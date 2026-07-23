package cmd

import (
	"fmt"
	"os"

	"github.com/base-infrastructure/platform/internal/bootstrap"
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
	if err := bootstrap.Current.Initialize(cfgFile); err != nil {
		fmt.Println("Error initializing platform:", err)
		os.Exit(1)
	}
}
