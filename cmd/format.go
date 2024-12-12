package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// formatCmd represents the format command
var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "format the bucket to this filesystem",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("format called")
	},
}

func init() {
	rootCmd.AddCommand(formatCmd)
}
