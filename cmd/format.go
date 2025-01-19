package cmd

import (
	"github.com/hailelagi/flubber/internal/storage"
	"github.com/spf13/cobra"
)

// formatCmd represents the format command
var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "mkfs the bucket",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.MarkFlagsRequiredTogether("page", "size")

		imageName, err := cmd.Flags().GetString("image")
		if err != nil {
			cmd.PrintErrln("Error getting image name:", err)
			return
		}

		blockSize, err := cmd.Flags().GetInt("size")
		if err != nil {
			cmd.PrintErrln("Error getting image size:", err)
			return
		}

		pageSize, err := cmd.Flags().GetInt("page")
		if err != nil {
			cmd.PrintErrln("Error getting page size:", err)
			return
		}

		storage.FormatBucket(imageName, pageSize, blockSize)
	},
}

func init() {
	rootCmd.AddCommand(formatCmd)

	formatCmd.Flags().StringP("image", "i", "", "Image name (required)")
	formatCmd.Flags().IntP("size", "s", 5120, "blocks to preallocate in bucket")
	formatCmd.Flags().IntP("page", "p", 4096, "page size (defaults to 4KiB)")
}

func init() {
	rootCmd.AddCommand(formatCmd)
}
