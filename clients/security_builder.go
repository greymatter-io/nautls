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

package clients

import "crypto/tls"

// SecurityBuilder provides an builder for client tls.Config instances.
//
// Deprecated: SecurityBuilder should no longer be used and implementations should move to ConfigurationBuilder.
type SecurityBuilder struct {
	config SecurityConfig
}

// NewSecurityBuilder returns a new instance of the SecurityBuilder structure.
func NewSecurityBuilder() *SecurityBuilder {
	return &SecurityBuilder{}
}

// Build creates a tls.Config from the SecurityBuilder.
func (b *SecurityBuilder) Build() (*tls.Config, error) {
	return b.config.Build()
}

// WithAuthorities sets the certificate authorities trusted by the built tls.Config. The values must be URLs that point
// to the locations of PEM encoded certificates.
//
// Note that in addition to those schemes supported by [getter](https://godoc.org/github.com/hashicorp/go-getter) a
// "base64" scheme is supported for providing the PEM encoded certifiate in the path of the URL directly. This is most
// applicable when the certificate data must be provided via an environement variable.
func (b *SecurityBuilder) WithAuthorities(authorities []string) *SecurityBuilder {
	b.config.Authorities = authorities
	return b
}

// WithCertificate sets the client certificate used for mTLS connections. The value must be a URL that points to the
// location of a PEM encoded certificate.
//
// Note that in addition to those schemes supported by [getter](https://godoc.org/github.com/hashicorp/go-getter) a
// "base64" scheme is supported for providing the PEM encoded certifiate in the path of the URL directly. This is most
// applicable when the certificate data must be provided via an environement variable.
func (b *SecurityBuilder) WithCertificate(certificate string) *SecurityBuilder {
	b.config.Certificate = certificate
	return b
}

// WithKey sets the client key used for mTLS connections. The value must be a URL that points to the location of a PEM
// encoded key.
//
// Note that in addition to those schemes supported by [getter](https://godoc.org/github.com/hashicorp/go-getter) a
// "base64" scheme is supported for providing the PEM encoded certifiate in the path of the URL directly. This is most
// applicable when the certificate data must be provided via an environement variable.
func (b *SecurityBuilder) WithKey(key string) *SecurityBuilder {
	b.config.Key = key
	return b
}

// WithServer sets the server name used for certificate verification.
func (b *SecurityBuilder) WithServer(server string) *SecurityBuilder {
	b.config.Server = server
	return b
}
