package cmd

import (
	"log"
	"time"

	fs "github.com/hailelagi/flubber/internal/fuse"
	"github.com/spf13/cobra"
)

var mountpoint string
var config fs.MntConfig

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

		return fs.InitMount(mp, &config)
	},
}

func init() {
	mountCmd.Flags().StringVarP(&mountpoint, "mountpoint", "m", "", "mount point required!")
	mountCmd.Flags().StringVarP(&config.Profile, "profile", "p", "profile.dat", "cpu profile")
	mountCmd.Flags().StringVarP(&config.MemProfile, "memprofile", "mem", "memprofile.dat", "memory profile")
	config.Ttl = mountCmd.Flags().DurationP("ttl", "ttl", time.Second, "attribute/entry cache TTL")
	config.Debug = *mountCmd.Flags().BoolP("debug", "d", true, "debug")

	if err := mountCmd.MarkFlagRequired("mountpoint"); err != nil {
		log.Fatal(err)
	}

	rootCmd.AddCommand(mountCmd)
}
