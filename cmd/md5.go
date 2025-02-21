/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"path/filepath"

	"github.com/WXZhang365/go-chunk-transform/utils"
	"github.com/spf13/cobra"
)

// md5Cmd represents the md5 command
var md5Cmd = &cobra.Command{
	Use:   "md5",
	Short: "计算单个文件的hash值",
	Long:  `计算单个文件的hash值`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			utils.ErrHandle("参数错误")
		}
		abs, err := filepath.Abs(args[0])
		utils.ErrHandle("无法转化为绝对路径", err)
		hash, err := utils.GetFileHash(abs)
		utils.ErrHandle("获取文件hash失败", err)
		cmd.Println(hash)
	},
}

func init() {
	rootCmd.AddCommand(md5Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// md5Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// md5Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
