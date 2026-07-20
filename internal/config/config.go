package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	System    SystemConfig    `mapstructure:"system"`
	Bootstrap BootstrapConfig `mapstructure:"bootstrap"`
	Plugins   PluginsConfig   `mapstructure:"plugins"`
	Profiles  ProfilesConfig  `mapstructure:"profiles"`
	Logging   LoggingConfig   `mapstructure:"logging"`
}

type SystemConfig struct {
	LogLevel string `mapstructure:"log_level"`
	DataDir  string `mapstructure:"data_dir"`
	StateDB  string `mapstructure:"state_db"`
}

type BootstrapConfig struct {
	StrictMode bool `mapstructure:"strict_mode"`
	AutoUpdate bool `mapstructure:"auto_update"`
}

type PluginsConfig struct {
	Paths []string `mapstructure:"paths"`
}

type ProfilesConfig struct {
	Active string `mapstructure:"active"`
}

type LoggingConfig struct {
	Format   string `mapstructure:"format"`
	Rotation string `mapstructure:"rotation"`
}

var Cfg Config

// Load reads in config file and ENV variables if set.
func Load(cfgFile string) error {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		viper.AddConfigPath(filepath.Join(home, ".config", "platform"))
		viper.AddConfigPath("/etc/platform/")
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		// It's ok if there isn't a config file, we can use defaults
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("error reading config file: %w", err)
		}
	}

	// Set Defaults
	viper.SetDefault("system.log_level", "info")
	viper.SetDefault("logging.format", "json")

	if err := viper.Unmarshal(&Cfg); err != nil {
		return fmt.Errorf("unable to decode into struct: %w", err)
	}

	return nil
}
