package jwt_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/isutare412/bloated/api/pkg/core/model"
	"github.com/isutare412/bloated/api/pkg/jwt"
)

var _ = Describe("Custom", func() {
	var customClient *jwt.CustomClient

	BeforeEach(func() {
		client, err := jwt.NewCustomClient(jwt.CustomClientConfig{
			TokenTTL:       time.Hour,
			PrivateKeyPath: "test.key",
			PublicKeyPath:  "test.key.pub",
		})
		Expect(err).NotTo(HaveOccurred())

		customClient = client
	})

	It("verifies signed token", func() {
		var (
			givenToken = model.CustomToken{
				Name:       "Tester Lee",
				GivenName:  "Tester",
				FamilyName: "Lee",
				Picture:    "picture-url",
				Email:      "foo@gmail.com",
			}
		)

		tokenString, err := customClient.SignCustomToken(givenToken)
		Expect(err).NotTo(HaveOccurred())

		verifiedToken, err := customClient.VerifyCustomToken(tokenString)
		Expect(err).NotTo(HaveOccurred())
		Expect(verifiedToken).To(Equal(givenToken))
	})
})
