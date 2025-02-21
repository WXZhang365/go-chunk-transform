/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/WXZhang365/go-chunk-transform/utils"
	"github.com/spf13/cobra"
)

// splitCmd represents the split command
var splitCmd = &cobra.Command{
	Use:   "split",
	Short: "用于将文件切分为对应的块大小",
	Long:  `用于将文件切分为对应的块大小,并且生成对应的chunk.list的chunk<->hash索引文件`,
	Run: func(cmd *cobra.Command, args []string) {
		size, _ := cmd.Flags().GetInt("size")
		output, _ := cmd.Flags().GetString("output")
		target, _ := cmd.Flags().GetString("target")
		if size <= 0 || len(output) == 0 || len(target) == 0 {
			utils.ErrHandle("参数错误")
		}
		tAbs, tErr := filepath.Abs(target)
		oAbs, oErr := filepath.Abs(output)
		fmt.Printf("参数: size: %dm |targetpath abs:%s | outputpath abs:%s\n", size, tAbs, oAbs)
		utils.ErrHandle("无法转化为绝对路径", tErr, oErr)
		open, err := os.Open(tAbs)
		utils.ErrHandle("文件打开失败", err)
		defer open.Close()
		chunkList := ""
		buffer := make([]byte, size*1024*1024)
		chunkNumber := 0
		for {
			bytesRead, err := open.Read(buffer)
			if err != nil && err != io.EOF {
				utils.ErrHandle("读取文件内容失败", err)
			}
			if bytesRead == 0 {
				break
			}
			chunkFileName := fmt.Sprintf("%s.%d.chunk", path.Join(output, target), chunkNumber)
			chunkFilePath := chunkFileName
			chunkFile, err := os.Create(chunkFilePath)
			utils.ErrHandle("创建文件失败", err)
			_, err = chunkFile.Write(buffer[:bytesRead])
			if err != nil {
				chunkFile.Close()
				utils.ErrHandle("写入文件失败", err)
			}

			chunkFile.Close()
			chunkNumber++
			hash, err := utils.GetFileHash(chunkFilePath)
			utils.ErrHandle("获取文件hash失败", err)
			chunkFileNameWithLength := fmt.Sprintf("%04X%s", len(chunkFileName), chunkFileName)
			chunkList += fmt.Sprintf("%s%s\n", chunkFileNameWithLength, hash)
		}
		err = os.WriteFile(fmt.Sprintf("%s/chunk.list", output), []byte(chunkList), 0644)
		utils.ErrHandle("写入chunk.list失败", err)
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(splitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// splitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	splitCmd.Flags().IntP("size", "s", 64, "每个chunk的大小,默认64M(单位为M)")
	splitCmd.Flags().StringP("output", "o", "output/", "输出到哪个目录")
	splitCmd.Flags().StringP("target", "t", "", "要切分的目标文件")
}
