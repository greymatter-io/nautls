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

package urls

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"testing"

	"github.com/greymatter-io/nautls/internal/temporary"
	"github.com/greymatter-io/nautls/internal/tests"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReadFile(test *testing.T) {

	Convey("When .ReadFile is invoked", test, func() {

		Convey("with no scheme", func() {

			expectedContent := tests.MustGenerateBytes(test)
			actualContent, err := temporary.WithFile(expectedContent, 0777, func(path string) (interface{}, error) {
				resource := tests.MustRelativePath(path, test)
				return ReadFile(resource)
			})

			Convey("it returns the content", func() {
				So(actualContent, ShouldResemble, expectedContent)
			})

			Convey("it returns a nil error", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("with a base64 scheme", func() {

			expectedContent := tests.MustGenerateBytes(test)
			resource := fmt.Sprintf("base64:///%s", url.PathEscape(base64.StdEncoding.EncodeToString(expectedContent)))
			actualContent, err := ReadFile(resource)

			Convey("it returns the content", func() {
				So(actualContent, ShouldResemble, expectedContent)
			})

			Convey("it returns a nil error", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("with a file scheme", func() {

			expectedContent := tests.MustGenerateBytes(test)
			actualContent, err := temporary.WithFile(expectedContent, 0777, func(path string) (interface{}, error) {
				resource := fmt.Sprintf("file://%s", path)
				return ReadFile(resource)
			})

			Convey("it returns the content", func() {
				So(actualContent, ShouldResemble, expectedContent)
			})

			Convey("it returns a nil error", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
