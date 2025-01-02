package cmd

import (
	"log"

	"github.com/hailelagi/flubber/internal/fuse"
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
		rootCmd.Flags().StringVarP(&directory, "dir", "d", "", "mount dir required!!")

		if err := rootCmd.MarkFlagRequired("dir"); err != nil {
			log.Fatal(err)
			return err
		}

		return fuse.Mount(directory)
	},
}

func init() {
	rootCmd.AddCommand(mountCmd)
}
