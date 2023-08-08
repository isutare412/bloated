package jwt_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/isutare412/bloated/api/pkg/log"
)

func TestJwt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Jwt Suite")
}

var _ = BeforeSuite(func() {
	log.AdaptGinkgo()
})
