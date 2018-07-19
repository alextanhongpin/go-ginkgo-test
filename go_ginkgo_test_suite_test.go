package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoGinkgoTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoGinkgoTest Suite")
}
