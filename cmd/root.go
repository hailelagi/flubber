package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "flubber",
		Short: "A FUSE filesystem built on s3",
		Long: `flubber is a filesystem in userspace (FUSE)
that shims out the backing storage to a block storage service such as s3`,
		Version: "0.0.1",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
