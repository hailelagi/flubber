package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"
)

// unmountCmd represents the mount command
var unmountCmd = &cobra.Command{
	Use:   "unmount",
	Short: "unmount a filesystem",
	Long:  `unmount a previously mounted filesystem at a mountpoint`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if mp := args[0]; mp != "" {
			mountpoint = mp
		}

		exeCmd := exec.Command("fusermount", "-u", mountpoint)

		if err := exeCmd.Run(); err != nil {
			return err
		} else {
			cmd.Println("\n unmounted filesystem")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(unmountCmd)

	unmountCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		cmd.Println("Usage:")
		cmd.Println("  unmount <./example_dir>")
		cmd.Println()
		cmd.Println("Flags:")
		cmd.Println("  -h, --help   help for config")
	})
}
