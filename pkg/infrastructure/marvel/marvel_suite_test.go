package marvel_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMarvel(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Marvel Client Suite")
}
