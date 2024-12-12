package cmd

import (
	"github.com/hailelagi/flubber/fuse"
	"github.com/spf13/cobra"
)

// mountCmd represents the mount command
var mountCmd = &cobra.Command{
	Use:   "mount",
	Short: "mount a filesystem at a directory",
	Long:  `mounts a filesystem at the specified directory`,
	Run: func(cmd *cobra.Command, args []string) {
		fuse.Mount()
	},
}

func init() {
	rootCmd.AddCommand(mountCmd)
}
