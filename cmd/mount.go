package cmd

import (
	"time"

	fs "github.com/hailelagi/flubber/internal"
	"github.com/spf13/cobra"
)

var mountpoint string
var config fs.MntConfig

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

		return fs.InitMount(mountpoint, &config)
	},
}

func init() {
	rootCmd.AddCommand(mountCmd)

	mountCmd.Flags().StringVarP(&config.Profile, "profile", "p", "profile.dat", "cpu profile")
	mountCmd.Flags().StringVar(&config.MemProfile, "memprofile", "memprofile.dat", "memory profile")
	config.Ttl = mountCmd.Flags().Duration("ttl", time.Second, "attribute/entry cache TTL")
	config.Debug = *mountCmd.Flags().BoolP("debug", "d", true, "debug")
}
