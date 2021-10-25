package coinbase

import (
	"log"
	"strconv"

	"github.com/rodgeraraujo/vwap/pkg/types"
)

// Match a minimal type for the match returned by the websocket
type Match struct {
	ProductId string `json:"product_id"`
	Price     string `json:"price"`
	Size      string `json:"size"`
}

// SizedPrice is a price and the associated size
type SizedPrice struct {
	Price float64 `json:"price"`
	Size  float64 `json:"size"`
}

// PairAggregator aggregates the VWAP for a pair
type PairAggregator struct {
	PairName             string
	SizedPrices          []*SizedPrice
	TotalSize, TotalVWAP float64
	WindowSize           int
}

func toSizedPrice(match *types.ChannelMessage) (*SizedPrice, error) {
	size, err := strconv.ParseFloat(match.Size, 64)
	if err != nil {
		return nil, err
	}

	price, err := strconv.ParseFloat(match.Price, 64)
	if err != nil {
		return nil, err
	}

	return &SizedPrice{
		Price: price,
		Size:  size,
	}, nil
}

// addMatch adds a sized price to the aggregator
func (pairAggregator *PairAggregator) addMatch(sizedPrice *SizedPrice) {
	pairAggregator.TotalSize += sizedPrice.Size
	pairAggregator.TotalVWAP += sizedPrice.Size * sizedPrice.Price
	pairAggregator.SizedPrices = append(pairAggregator.SizedPrices, &SizedPrice{
		Price: sizedPrice.Price,
		Size:  sizedPrice.Size,
	})
}

// calcVwap calculate the VWAP for a pair
func (pairAggregator *PairAggregator) calcVwap() float64 {
	if pairAggregator.TotalSize == 0 {
		return 0
	}
	return pairAggregator.TotalVWAP / pairAggregator.TotalSize
}

// removeOldestMatch removes the oldest match from the aggregator
func (pairAggregator *PairAggregator) removeOldestMatch() {
	oldMatch := pairAggregator.SizedPrices[0]

	pairAggregator.TotalSize -= oldMatch.Size
	pairAggregator.TotalVWAP -= oldMatch.Size * oldMatch.Price
}

// log logs the VWAP for a pair
func (pairAggregator *PairAggregator) log() {
	vwap := pairAggregator.calcVwap()
	log.Printf("%s: %f", pairAggregator.PairName, vwap)
}

// updateMatch updates the aggregator with a new match
func (pairAggregator *PairAggregator) updateMatch(match *types.ChannelMessage) {
	if len(pairAggregator.SizedPrices) == pairAggregator.WindowSize {
		pairAggregator.removeOldestMatch()
	}

	sizedPrice, err := toSizedPrice(match)

	if err != nil {
		return
	}

	pairAggregator.addMatch(sizedPrice)
}

// listenToMatches listens to matches and updates the aggregator
func (pairAggregator *PairAggregator) ListenToMatches(match chan *types.ChannelMessage) {
	for match := range match {
		pairAggregator.updateMatch(match)
		pairAggregator.log()
	}
}

// NewPairAggregator creates a new pair aggregator
func NewPairAggregator(pairName string, windowSize int) *PairAggregator {
	return &PairAggregator{
		PairName:    pairName,
		WindowSize:  windowSize,
		SizedPrices: make([]*SizedPrice, 0),
	}
}
