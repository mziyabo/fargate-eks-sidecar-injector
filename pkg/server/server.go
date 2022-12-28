// Fargate Sidecar Injector HTTP WebServer
package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/mziyabo/fargate-sidecar-injector/m/v2/pkg/shared"
)

var config shared.FargateSidecarInjectorConfig

func init() {
	config = shared.NewConfig()
}

// Start webhook http server
func Start() {
	addr := strings.Join([]string{config.Host, fmt.Sprint(config.Port)}, ":")
	proxy := http.HandlerFunc(handler)

	log.Printf("Listening at: %s\n", addr)

	var listenErr error
	if config.TLSConfig.Enabled {
		listenErr = http.ListenAndServeTLS(addr, config.TLSConfig.Cert, config.TLSConfig.Key, proxy)
	} else {
		listenErr = http.ListenAndServe(addr, proxy)
	}

	if listenErr != nil {
		_ = fmt.Errorf("failed to listen on address: [%s]", addr)
		log.Panic(listenErr)
	}
}
