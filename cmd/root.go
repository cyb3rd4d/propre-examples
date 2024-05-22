package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	logLevel   string
	configFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "poc-propre",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringVar(&logLevel, "log-level", "info", "The level of the logger")
	rootCmd.Flags().StringVar(&configFile, "config-file", "", "The path of the configuration file")

	viper.BindPFlag("log_level", rootCmd.Flags().Lookup("log-level"))
}

func initConfig() {
	if configFile == "" {
		return
	}

	viper.SetConfigFile(configFile)
	viper.SetEnvPrefix("propre")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "could not load the configuration file: %s", err)
		os.Exit(1)
	}
}
