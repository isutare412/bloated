package pkgerror_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/isutare412/bloated/api/pkg/pkgerror"
)

var errOrigin = errors.New("some complex error occurred")

var _ = Describe("Known", func() {
	It("unwraps origin error", func() {
		err := FireError()
		Expect(err).To(Equal(errOrigin))

		err = pkgerror.Known{Origin: err}
		Expect(err).NotTo(Equal(errOrigin))
		Expect(errors.Is(err, errOrigin)).To(BeTrue())

		kerr := pkgerror.AsKnown(err)
		Expect(kerr).NotTo(BeNil())
		Expect(kerr.Origin).To(Equal(errOrigin))
	})

	It("does not cast to Known if it is not known error", func() {
		err := FireError()
		Expect(err).To(Equal(errOrigin))

		kerr := pkgerror.AsKnown(err)
		Expect(kerr).To(BeNil())
	})
})

func FireError() error {
	return errOrigin
}
