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

// mergeCmd represents the merge command
var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "合并文件",
	Long:  `合并文件，但是不会校验完整性`,
	Run: func(cmd *cobra.Command, args []string) {
		target, _ := cmd.Flags().GetString("target")
		output, _ := cmd.Flags().GetString("output")
		tDAbs, tErr := filepath.Abs(target)
		tOAbs, oErr := filepath.Abs(output)
		utils.ErrHandle("无法转化为绝对路径", tErr, oErr)
		filename := path.Join(tDAbs, "chunk.list")
		channel := make(chan *chunk)
		flag := make(chan error)
		go readChunkList(filename, channel)
		go func() {
			for c := range channel {
				if c == nil {
					break
				}
				fmt.Printf("合并文件: %s\n", c.FileName)
				out, err := os.Create(tOAbs)
				flag <- err
				defer out.Close()
				chunk, err := os.Open(c.FileName)
				flag <- err
				defer chunk.Close()
				_, err = io.Copy(out, chunk)
				flag <- err
				chunk.Close()
			}
			flag <- nil
		}()
		status := <-flag
		if status == nil {
			fmt.Println("文件合并完成")
		} else {
			utils.ErrHandle("文件合并失败", status)
		}
	},
}

func init() {
	rootCmd.AddCommand(mergeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mergeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	mergeCmd.Flags().StringP("output", "o", "", "输出的文件名,默认使用chunk.list中第一个文件名")
	mergeCmd.Flags().StringP("target", "t", "output", "默认从output文件夹中合并")
}
