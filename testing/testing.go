package testing

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	uiza "api-wrapper-go"
	"api-wrapper-go/form"

	"golang.org/x/net/http2"
)

func init() {
	// Enable strict mode on form encoding so that we'll panic if any kind of
	// malformed param struct is detected
	form.Strict = true

	port := os.Getenv("AZUI_MOCK_PORT")
	if port == "" {
		port = "12112"
	}

	// uiza-mock's certificate for localhost is self-signed so configure a
	// specialized client that skips the certificate authority check.
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	// Go can often enable HTTP/2 automatically if it's supported, but
	// confusingly, if you set `TLSClientConfig`, it disables it and you have
	// to explicitly invoke http2's `ConfigureTransport` to get it back.
	//
	// See the incorrectly closed bug report here:
	//
	//     https://github.com/golang/go/issues/20645
	//
	err := http2.ConfigureTransport(transport)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize HTTP/2 transport: %v\n", err)
		os.Exit(1)
	}

	httpClient := &http.Client{
		Transport: transport,
	}

	// resp, err := httpClient.Get("https://localhost:" + port)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Couldn't reach uiza-mock at `localhost:%s` (%v). Is "+
	// 		"it running? Please see README for setup instructions.\n", port, err)
	// 	os.Exit(1)
	// }
	// version := resp.Header.Get("uiza-Mock-Version")
	// if version != "master" && compareVersions(version, MockMinimumVersion) > 0 {
	// 	fmt.Fprintf(os.Stderr, "Your version of uiza-mock (%s) is too old. The "+
	// 		"minimum version to run this test suite is %s. Please see its "+
	// 		"repository for upgrade instructions.\n", version, MockMinimumVersion)
	// 	os.Exit(1)
	// }

	uiza.Key = "uap-a2aaa7b2aea746ec89e67ad2f8f9ebbf-fdf5bdca"

	// Configure a backend for uiza-mock and set it for both the API and
	// Uploads (unlike the real uiza API, uiza-mock supports both these
	// backends).
	uizaMockBackend := uiza.GetBackendWithConfig(
		uiza.APIBackend,
		&uiza.BackendConfig{
			URL:        "https://localhost:" + port,
			HTTPClient: httpClient,
			Logger:     uiza.Logger,
		},
	)
	uiza.SetBackend(uiza.APIBackend, uizaMockBackend)
	uiza.SetBackend(uiza.UploadsBackend, uizaMockBackend)
}
