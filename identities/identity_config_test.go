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

package identities

import (
	"encoding/json"
	"testing"

	"github.com/greymatter-io/nautls/internal/tests"
	"gopkg.in/yaml.v2"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIdentityConfig(t *testing.T) {

	Convey("When IdentityConfig", t, func() {

		Convey(".Build is invoked", func() {

			config := &IdentityConfig{
				Authorities: "./testdata/multiple.crt",
				Certificate: "./testdata/single.crt",
				Key:         "./testdata/single.key",
			}

			Convey("with empty authorities", func() {

				config.Authorities = "./testdata/empty.crt"
				identity, err := config.Build()

				Convey("it should return a non-nil identity", func() {
					So(identity, ShouldNotBeNil)
				})

				Convey("it should return a nil error", func() {
					So(err, ShouldBeNil)
				})
			})

			Convey("with invalid authorities", func() {

				config.Authorities = "./testdata/invalid.crt"
				identity, err := config.Build()

				Convey("it should return a nil identity", func() {
					So(identity, ShouldBeNil)
				})

				Convey("it should return a non-nil error", func() {
					So(err, ShouldNotBeNil)
				})
			})

			Convey("with an empty certificate", func() {

				config.Certificate = "./testdata/empty.crt"
				identity, err := config.Build()

				Convey("it should return a nil identity", func() {
					So(identity, ShouldBeNil)
				})

				Convey("it should return a non-nil error", func() {
					So(err, ShouldNotBeNil)
				})
			})

			Convey("with an invalid certificate", func() {

				config.Certificate = "./testdata/invalid.crt"
				identity, err := config.Build()

				Convey("it should return a nil identity", func() {
					So(identity, ShouldBeNil)
				})

				Convey("it should return a non-nil error", func() {
					So(err, ShouldNotBeNil)
				})
			})

			Convey("with multiple certificates", func() {

				config.Certificate = "./testdata/multiple.crt"
				identity, err := config.Build()

				Convey("it should return a nil identity", func() {
					So(identity, ShouldBeNil)
				})

				Convey("it should return a non-nil error", func() {
					So(err, ShouldNotBeNil)
				})
			})

			Convey("with an empty key", func() {

				config.Key = "./testdata/empty.key"
				identity, err := config.Build()

				Convey("it should return a nil identity", func() {
					So(identity, ShouldBeNil)
				})

				Convey("it should return a non-nil error", func() {
					So(err, ShouldNotBeNil)
				})
			})

			Convey("with an invalid key", func() {

				config.Key = "./testdata/invalid.key"
				identity, err := config.Build()

				Convey("it should return a nil identity", func() {
					So(identity, ShouldBeNil)
				})

				Convey("it should return a non-nil error", func() {
					So(err, ShouldNotBeNil)
				})
			})

			Convey("with multiple keys", func() {

				config.Key = "./testdata/multiple.key"
				identity, err := config.Build()

				Convey("it should return a nil identity", func() {
					So(identity, ShouldBeNil)
				})

				Convey("it should return a non-nil error", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})

		Convey(" is deserialized", func() {

			var actual IdentityConfig

			expected := map[string]string{
				"authorities": "authorities",
				"certificate": "certificate",
				"key":         "key",
			}

			Convey("from JSON", func() {

				err := json.Unmarshal(tests.MustTemplate("testdata/config.json.tpl", expected, t), &actual)

				Convey("it should populate the authorities", func() {
					So(actual.Authorities, ShouldEqual, expected["authorities"])
				})

				Convey("it should populate the certificate", func() {
					So(actual.Certificate, ShouldEqual, expected["certificate"])
				})

				Convey("it should populate the key", func() {
					So(actual.Key, ShouldEqual, expected["key"])
				})

				Convey("it should return a nil error", func() {
					So(err, ShouldBeNil)
				})
			})

			Convey("from YAML", func() {

				err := yaml.Unmarshal(tests.MustTemplate("testdata/config.yaml.tpl", expected, t), &actual)

				Convey("it should populate the authorities", func() {
					So(actual.Authorities, ShouldEqual, expected["authorities"])
				})

				Convey("it should populate the certificate", func() {
					So(actual.Certificate, ShouldEqual, expected["certificate"])
				})

				Convey("it should populate the key", func() {
					So(actual.Key, ShouldEqual, expected["key"])
				})

				Convey("it should return a nil error", func() {
					So(err, ShouldBeNil)
				})
			})
		})
	})
}
