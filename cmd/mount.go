package cmd

import (
	"log"

	"github.com/hailelagi/flubber/fuse"
	"github.com/spf13/cobra"
)

var directory string

// mountCmd represents the mount command
var mountCmd = &cobra.Command{
	Use:   "mount",
	Short: "mount a filesystem at a directory",
	Long:  `mounts a filesystem at the specified directory`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Flags().StringVar(&directory, "dir", "/flubber-fuse", "mount directory")

		if err := cmd.MarkFlagRequired("dir"); err != nil {
			log.Fatal(err)
			return err
		}

		return fuse.Mount(directory)
	},
}

func init() {
	rootCmd.AddCommand(mountCmd)
}
