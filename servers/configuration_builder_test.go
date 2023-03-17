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

package servers

import (
	"testing"

	"github.com/greymatter-io/nautls/internal/tests"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConfigurationBuilder(t *testing.T) {

	Convey("When ConfigurationBuilder", t, func() {

		builder := NewConfigurationBuilder()

		Convey(".Build() is invoked", func() {

			config := builder.Build()

			Convey("it returns the configuration", func() {
				So(config, ShouldNotBeZeroValue)
			})
		})

		Convey(".WithAuthority is invoked", func() {

			authorities := tests.MustGenerateStrings(t)

			builder.WithAuthorities(authorities)

			Convey("it sets the authorities", func() {
				So(builder.Authorities, ShouldResemble, authorities)
			})
		})

		Convey(".WithCertificate is invoked", func() {

			certificate := tests.MustGenerateString(t)

			builder.WithCertificate(certificate)

			Convey("it sets the certificate", func() {
				So(builder.Certificate, ShouldEqual, certificate)
			})
		})

		Convey(".WithKey is invoked", func() {

			key := tests.MustGenerateString(t)

			builder.WithKey(key)

			Convey("it sets the key", func() {
				So(builder.Key, ShouldEqual, key)
			})
		})

		Convey(".WithAuthentication is invoked", func() {

			authentication := MustGenerateAuthentication(t)

			builder.WithAuthentication(authentication)

			Convey("it sets the authentication", func() {
				So(builder.Authentication, ShouldEqual, authentication)
			})
		})
	})
}
