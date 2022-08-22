package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"os/signal"
	"skychart/config"
	"skychart/server"
	"syscall"
)

var (
	defaultUpdateFreq = "@daily"
	configFile        string
	defaultConfigFile = "~/.skychart/config.toml"
	ServerCmd         = &cobra.Command{
		Use: "skychart [--config config]",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runServer()
		},
	}
)

func init() {
	ServerCmd.Flags().StringVar(&configFile, "config", defaultConfigFile, "Config file")
}

func runServer() error {
	cfg, err := config.ReadConfigFile(configFile)
	if err != nil {
		fmt.Print(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	return server.Serve(ctx, cfg.RegistryUrl, cfg.Port, defaultUpdateFreq)
}
