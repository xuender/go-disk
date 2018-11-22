package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/gommon/color"
	"github.com/spf13/cobra"

	"../../server/gds"
)

var listCmd = &cobra.Command{
	Use:     "list [文件名...]",
	Aliases: []string{"l"},
	Short:   "列远程目录内容",
	Long: `
  列出目录下所有文件信息`,
	RunE: func(cmd *cobra.Command, args []string) error {
		color.Println("服务器地址:", color.Blue(serverURL))
		if len(args) == 0 {
			return list(".")
		}
		for _, path := range args {
			if err := list(path); err != nil {
				return err
			}
		}
		return nil
	},
}

func list(path string) error {
	color.Println("显示目录:", color.Green(path))
	bs, err := getBytes(fmt.Sprintf("%s/api/files/%s", serverURL, path))
	if err != nil {
		return err
	}
	files := []gds.File{}
	err = json.Unmarshal(bs, &files)
	if err != nil {
		return err
	}
	max := 1
	for _, f := range files {
		l := len(fmt.Sprintf("%d", f.Size))
		if max < l {
			max = l
		}
	}
	for _, f := range files {
		color.Println(
			f.Ca.Format(_format),                           // 创建时间
			color.Yellow(f.Mod.Format(_format)),            // 修改时间
			fmt.Sprintf(fmt.Sprintf("%%%dd", max), f.Size), // 文件尺寸
			color.Green(f.Name),                            // 文件名
		)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)
	// flags := listCmd.Flags()
	// flags.StringP(_server, "s", "http://localhost:6181", "服务器地址")
}
