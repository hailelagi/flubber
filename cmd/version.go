package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version",
	Long:  `Prints the version, yay semver!`,
	Run: func(cmd *cobra.Command, args []string) {
		version := cmd.Root().Version
		cmd.Println("flubber version: ", version)
	},
}
