package cmd

import (
	"log"

	"github.com/hailelagi/flubber/internal/fuse"
	"github.com/spf13/cobra"
)

var mountpoint string

// mountCmd represents the mount command
var mountCmd = &cobra.Command{
	Use:   "mount -m=<dir>",
	Short: "mount a filesystem at a directory",
	Long:  `mounts a filesystem at the specified mount point`,
	RunE: func(cmd *cobra.Command, args []string) error {
		mp, err := cmd.Flags().GetString("mountpoint")
		if err != nil {
			log.Fatal(err)
			return err
		}

		return fuse.Mount(mp)
	},
}

func init() {
	mountCmd.Flags().StringVarP(&mountpoint, "mountpoint", "m", "", "mount point required!")

	if err := mountCmd.MarkFlagRequired("mountpoint"); err != nil {
		log.Fatal(err)
	}

	rootCmd.AddCommand(mountCmd)
}
