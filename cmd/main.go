package main

import (
	uiza "api-wrapper-go"
	"api-wrapper-go/entity"
	_ "api-wrapper-go/testing"
	"crypto/tls"
	"log"
	"net/http"
)

func init() {
	uiza.Key = "uap-a2aaa7b2aea746ec89e67ad2f8f9ebbf-fdf5bdca"

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
			URL:        "https://apiwrapper.uiza.co",
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
	response, _ := entity.Retrieve(params)
	log.Printf("%s\n", response)
}
