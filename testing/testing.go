package testing

import (
	uiza "api-wrapper-go"
	"crypto/tls"
	"net/http"
)

func init() {
	uiza.Key = ""

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
			URL:        "https://uiza.io",
			HTTPClient: httpClient,
			Logger:     uiza.Logger,
			LogLevel:   3,
		},
	)
	uiza.SetBackend(uiza.APIBackend, uizaMockBackend)
	uiza.SetBackend(uiza.UploadsBackend, uizaMockBackend)
}
