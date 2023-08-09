package postgres

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Helper", func() {
	Context("pointerIfNotZero", func() {
		It("works on int", func() {
			Expect(pointerIfNotZero(123)).NotTo(BeNil())
			Expect(pointerIfNotZero(0)).To(BeNil())
		})

		It("works on string", func() {
			Expect(pointerIfNotZero("123")).NotTo(BeNil())
			Expect(pointerIfNotZero("")).To(BeNil())
		})

		It("works on struct", func() {
			type user struct {
				name string
			}

			Expect(pointerIfNotZero(user{name: "tester"})).NotTo(BeNil())
			Expect(pointerIfNotZero(user{})).To(BeNil())
		})

		It("works on pointer", func() {
			num := 123
			Expect(pointerIfNotZero[*int](&num)).NotTo(BeNil())
			Expect(pointerIfNotZero[*int](nil)).To(BeNil())
		})
	})
})
