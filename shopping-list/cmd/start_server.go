package cmd

import (
	"context"
	"os/signal"
	"syscall"

	"shopping-list/internal/article/driver"
	"shopping-list/internal/article/driver/logger"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// startServerCmd represents the startServer command
var startServerCmd = &cobra.Command{
	Use: "start-server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer cancel()

		ctx = logger.ToContext(ctx, logger.New(viper.GetString("log_level")))
		srv := driver.NewHTTPServer(viper.GetString("server_addr"), driver.NewRouter(ctx))
		srv.Run(ctx)
	},
}

func init() {
	rootCmd.AddCommand(startServerCmd)

	startServerCmd.Flags().String("server-addr", ":8888", "The address of the server")
	viper.BindPFlag("server_addr", startServerCmd.Flags().Lookup("server-addr"))
}
