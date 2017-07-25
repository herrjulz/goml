package goml_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoml(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goml Suite")
}
