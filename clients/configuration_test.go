package clients

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"

	"github.com/greymatter-io/nautls/internal/tests"
	"github.com/greymatter-io/nautls/servers"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	AuthorityCertificate = "./testdata/authority.crt"
	ClientCertificate    = "./testdata/client.crt"
	ClientKey            = "./testdata/client.key"
	ServerCertificate    = "./testdata/server.crt"
	ServerKey            = "./testdata/server.key"
)

// shouldBeClient returns function that validates whether a client is compatible with a server
func shouldBeClient(t *testing.T, scheme string, configuration *servers.Configuration) func(interface{}, ...interface{}) string {

	return func(actual interface{}, expected ...interface{}) string {

		client, ok := actual.(*http.Client)
		if !ok {
			return "expected http client but was not http client"
		}

		config, err := configuration.TLS()
		if err != nil {
			t.Errorf("error creating server tls configuration [%s]", err)
		}

		address, server := tests.MustServe(t, config)
		defer server.Close()

		response, err := client.Get(fmt.Sprintf("%s://%s", scheme, address))
		if err != nil {
			return fmt.Sprintf("expected nil error but was [%s]", err.Error())
		}

		if response.StatusCode != http.StatusNotFound {
			return fmt.Sprintf("expected 404 status code but was [%d]", response.StatusCode)
		}

		return ""
	}

}

// shouldNotBeClient returns function that validates whether a client is not compatible with a server
func shouldNotBeClient(t *testing.T, scheme string, configuration *servers.Configuration) func(interface{}, ...interface{}) string {

	return func(actual interface{}, expected ...interface{}) string {

		client, ok := actual.(*http.Client)
		if !ok {
			return "expected http client but was not http client"
		}

		config, err := configuration.TLS()
		if err != nil {
			t.Errorf("error creating server tls configuration [%s]", err)
		}

		address, server := tests.MustServe(t, config)
		defer server.Close()

		response, err := client.Get(fmt.Sprintf("%s://%s", scheme, address))
		if err != nil {
			return ""
		}

		if response.StatusCode != http.StatusNotFound {
			return ""
		}

		return "expected incompatable client but was valid"
	}

}

func TestConfiguration(t *testing.T) {

	shouldBePlaintextClient := shouldBeClient(t, "http", nil)

	shouldBeTLSClient := shouldBeClient(t, "https", &servers.Configuration{
		Certificate: ServerCertificate,
		Key:         ServerKey,
	})

	shouldBeMTLSClient := shouldBeClient(t, "https", &servers.Configuration{
		Authorities: []string{AuthorityCertificate},
		Certificate: ServerCertificate, Key: ServerKey,
	})

	shouldNotBeTLSClient := shouldNotBeClient(t, "https", &servers.Configuration{
		Certificate: ServerCertificate,
		Key:         ServerKey,
	})

	shouldNotBeMTLSClient := shouldNotBeClient(t, "https", &servers.Configuration{
		Authorities:    []string{AuthorityCertificate},
		Certificate:    ServerCertificate,
		Key:            ServerKey,
		Authentication: servers.Authentication(tls.RequireAndVerifyClientCert),
	})

	Convey("When Configuration", t, func() {

		var configuration *Configuration

		Convey(".HTTP is invoked", func() {

			Convey("and the configuration is nil", func() {

				client, err := configuration.HTTP()

				Convey("it returns a nil error", func() {
					So(err, ShouldBeNil)
				})

				Convey("it returns a plaintext client", func() {
					So(client, shouldBePlaintextClient)
				})

				Convey("it returns an invalid TLS client", func() {
					So(client, shouldNotBeTLSClient)
				})

				Convey("it returns an invalid mTLS client", func() {
					So(client, shouldNotBeMTLSClient)
				})
			})

			Convey("and the configuration is zero", func() {

				configuration := &Configuration{}

				client, err := configuration.HTTP()

				Convey("it returns a nil error", func() {
					So(err, ShouldBeNil)
				})

				Convey("it returns a plaintext client", func() {
					So(client, shouldBePlaintextClient)
				})

				Convey("it returns an invalid TLS client", func() {
					So(client, shouldNotBeTLSClient)
				})

				Convey("it returns an invalid mTLS client", func() {
					So(client, shouldNotBeMTLSClient)
				})
			})

			Convey("and the configuration is TLS", func() {

				configuration := &Configuration{
					Authorities: []string{AuthorityCertificate},
				}

				client, err := configuration.HTTP()

				Convey("it returns a nil error", func() {
					So(err, ShouldBeNil)
				})

				Convey("it returns a valid plaintext client", func() {
					So(client, shouldBePlaintextClient)
				})

				Convey("it returns a valid TLS client", func() {
					So(client, shouldBeTLSClient)
				})

				Convey("it returns an invalid mTLS client", func() {
					So(client, shouldNotBeMTLSClient)
				})
			})

			Convey("and the configuration is mTLS", func() {

				configuration := &Configuration{
					Authorities: []string{AuthorityCertificate},
					Certificate: ClientCertificate,
					Key:         ClientKey,
				}

				client, err := configuration.HTTP()

				Convey("it returns a nil error", func() {
					So(err, ShouldBeNil)
				})

				Convey("it returns a valid plaintext client", func() {
					So(client, shouldBePlaintextClient)
				})

				Convey("it returns a valid TLS client", func() {
					So(client, shouldBeTLSClient)
				})

				Convey("it returns a valid mTLS client", func() {
					So(client, shouldBeMTLSClient)
				})
			})
		})
	})
}
