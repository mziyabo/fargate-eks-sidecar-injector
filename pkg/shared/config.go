package shared

import (
	"fmt"
	"net/url"

	"encoding/base64"

	"github.com/spf13/viper"
)

func init() {
	parseConfig()
}

// Initialize webhook config
func NewConfig() FargateSidecarInjectorConfig {
	// TODO: Clean this up
	cert, _ := base64.StdEncoding.DecodeString(viper.GetString("serve.tls.cert"))
	key, _ := base64.StdEncoding.DecodeString(viper.GetString("serve.tls.certKey"))
	ca, _ := base64.StdEncoding.DecodeString(viper.GetString("serve.tls.ca"))

	config := FargateSidecarInjectorConfig{
		Port: viper.GetInt("serve.port"),
		Host: viper.GetString("serve.host"),

		TLSConfig: TLSClientConfig{
			Enabled: viper.GetBool("serve.tls.enabled"),
			Cert:    string(cert),
			Key:     string(key),
			CA:      string(ca),
		},
	}

	return config
}

// Fargate sidecar injector webhook configuration
type FargateSidecarInjectorConfig struct {
	ProxyURL  *url.URL
	Port      int
	Host      string
	TLSConfig TLSClientConfig
	Token     string // serviceaccount token
}

type TLSClientConfig struct {
	Enabled bool
	Cert    string
	Key     string
	CA      string
}

// parseConfig reads config from file
func parseConfig() {
	// Name of config file (without extension)
	viper.SetConfigName("fargatesidecarinjector.conf")

	// REQUIRED: if the config file does not have the extension in the name
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/fargatesidecarinjector")
	viper.AddConfigPath("$HOME/.fargatesidecarinjector")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
