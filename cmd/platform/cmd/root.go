package cmd

import (
	"fmt"
	"os"

	"github.com/base-infrastructure/platform/internal/config"
	"github.com/base-infrastructure/platform/internal/logger"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "platform",
	Short: "Universal Bootstrap Framework",
	Long:  `A production-grade, self-hosted engineering platform designed to provision, configure, validate, update, recover, and maintain heterogeneous environments.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/platform/config.yaml)")
}

func initConfig() {
	if err := config.Load(cfgFile); err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	logger.Init(config.Cfg.System.LogLevel, config.Cfg.Logging.Format)
}
