package server

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HTTP handler for server
func handler(rw http.ResponseWriter, req *http.Request) {

	req.Close = true
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// remoteAddr, _, _ := net.SplitHostPort(req.RemoteAddr)

	// TODO: Do we need this
	// req.Header.Set("X-Forwarded-For", remoteAddr)
	// req.Header.Add("Authorization", strings.Join([]string{"Bearer", config.Token}, " "))

	client := http.Client{Timeout: 300 * time.Second}
	resp, err := client.Do(req)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(rw, err)
		return
	}

	for key, values := range resp.Header {
		for _, value := range values {
			rw.Header().Set(key, value)
		}
	}

	defer resp.Body.Close()

	rw.WriteHeader(resp.StatusCode)

	var data []byte
	reader := bytes.NewBuffer(data)

	_, readerErr := io.Copy(reader, resp.Body)
	data = reader.Bytes()

	if readerErr != nil {
		_ = fmt.Errorf("ReaderError: %s", readerErr)
	}

	// Do work here
	// TODO: replace with actual injected data
	rw.Write(data)
}
