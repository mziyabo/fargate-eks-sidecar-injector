package server

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"

	injector "github.com/mziyabo/fargate-eks-sidecar-injector/m/v2/pkg/fargateInjector"
	"k8s.io/api/admission/v1beta1"
)

// HTTP handler for server
func mutatingWebhookHandler(rw http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "%s", err)
	}

	ar := v1beta1.AdmissionReview{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		panic(err)
	}

	response, err := injector.Mutate(ar)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "%s", err)
	}

	ar.Response = response
	responseBody, err := json.Marshal(ar)
	if err != nil {
		panic(err)
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(responseBody)
}

// HTTP root handler
func rootHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "%q", html.EscapeString(req.URL.Path))
}
