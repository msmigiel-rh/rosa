/*
Copyright (c) 2025 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package url

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestURLValidation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "URL Validation Suite")
}

var _ = Describe("ValidateURLCredentials", func() {
	Context("when URL has no scheme separator", func() {
		It("returns nil for URL without scheme", func() {
			err := ValidateURLCredentials("example.com")
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns nil for URL with partial scheme", func() {
			err := ValidateURLCredentials("http:/example.com")
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("when URL has no credentials", func() {
		It("returns nil for URL without @", func() {
			err := ValidateURLCredentials("http://example.com")
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns nil for URL with port but no credentials", func() {
			err := ValidateURLCredentials("http://example.com:8080")
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("when URL has valid credentials", func() {
		It("returns nil for URL with username only", func() {
			err := ValidateURLCredentials("http://user@example.com")
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns nil for URL with username and password", func() {
			err := ValidateURLCredentials("http://user:pass@example.com")
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns nil for URL with empty password", func() {
			err := ValidateURLCredentials("http://user:@example.com")
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("when username contains invalid characters", func() {
		DescribeTable("returns error for invalid username character",
			func(url string, expectedChar rune) {
				err := ValidateURLCredentials(url)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("username contains invalid character '" + string(expectedChar) + "'"))
			},
			Entry("slash in username", "http://us/er:pass@example.com", '/'),
			Entry("question mark in username", "http://us?er:pass@example.com", '?'),
			Entry("hash in username", "http://us#er:pass@example.com", '#'),
		)
	})

	Context("when password contains invalid characters", func() {
		DescribeTable("returns error for invalid password character",
			func(url string, expectedChar rune) {
				err := ValidateURLCredentials(url)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("password contains invalid character '" + string(expectedChar) + "'"))
			},
			Entry("slash in password", "http://user:pa/ss@example.com", '/'),
			Entry("question mark in password", "http://user:pa?ss@example.com", '?'),
			Entry("hash in password", "http://user:pa#ss@example.com", '#'),
			Entry("bracket in password", "http://user:pa[ss@example.com", '['),
		)
	})

	Context("when URL has multiple @ signs", func() {
		It("returns error indicating @ in password", func() {
			err := ValidateURLCredentials("http://user:p@ss@example.com")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("password contains invalid character '@'"))
		})
	})
})
