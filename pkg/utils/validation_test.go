package utils_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rodgeraraujo/vwap/pkg/utils"
)

func TestHelpers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Helpers Suite")
}

var _ = Describe("helpers.IsValidTradingPairs()", func() {
	Context("Is valid trading pairs input", func() {
		It("should return true", func() {
			Expect(utils.IsValidTradingPairs("BTC-USD,ETH-USD,LTC-USD")).To(BeTrue())
		})

		It("should return false", func() {
			Expect(utils.IsValidTradingPairs("BTC,USD,ltc-usb")).To(BeFalse())
		})
	})
})
