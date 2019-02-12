package main

import (
	"crypto/tls"
	"log"
	"net/http"
	uiza "uiza-api-wrapper"
	"uiza-api-wrapper/entity"
	_ "uiza-api-wrapper/testing"
)

func init() {
	uiza.Key = "uap-a2aaa7b2aea746ec89e67ad2f8f9ebbf-fdf5bdca"
	// var port = "80"
	// Configure a backend for stripe-mock and set it for both the API and
	// Uploads (unlike the real Stripe API, stripe-mock supports both these
	// backends).
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	httpClient := &http.Client{
		Transport: transport,
	}

	uizaMockBackend := uiza.GetBackendWithConfig(
		uiza.APIBackend,
		&uiza.BackendConfig{
			URL:        "https://apiwrapper.uiza.co:",
			HTTPClient: httpClient,
			Logger:     uiza.Logger,
			LogLevel:   3,
		},
	)
	uiza.SetBackend(uiza.APIBackend, uizaMockBackend)
	uiza.SetBackend(uiza.UploadsBackend, uizaMockBackend)
}

func main() {
	params := &uiza.EntitySpecParams{ID: uiza.String("650131dc-c024-40e5-bde1-8bdb0cf898c6")}
	ch, _ := entity.Retrieve(params)
	log.Printf("%v\n", ch.RequestID)
}
