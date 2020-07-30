// Copyright 2019 Decipher Technology Studios
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

// Package encoding provies helper methods for PEM encoding x509 certificates and RSA private keys.
package encoding

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func PEMEncodeCertificate(certificate *x509.Certificate) []byte {
	if certificate == nil {
		return []byte{}
	}

	return pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certificate.Raw,
	})
}

func PEMEncodeCertificates(certificates []*x509.Certificate) [][]byte {
	pems := make([][]byte, len(certificates))
	for i, cert := range certificates {
		pems[i] = PEMEncodeCertificate(cert)
	}
	return pems
}

func PEMEncodeKey(key *rsa.PrivateKey) []byte {
	if key == nil {
		return []byte{}
	}

	return pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})
}

func PEMEncodeKeys(keys []*rsa.PrivateKey) [][]byte {
	pems := make([][]byte, len(keys))
	for i, key := range keys {
		pems[i] = PEMEncodeKey(key)
	}
	return pems
}
