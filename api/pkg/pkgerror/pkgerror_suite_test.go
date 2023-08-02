package pkgerror_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPkgerror(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pkgerror Suite")
}
