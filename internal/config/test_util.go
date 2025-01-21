package config

import "github.com/spf13/viper"

func SetupTestConfig() {
	viper.Set("bucket.url", "localhost:9000")
	viper.Set("bucket.name", "test")
	viper.Set("credentials.access_key_id", "minioadmin")
	viper.Set("credentials.secret_access_key", "minioadmin")

}
