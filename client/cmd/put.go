package cmd

import (
	"fmt"
	"log"

	"github.com/labstack/gommon/color"
	"github.com/spf13/cobra"
	utils "github.com/xuender/go-utils"

	"../../cmds"
	"../gdc"
)

var upCmd = &cobra.Command{
	Use:     "up [文件名...]",
	Aliases: []string{"l"},
	Short:   "上传文件",
	Long: `
  将文件上传到指定目录`,
	RunE: func(cmd *cobra.Command, args []string) error {
		color.Println("服务器地址:", color.Blue(serverURL))
		path := cmds.GetString(cmd, _path)
		for _, f := range args {
			color.Println("上传文件:", color.Green(f))
			fid, err := utils.NewFileID(f)
			if err != nil {
				return err
			}
			// TODO 校验FileID
			log.Println(fid)
			code, body, err := gdc.PostFile(f, fmt.Sprintf("%s/api/files/%s", serverURL, path))
			if err != nil {
				return err
			}
			fmt.Println(string(body), code)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
	flags := upCmd.Flags()
	flags.StringP(_path, "p", ".", "上传目录")
}
