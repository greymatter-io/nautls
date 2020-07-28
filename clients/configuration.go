package clients

import (
	"crypto/tls"
	"net/http"

	"github.com/greymatter-io/nautls/builders"
	"github.com/pkg/errors"
)

// Configuration provides a serializable representation of an tls.Config configuration supporting plaintext, TLS and
// mTLS communications.
type Configuration struct {

	// Authorities defines the trusted certificate authorities. The values must be URLs that point to the location of
	// PEM encoded certificates.
	//
	// Note that in addition to those schemes supported by [getter](https://godoc.org/github.com/hashicorp/go-getter) a
	// "base64" scheme is supported for providing the PEM encoded certifiate in the path of the URL directly. This is
	//  most applicable when the certificate data must be provided via an environement variable.
	Authorities []string `json:"authorities" mapstructure:"authorities" yaml:"authorities"`

	// Certificate defines the client certificate used for mTLS connections. The value must be a URL that points to the
	// location of a PEM encoded certificate.
	//
	// Note that in addition to those schemes supported by [getter](https://godoc.org/github.com/hashicorp/go-getter) a
	// "base64" scheme is supported for providing the PEM encoded certifiate in the path of the URL directly. This is
	// most applicable when the certificate data must be provided via an environement variable.
	Certificate string `json:"certificate" mapstructure:"certificate" yaml:"certificate"`

	// Key defines the client key used for mTLS connections. The value must be a URL that points to the location of a
	// PEM encoded key.
	//
	// Note that in addition to those schemes supported by [getter](https://godoc.org/github.com/hashicorp/go-getter) a
	// "base64" scheme is supported for providing the PEM encoded certifiate in the path of the URL directly. This is
	// most applicable when the certificate data must be provided via an environement variable.
	Key string `json:"key" mapstructure:"key" yaml:"key"`

	// Server defines the server name used for certificate verification.
	Server string `json:"server" mapstructure:"server" yaml:"server"`
}

// HTTP returns an http.Client from the configuration. Note that invoking this method on a nil instance is not an error
// and returns the value if http.DefaultClient.
func (c *Configuration) HTTP() (*http.Client, error) {

	configuration, err := c.TLS()
	if err != nil {
		return nil, errors.Wrap(err, "error building tls configuration for client")
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: configuration,
		},
	}

	return client, nil
}

// TLS returns a tls.Config instance from the configuration. Note that invoking this method on a nil instance is not an
// error and returns nil.
func (c *Configuration) TLS() (*tls.Config, error) {

	if c == nil {
		return nil, nil
	}

	pool, err := builders.BuildCertificatePool(c.Authorities)
	if err != nil {
		return nil, errors.Wrap(err, "error building certificate authority pool")
	}

	certificates, err := builders.BuildCertificates(c.Certificate, c.Key)
	if err != nil {
		return nil, errors.Wrap(err, "error building certificates")
	}

	configuration := &tls.Config{
		Certificates: certificates,
		RootCAs:      pool,
		ServerName:   c.Server,
	}

	return configuration, nil

}
