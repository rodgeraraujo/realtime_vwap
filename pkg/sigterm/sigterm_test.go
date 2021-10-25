package sigterm_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rodgeraraujo/vwap/pkg/sigterm"
)

func TestSigterm(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sigterm Suite")
}

var _ = Describe("sigterm.HnderSigterm()", func() {
	Context("Sigterm handler", func() {
		It("should exit the program", func() {
			sigterm.HnderSigterm()
			fmt.Println("Sigterm handler called")
		})
	})
})
