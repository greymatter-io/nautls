package tests

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"
)

// MustServe returns a server listening on the provided address with the provided TLS configuration or fails the test.
func MustServe(test *testing.T, configuration *tls.Config) (string, io.Closer) {

	var err error
	var listener net.Listener

	if configuration == nil {
		listener, err = net.Listen("tcp", "localhost:0")
		if err != nil {
			test.Errorf("unable to generate net listener")
		}
	} else {
		listener, err = tls.Listen("tcp", "localhost:0", configuration)
		if err != nil {
			test.Errorf("unable to generate tls listener")
		}
	}

	port := listener.Addr().(*net.TCPAddr).Port
	server := &http.Server{Handler: http.NotFoundHandler()}

	go server.Serve(listener)

	return fmt.Sprintf("localhost:%d", port), server
}
