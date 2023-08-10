package config_test

import (
	"os"
	"path"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/isutare412/bloated/api/pkg/config"
	"github.com/isutare412/bloated/api/pkg/validation"
)

var _ = Describe("Loader", func() {
	It("overwrites config using environment variables", func() {
		os.Setenv("APP_POSTGRES_DATABASE", "test")

		validator, err := validation.NewValidator()
		Expect(err).NotTo(HaveOccurred())

		cfg, err := config.LoadValidated(path.Join("..", "..", "config.yaml"), validator)
		Expect(err).NotTo(HaveOccurred())
		Expect(cfg.Postgres.Database).To(BeEquivalentTo("test"))
	})
})
