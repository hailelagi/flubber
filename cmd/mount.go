package cmd

import (
	"log"
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
	Long:  `mounts a filesystem at the specified mount point`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if mp := args[0]; mp != "" {
			mountpoint = mp
		}

		config.Ttl = cmd.Flags().Duration("ttl", time.Second, "attribute/entry cache TTL")
		config.Debug = *cmd.Flags().BoolP("debug", "d", true, "debug")

		if err := cmd.MarkFlagRequired("mountpoint"); err != nil {
			log.Fatal(err)
		}

		return fs.InitMount(mountpoint, &config)
	},
}

func init() {
	rootCmd.AddCommand(mountCmd)

	mountCmd.Flags().StringVarP(&mountpoint, "mountpoint", "m", "", "mount point required!")
	mountCmd.Flags().StringVarP(&config.Profile, "profile", "p", "profile.dat", "cpu profile")
	mountCmd.Flags().StringVar(&config.MemProfile, "memprofile", "memprofile.dat", "memory profile")
}
