package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "flubber",
	Short:   "A FUSE filesystem built on s3",
	Long:    `flubber is a filesystem in userspace that shims out the backing storage to a block storage service such as s3`,
	Version: "0.0.1",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.flubber.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
