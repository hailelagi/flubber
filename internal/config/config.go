package config

import (
	"time"

	"github.com/spf13/viper"
)

type Storage struct {
	Endpoint             string
	AccessKeyID          string
	SecretAccessKey      string
	AllowedBuckets       map[string]string
	EnableBucketPolicies bool
	UseSSL               bool
}

type Mount struct {
	Debug      bool
	Profile    string
	MemProfile string
	Ttl        *time.Duration
}

func GetStorageConfig() *Storage {
	return &Storage{
		Endpoint:             viper.GetString("bucket.url"),
		AccessKeyID:          viper.GetString("credentials.access_key_id"),
		SecretAccessKey:      viper.GetString("credentials.secret_access_key"),
		UseSSL:               viper.GetBool("bucket.ssl") || false,
		EnableBucketPolicies: viper.GetBool("bucket.policies") || false,
		AllowedBuckets:       map[string]string{"public": "read", "private": "all", "local": "all"},
	}
}
