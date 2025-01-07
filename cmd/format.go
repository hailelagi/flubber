package cmd

import (
	objectstore "github.com/hailelagi/flubber/internal/object_store"
	"github.com/spf13/cobra"
)

// formatCmd represents the format command
var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "mkfs the bucket",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		flags := map[string]string{
			"image": "Error getting image name:",
			"size":  "Error getting image size:",
			"block": "Error getting block size:",
		}

		values := make(map[string]string)
		for flag, errMsg := range flags {
			value, err := cmd.Flags().GetString(flag)
			if err != nil {
				cmd.PrintErrln(errMsg, err)
				return
			}
			values[flag] = value
		}

		cmd.MarkFlagsRequiredTogether("image", "size", "block")

		imageName := values["image"]
		imageSize := values["size"]
		blockSize := values["block"]

		objectstore.FormatBucket(imageName, imageSize, blockSize)
	},
}

func init() {
	formatCmd.Flags().StringP("image", "i", "", "Image name (required)")
	formatCmd.Flags().IntP("size", "s", 5120, "Size of the image to create on the bucket(required)")
	formatCmd.Flags().IntP("block", "b", 4096, "Size of blocks (defaults to 4KiB)")
	rootCmd.AddCommand(formatCmd)
}

func init() {
	rootCmd.AddCommand(formatCmd)
}
