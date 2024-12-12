package cmd

import (
	"github.com/hailelagi/flubber/fuse"
	"github.com/spf13/cobra"
)

// unmountCmd represents the mount command
var unmountCmd = &cobra.Command{
	Use:   "unmount",
	Short: "unmount a filesystem at a directory",
	Long:  `mounts a filesystem at the specified directory`,
	Run: func(cmd *cobra.Command, args []string) {
		fuse.Mount()
	},
}

func init() {
	rootCmd.AddCommand(unmountCmd)
}
