package cmd

import (
	"github.com/WXZhang365/go-chunk-transform/serve"
	"github.com/spf13/cobra"
)

// splitCmd represents the split command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "启动一个文件服务器",
	Long:  `启动一个简易的文件传输服务器/渲染服务器`,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		host, _ := cmd.Flags().GetString("host")
		root, _ := cmd.Flags().GetString("root")
		httpServe := serve.FileServe(root)
		httpServe(host, port)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// splitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	serveCmd.Flags().IntP("port", "p", 20333, "Port")
	serveCmd.Flags().StringP("host", "H", "0.0.0.0", "Host")
	serveCmd.Flags().StringP("root", "r", "./", "root path")
	// serveCmd.Flags().StringP("mode", "m", "f", "Serve mode(f:文件传输模式,s:web服务器模式)")
}
