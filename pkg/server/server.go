// Fargate Sidecar Injector HTTP WebServer
package server

import (
	"crypto/tls"
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

// Start starts webhook http server
func Start() {
	addr := strings.Join([]string{config.Host, fmt.Sprint(config.Port)}, ":")

	mux := http.NewServeMux()
	mux.HandleFunc("/mutate", mutatingWebhookHandler)
	mux.HandleFunc("/", rootHandler)

	log.Printf("Listening at: %s\n", addr)

	var listenErr error
	if config.TLSConfig.Enabled {
		serveTLS(addr, mux)
	} else {
		listenErr = http.ListenAndServe(addr, mux)
	}

	if listenErr != nil {
		_ = fmt.Errorf("failed to listen on address: [%s]", addr)
		log.Panic(listenErr)
	}
}

// serveTLS starts an HTTP server with TLS support using the provided address and handler.
func serveTLS(addr string, webhookMux *http.ServeMux) {
	certs, err := tls.X509KeyPair([]byte(config.TLSConfig.Cert), []byte(config.TLSConfig.Key))
	if err != nil {
		panic(fmt.Errorf("error loading TLS certificate and key: %w", err))
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{
			certs,
		},
	}

	listener, err := tls.Listen("tcp", addr, tlsConfig)
	if err != nil {
		panic(fmt.Errorf("listener error: %w", err))
	}
	defer listener.Close()

	httpServer := &http.Server{
		Handler: webhookMux,
	}
	log.Fatal(httpServer.Serve(listener))
}
