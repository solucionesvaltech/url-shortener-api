package details

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestURL(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "details adapter suite")
}
