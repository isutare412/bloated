package config_test

import (
	"os"
	"path"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/isutare412/bloated/api/pkg/config"
)

var _ = Describe("Loader", func() {
	It("overwrites config using environment variables", func() {
		os.Setenv("APP_LOG_FORMAT", "test")

		cfg, err := config.LoadValidated(path.Join("..", "..", "config.yaml"))
		Expect(err).To(HaveOccurred())
		Expect(cfg.Log.Format).To(BeEquivalentTo("test"))
	})
})
