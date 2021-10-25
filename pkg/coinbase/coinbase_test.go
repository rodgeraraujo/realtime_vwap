package coinbase

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCoinbase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Calculator Suite")
}

var _ = Describe("coinbase", func() {

	Context("Vwap Calculator", func() {
		It("should create a new pair aggregator", func() {
			pairAggregator := NewPairAggregator("BTC-USD", 200)
			Expect(pairAggregator.PairName).To(Equal("BTC-USD"))
			Expect(pairAggregator.WindowSize).To(Equal(200))
			Expect(pairAggregator.TotalSize).To(Equal(0.0))
		})

		It("should add match to the aggregator", func() {
			sizedPrice := &SizedPrice{
				Price: 8.31,
				Size:  40.11,
			}

			pairAggregator := &PairAggregator{
				PairName:    "ETH-USD",
				SizedPrices: make([]*SizedPrice, 0),
				TotalSize:   180.88,
				TotalVWAP:   321.00,
				WindowSize:  200,
			}

			pairAggregator.addMatch(sizedPrice)

			Expect(pairAggregator.TotalSize).To(Equal(180.88 + 40.11))
			Expect(pairAggregator.TotalVWAP).To(Equal(321.00 + 8.31*40.11))

		})

		It("should return a new sized price", func() {
			match := &ChannelMessage{
				ProductId: "BTC-USD",
				Price:     "100.0",
				Size:      "1.0",
			}

			_, err := toSizedPrice(match)

			Expect(err).To(BeNil())
		})
	})
})
