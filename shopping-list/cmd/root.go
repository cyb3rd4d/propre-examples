package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "shopping-list",
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
	rootCmd.PersistentFlags().StringVar(&configFile, "config-file", ".env", "The path of the configuration file")
}

func initConfig() {
	viper.SetEnvPrefix("propre")
	viper.AutomaticEnv()

	if configFile == "" {
		return
	}

	if err := godotenv.Load(configFile); err != nil {
		fmt.Fprintf(os.Stderr, "could not load the configuration file: %s\n", err)
		os.Exit(1)
	}
}
