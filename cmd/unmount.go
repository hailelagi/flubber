package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"
)

// unmountCmd represents the mount command
var unmountCmd = &cobra.Command{
	Use:   "unmount",
	Short: "unmount a filesystem",
	Long:  `mounts a filesystem at the specified mount point`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if mp := args[0]; mp != "" {
			mountpoint = mp
		}

		exeCmd := exec.Command("fusermount", "-u", mountpoint)

		if err := exeCmd.Run(); err != nil {
			return err
		} else {
			cmd.Println("unmounted filesystem")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(unmountCmd)
}
