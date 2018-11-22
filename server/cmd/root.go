package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xuender/go-kit"

	"../../cmds"
	"../gds"
)

var _cfgFile string

var rootCmd = &cobra.Command{
	Use:     "gds",
	Short:   "go disk server",
	Version: "v0.0.1",
	Long:    `网盘服务器`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dbDir, _ := filepath.Abs(cmds.GetString(cmd, cmds.DBPathStr))
		log.Println("数据库:", dbDir)
		db, err := kit.NewDB(dbDir)
		if err != nil {
			return err
		}
		defer db.Close()
		gds.Init(db, cmds.GetString(cmd, _tempPathStr), cmds.GetString(cmd, _filesPathStr))

		address := cmds.GetString(cmd, _address)
		// 地址端口号
		if !strings.HasPrefix(address, ":") {
			address = ":" + address
		}
		log.Println("端口号:", address)

		// 退出
		quitChan := make(chan os.Signal)
		signal.Notify(quitChan,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGHUP,
		)
		// 运行
		go gds.WebStart(address)

		<-quitChan
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	pflags := rootCmd.PersistentFlags()
	pflags.StringVarP(&_cfgFile, "config", "c", "", "配置文件")
	pflags.StringP(_tempPathStr, "t", "temp", "临时目录")
	pflags.StringP(_filesPathStr, "f", "files", "文件目录")

	flags := rootCmd.Flags()
	flags.StringP(cmds.DBPathStr, "d", "db", "数据库目录")
	flags.StringP(_address, "a", "6181", "访问地址端口号")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if _cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(_cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName(".gds")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		_cfgFile = viper.ConfigFileUsed()
		log.Println("读取配置文件:", _cfgFile)
	}
}
