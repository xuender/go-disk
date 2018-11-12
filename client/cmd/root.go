package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var serverURL string

var rootCmd = &cobra.Command{
	Use:     "gdc",
	Short:   "go disk client",
	Version: "v0.0.1",
	Long:    `网盘客户端`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("连接服务器")
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
	rootCmd.PersistentFlags().StringVarP(&serverURL, "server", "s", "http://localhost:6181", "服务器地址")
	// flags := rootCmd.Flags()
	// flags.StringVarP(&_dbPath, _db, "d", "db", "数据库目录")
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
		viper.SetConfigName(".gdc")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		cfgFile = viper.ConfigFileUsed()
		log.Println("读取配置文件:", cfgFile)
	}
}
