// Copyright 2023 greymatter.io Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package builders

import (
	"crypto/tls"
	"crypto/x509"

	"github.com/greymatter-io/nautls/internal/urls"
	"github.com/pkg/errors"
)

// BuildCertificatePool provides a utility function for creating a certificate pool from an array of resources. Note that if
// the array of URLs is empty the system certificates will be used.
func BuildCertificatePool(certificateResources []string) (*x509.CertPool, error) {

	if len(certificateResources) == 0 {
		return x509.SystemCertPool()
	}

	pool := x509.NewCertPool()
	for _, certificateResource := range certificateResources {

		bytes, err := readResource(certificateResource)
		if err != nil {
			return nil, errors.Wrapf(err, "error reading certificate [%s]", certificateResource)
		}

		if !pool.AppendCertsFromPEM(bytes) {
			return nil, errors.Wrapf(err, "error appending certificate from [%s]", certificateResource)
		}
	}

	return pool, nil
}

// BuildCertificates provides a utility function for loading a certificate from certificate and key resources.
func BuildCertificates(certificateResource string, keyResource string) ([]tls.Certificate, error) {

	certificates := []tls.Certificate{}

	if (certificateResource == "") && (keyResource == "") {
		return certificates, nil
	}

	certificate, err := readKeyPair(certificateResource, keyResource)
	if err != nil {
		return nil, errors.Wrapf(err, "error loading key pair from [%s] and [%s]", certificateResource, keyResource)
	}

	return append(certificates, certificate), nil
}

// readKeyPair reads an X.509 key pair from certificate and key resources.
func readKeyPair(certificateResource string, keyResource string) (tls.Certificate, error) {

	certificateBytes, err := readResource(certificateResource)
	if err != nil {
		return tls.Certificate{}, errors.Wrapf(err, "error reading certificate [%s]", certificateResource)
	}

	keyBytes, err := readResource(keyResource)
	if err != nil {
		return tls.Certificate{}, errors.Wrapf(err, "error reading key [%s]", keyResource)
	}

	return tls.X509KeyPair(certificateBytes, keyBytes)
}

// readResource reads a resource URL into a byte array.
func readResource(resource string) ([]byte, error) {

	resourceBytes, err := urls.ReadFile(resource)
	if err != nil {
		return nil, errors.Wrapf(err, "error reading resource from [%s]", resource)
	}

	return resourceBytes, nil
}
