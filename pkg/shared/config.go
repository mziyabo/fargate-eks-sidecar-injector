package shared

import (
	"fmt"
	"log"
	"net/url"

	"encoding/base64"

	"github.com/spf13/viper"
)

func init() {
	parseConfig()
}

// Initialize webhook config
func NewConfig() FargateSidecarInjectorConfig {
	ca := viperGetB64("serve.tls.ca")
	cert := viperGetB64("serve.tls.cert")
	key := viperGetB64("serve.tls.certKey")

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
		log.Fatalln(fmt.Errorf("fatal error config file: %w", err))
	}
}

// viperGetB64 base64 decodes and returns config setting
func viperGetB64(setting string) []byte {
	d, err := base64.StdEncoding.DecodeString(viper.GetString(setting))
	if err != nil {
		log.Fatalln(err)
	}
	return d
}
