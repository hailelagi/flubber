package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "configure bucket credentials",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		requiredFlags := map[string]string{
			"bucket_url":        "",
			"bucket_name":       "",
			"access_key_id":     "",
			"secret_access_key": "",
		}

		for flag := range requiredFlags {
			value, err := cmd.Flags().GetString(flag)
			if err != nil || value == "" {
				cmd.PrintErrf("Error: %s is required and cannot be empty\n", flag)
				return
			}
			requiredFlags[flag] = value
		}

		viper.Set("bucket.url", requiredFlags["bucket_url"])
		viper.Set("bucket.name", requiredFlags["bucket_name"])
		viper.Set("credentials.access_key_id", requiredFlags["access_key_id"])
		viper.Set("credentials.secret_access_key", requiredFlags["secret_access_key"])

		requiredFields := []string{
			"bucket.url",
			"bucket.name",
			"credentials.access_key_id",
			"credentials.secret_access_key",
		}

		for _, field := range requiredFields {
			if !viper.IsSet(field) {
				cmd.PrintErrf("Missing required config: %s\n", field)
				return
			}
		}

		if err := viper.SafeWriteConfigAs(viper.ConfigFileUsed()); err != nil {
			if viper.ConfigFileUsed() == "" {
				home, err := os.UserHomeDir()
				if err != nil {
					cmd.PrintErrf("Error getting user home directory: %v\n", err)
					return
				}
				configFile := filepath.Join(home, ".config", "flubber", "config.yaml")

				err = os.MkdirAll(filepath.Dir(configFile), 0755)
				if err != nil {
					cmd.PrintErrf("Error creating config directory: %v\n", err)
					return
				}

				if err := viper.SafeWriteConfigAs(configFile); err != nil {
					cmd.PrintErrln("warning overwriting previous config.")
					if err := viper.WriteConfigAs(configFile); err != nil {
						cmd.PrintErrf("Error saving config: %v\n", err)
						return
					}
				}

			} else {
				cmd.PrintErrf("Error saving config: %v\n", err)
				return
			}
		}

		cmd.Println("S3 credentials configured and saved successfully")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().String("bucket_url", "", "URL of the S3 bucket (required)")
	configCmd.Flags().String("bucket_name", "", "Name of the S3 bucket (required)")
	configCmd.Flags().String("access_key_id", "", "Access key ID for the S3 bucket (required)")
	configCmd.Flags().String("secret_access_key", "", "Secret access key for the S3 bucket (required)")

	configCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		cmd.Println("Usage:")
		cmd.Println("  config [flags]")
		cmd.Println()
		cmd.Println("Flags:")
		cmd.Println("  --bucket_url string        URL of the S3 bucket (required)")
		cmd.Println("  --bucket_name string       Name of the S3 bucket (required)")
		cmd.Println("  --access_key_id string     Access key ID for the S3 bucket (required)")
		cmd.Println("  --secret_access_key string Secret access key for the S3 bucket (required)")
		cmd.Println()
		cmd.Println("Global Flags:")
		cmd.Println("  -h, --help   help for config")
	})
}
