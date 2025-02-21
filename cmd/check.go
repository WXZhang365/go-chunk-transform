/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/WXZhang365/go-chunk-transform/utils"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "校验文件完整性",
	Long:  `通过chunkList校验文件完整性`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := cmd.Flags().GetString("dir")
		chunkListPath, err := filepath.Abs(filepath.Join(dir, "chunk.list"))
		utils.ErrHandle("无法转化为绝对路径", err)
		file, err := os.Open(chunkListPath)
		utils.ErrHandle("无法打开chunk.list文件", err)
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			// 提取文件名长度和文件名
			fileNameLengthHex := line[0:4]
			fileNameLength, err := strconv.ParseInt(fileNameLengthHex, 16, 32)
			if err != nil {
				utils.ErrHandle("无法解析文件名长度", err)
			}
			chunkFileName := line[4 : 4+fileNameLength]
			expectedHash := line[4+fileNameLength:]
			abs, err := filepath.Abs(chunkFileName)
			utils.ErrHandle("转换到绝对路径错误", err)
			// 计算实际的文件哈希值
			actualHash, err := utils.GetFileHash(abs)
			if err != nil {
				utils.ErrHandle("无法计算文件哈希值", err)
			}

			// 比较哈希值
			if actualHash != expectedHash {
				fmt.Printf("文件 %s 的哈希值不匹配: 期望 %s, 实际 %s\n", chunkFileName, expectedHash, actualHash)
			} else {
				fmt.Printf("文件 %s 的哈希值匹配\n", chunkFileName)
			}
		}

		if err := scanner.Err(); err != nil {
			utils.ErrHandle("读取chunk.list文件时出错", err)
		}
	},
}

type chunk struct {
	FileNameLen int
	FileName    string
	Hash        string
}

func readChunkList(filename string, channel chan *chunk) {
	file, err := os.Open(filename)
	utils.ErrHandle("无法打开chunk.list文件", err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fileNameLengthHex := line[0:4]
		fileNameLength, err := strconv.ParseInt(fileNameLengthHex, 16, 32)
		utils.ErrHandle("无法解析文件名长度", err)
		chunkFileName := line[4 : 4+fileNameLength]
		expectedHash := line[4+fileNameLength:]
		channel <- &chunk{
			FileNameLen: int(fileNameLength),
			FileName:    chunkFileName,
			Hash:        expectedHash,
		}
	}
	channel <- nil
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	checkCmd.Flags().StringP("dir", "d", "output", "从output文件验证chunkList是否正确无误")
}
