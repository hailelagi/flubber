package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "flubber",
		Short: "A FUSE filesystem built on object storage",
		Long: `flubber is a filesystem in userspace (FUSE)
that shims out the backing storage to a block storage service such as object storage`,
		Version: "0.0.1",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile == "" {
		viper.SetConfigFile(cfgFile)
		home, err := os.UserHomeDir()

		if err != nil {
			cobra.CheckErr(err)
			return
		}

		configFile := filepath.Join(home, ".config", "flubber", "config.yaml")

		err = os.MkdirAll(filepath.Dir(configFile), 0o755)
		if err != nil {
			cobra.CheckErr(err)
			return
		}

		viper.SetConfigFile(configFile)
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		cobra.CheckErr(fmt.Errorf("Error reading config file: %v", err))
	}
}
