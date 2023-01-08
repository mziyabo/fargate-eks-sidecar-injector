package shared

import (
	"fmt"
	"net/url"

	"encoding/base64"

	"github.com/spf13/viper"
)

func init() {
	parseConfig()

	//parseConfig()
	// read config
}

// Initialize webhook config
func NewConfig() FargateSidecarInjectorConfig {
	// TODO: Clean this up - Lots of experimentation and repetition here
	cert, _ := base64.StdEncoding.DecodeString(viper.GetString("serve.tls.cert"))
	key, _ := base64.StdEncoding.DecodeString(viper.GetString("serve.tls.certKey"))

	config := FargateSidecarInjectorConfig{
		Port: viper.GetInt("serve.port"),
		Host: viper.GetString("serve.host"),

		TLSConfig: TLSClientConfig{
			Enabled: viper.GetBool("serve.tls.enabled"),
			Cert:    string(cert),
			Key:     string(key),
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
}

// Parse webhook server configuration
func parseConfig() {
	// Name of config file (without extension)
	viper.SetConfigName("fargatesidecarinjector.conf")
	// REQUIRED if the config file does not have the extension in the name
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/fargatesidecarinjector")
	viper.AddConfigPath("$HOME/.fargatesidecarinjector")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
