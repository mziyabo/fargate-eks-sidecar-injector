// Fargate Sidecar Injector HTTP WebServer
package server

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/mziyabo/eks-fargate-sidecar-injector/m/v2/pkg/shared"
)

var config shared.FargateSidecarInjectorConfig

func init() {
	config = shared.NewConfig()
}

// Listen starts webhook http server
func StartListening() {
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

// serveTLS starts an HTTP server for the Mutating WebHook
func serveTLS(addr string, webhookMux *http.ServeMux) {
	var certPool = x509.NewCertPool()
	certPool.AppendCertsFromPEM([]byte(config.TLSConfig.CA))
	certs, err := tls.X509KeyPair([]byte(config.TLSConfig.Cert), []byte(config.TLSConfig.Key))
	if err != nil {
		panic(fmt.Errorf("error loading TLS certificate and key: %w", err))
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{
			certs,
		},
		RootCAs:            certPool,
		InsecureSkipVerify: false, // TODO: InsecureSkipVerify - not ideal, find a better way.
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
