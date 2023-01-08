package server

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
)

// HTTP handler for server
func mutatingWebhookHandler(rw http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "%s", err)
	}

	// TODO: mutate the request
	// mutated, err := m.Mutate(body)
	mutated := body
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "%s", err)
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(mutated)
}

// HTTP root handler
func rootHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "hello %q", html.EscapeString(req.URL.Path))
}
