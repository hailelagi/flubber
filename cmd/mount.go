package cmd

import (
	"log"

	"github.com/hailelagi/flubber/internal/fuse"
	"github.com/spf13/cobra"
)

var mountpoint string

// mountCmd represents the mount command
var mountCmd = &cobra.Command{
	Use:   "mount",
	Short: "mount a filesystem at a directory",
	Long:  `mounts a filesystem at the specified mount point`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		rootCmd.Flags().StringVarP(&mountpoint, "mountpoint", "m", "", "mount point required!!")

		if err := rootCmd.MarkFlagRequired("mountpoint"); err != nil {
			log.Fatal(err)
			return err
		}

		return fuse.Mount(mountpoint)
	},
}

func init() {
	rootCmd.AddCommand(mountCmd)
}
