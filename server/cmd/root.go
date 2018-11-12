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

	"../../cmds"
	"../gds"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:     "gds",
	Short:   "go disk server",
	Version: "v0.0.1",
	Long:    `网盘服务器`,
	RunE: func(cmd *cobra.Command, args []string) error {
		db, _ := filepath.Abs(cmds.GetString(cmd, cmds.DBPathStr))
		log.Println("数据库:", db)
		address := cmds.GetString(cmd, _address)
		// 地址端口号
		if !strings.HasPrefix(address, ":") {
			address = ":" + address
		}
		log.Println("端口号:", address)
		web, err := gds.NewWeb(db)
		if err != nil {
			return err
		}
		defer web.Close()
		// 退出
		quitChan := make(chan os.Signal)
		signal.Notify(quitChan,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGHUP,
		)
		// 运行
		go web.Start(address)
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
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "配置文件")
	flags := rootCmd.Flags()
	flags.StringP(cmds.DBPathStr, "d", "db", "数据库目录")
	flags.StringP(_address, "a", "6181", "访问地址端口号")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
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
		cfgFile = viper.ConfigFileUsed()
		log.Println("读取配置文件:", cfgFile)
	}
}
