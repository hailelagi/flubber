package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// unmountCmd represents the mount command
var unmountCmd = &cobra.Command{
	Use:   "unmount",
	Short: "unmount a filesystem at a directory",
	Long:  `mounts a filesystem at the specified mount point`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		rootCmd.Flags().StringVarP(&mountpoint, "mountpoint", "m", "", "mount point required!!")

		if err := rootCmd.MarkFlagRequired("mountpoint"); err != nil {
			log.Fatal(err)
			return err
		}

		exeCmd := exec.Command(fmt.Sprintf("fusermount -u %s", mountpoint))

		if err := exeCmd.Run(); err != nil {
			return err
		}

		cmd.Printf("unmounted filesystem")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(unmountCmd)
}

func init() {
	formatCmd.Flags().StringP("mountpoint", "m", "", "mount point required!")
	rootCmd.AddCommand(unmountCmd)
}
