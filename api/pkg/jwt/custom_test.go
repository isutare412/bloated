package jwt_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/isutare412/bloated/api/pkg/core/model"
	"github.com/isutare412/bloated/api/pkg/jwt"
)

const (
	testPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA8D+zWMKfmKKmYyxsrDVFAsJ5uTk7aBtSAWYD9PX73gILIUsC
35qnKEd1I1S9l7YUoQ1vtpNR1FtvdYMWw6CeGx8p2hGjcB/Fj7xVV93wxcHBkacx
cfrCesDh4WGDZq5i9veF07wxdWFbNGZP0Npex7wiJBLvaIQCJAQbxCCFHf+oN70T
mkoCee7MlYMlkw01tkuW6dyOPRIUQtGUZ3OBaeUHA2IcQrMSScG7PTS1U8AlvGt6
qU05xkdrZRhXzS9zxIPuLR/ScVaAShbqlB/81b72WoTPsvwZZfm6vg1l612r+Osb
wnmgfeLppEl11NGQ106FWWU4WcAHM0vPtZDBFwIDAQABAoIBAHrpjB74C2Cyrf+F
BeAgrLrZth6+GheMCqtufs2/X2lYkEsrLkApxiVEUbiOrSTF4c33qtS0kCPd16s+
MtJJBTqI+gd4CK7fglqkFuGKSZlTJG4ZJKHUkdTtg2KkWe6Zf0YsoooN+Ru9gETR
pzoJzn3PUYQ1L8i/6Lx2YaI4pQIQAUILCI6zVHY+1WIS2fpjwtxx95WiyLppxkos
WJhl6iVEOzNzDDjaBY6rCU1KjVz/pFFkiQ03pIzC/mAjzNiO1PKgiys9WEeWglYP
/fGPiuu6ZHRiWOcRaqGOHSw+5HNEbeLjeKvKtwLtaevN5jfycn4CeGvO8nWDW6Ao
OY+BvlECgYEA+jnmxBppZVt95ZFmUDXUNuaoQwOeKVTKSxco2a7ZtlV15mNiS3Z9
vcAKbWwoAOaGFf5r1ApHrKqndrfiOEGrdIK8CKgB4QgE+AAu04CqRYfl4dAxsdQK
BycmS1VeWSux34QkJRw5RYqfAv5Ohsx29Oa5NF0vp+EnMn5NO5sFU28CgYEA9crc
yclf4C6UXO88extZKRPvnkTi+bFPl1vomXThZFEIGjTMp582W5xJs3/btHAPNd88
VuteQi5H851Cm1q2yE+JIw22zZKG9NGFKy7QHtula1DdIrs32AcBzkQC8o8PDgo5
uzMrxbc5/9HkOf0GoF1cqur5TD1ViIjVbVOfeNkCgYAia4qR2Sw8VAyKYqjXScBK
WQo2Ra4g81fsYmAtKX8rNLBGI8+mn8MGMf99M86A6TTym4/LmCpp1wO4Me09XdMu
8Lja2/l+MfM1Bhlq4J0LeXkmydO8KXXDrB/5ZuHWlEsgevXZ4PsD1OCwsbitlDyc
zYwWZuciSbdxm6wWEen2owKBgQDUTBoT+GllWxYuhzy4IFEMl/mZvGpHvQy/8VSg
Z0Hewda7u+sgxPXQftdxwPfli/y3TU/yy20owIzJMIW9ZccGkRwkOM4yFWOXxfi9
6bs9S/4/CSNXwlljr/mxTTE2jLY2LELdHD+skKv5+DKmm5PDo6BMyJP3c+qS+Y1O
1rGhUQKBgQCV9baePSGJ2B8qW+pd4GlZMrxMnkOavqUKvOwKXY5VdhnsLAfLyXrz
3PJhpBGarcua7XUnykbxZkAtvSgA+Lg35Ea0Ki73cVPoBjK8HBwmNyd/ULSdmfqB
xU2rRsCfjmznvh5bMZ50aNHTFSs3Wl+xX6OqWmrPcnTryyHddMBh7g==
-----END RSA PRIVATE KEY-----`

	testPublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA8D+zWMKfmKKmYyxsrDVF
AsJ5uTk7aBtSAWYD9PX73gILIUsC35qnKEd1I1S9l7YUoQ1vtpNR1FtvdYMWw6Ce
Gx8p2hGjcB/Fj7xVV93wxcHBkacxcfrCesDh4WGDZq5i9veF07wxdWFbNGZP0Npe
x7wiJBLvaIQCJAQbxCCFHf+oN70TmkoCee7MlYMlkw01tkuW6dyOPRIUQtGUZ3OB
aeUHA2IcQrMSScG7PTS1U8AlvGt6qU05xkdrZRhXzS9zxIPuLR/ScVaAShbqlB/8
1b72WoTPsvwZZfm6vg1l612r+OsbwnmgfeLppEl11NGQ106FWWU4WcAHM0vPtZDB
FwIDAQAB
-----END PUBLIC KEY-----`
)

var _ = Describe("Custom", func() {
	var customClient *jwt.CustomClient

	BeforeEach(func() {
		client, err := jwt.NewCustomClient(jwt.CustomClientConfig{
			TokenTTL:   time.Hour,
			PrivateKey: []byte(testPrivateKey),
			PublicKey:  []byte(testPublicKey),
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
