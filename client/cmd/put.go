package cmd

import (
	"github.com/labstack/gommon/color"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:     "up [文件名...]",
	Aliases: []string{"l"},
	Short:   "上传文件",
	Long: `
  将文件上传到指定目录`,
	RunE: func(cmd *cobra.Command, args []string) error {
		color.Println("服务器地址:", color.Blue(serverURL))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}
