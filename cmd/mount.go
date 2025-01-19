package cmd

import (
	"time"

	"github.com/hailelagi/flubber/internal/config"
	"github.com/hailelagi/flubber/internal/fs"
	"github.com/spf13/cobra"
)

var mountpoint string
var mntConfig config.Mount

// mountCmd represents the mount command
var mountCmd = &cobra.Command{
	Use:   "mount <./example_dir> [flags]",
	Short: "mount a filesystem at a directory",
	Long:  `mount a filesystem at the specified mount point`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if mp := args[0]; mp != "" {
			mountpoint = mp
		}

		return fs.InitMount(mountpoint, &mntConfig)
	},
}

func init() {
	rootCmd.AddCommand(mountCmd)

	mountCmd.Flags().StringVarP(&mntConfig.Profile, "profile", "p", "profile.dat", "cpu profile")
	mountCmd.Flags().StringVar(&mntConfig.MemProfile, "memprofile", "memprofile.dat", "memory profile")
	mntConfig.Ttl = mountCmd.Flags().Duration("ttl", time.Second, "attribute/entry cache TTL")
	mntConfig.Debug = *mountCmd.Flags().BoolP("debug", "d", true, "debug")
}
