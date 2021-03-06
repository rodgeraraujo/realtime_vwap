package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/rodgeraraujo/vwap/pkg/coinbase"
	"github.com/rodgeraraujo/vwap/pkg/sigterm"
	"github.com/rodgeraraujo/vwap/pkg/types"
	"github.com/rodgeraraujo/vwap/pkg/utils"
)

func main() {
	// handling SIGTERM signal to close the program gracefully
	sigterm.HnderSigterm()

	tradingPairsInput := flag.String("pairs", "BTC-USD,ETH-USD,ETH-BTC", "trading pairs to monitor (comma separated)")
	windowSizeInput := flag.Int("window", 200, "window size")

	fmt.Println(*tradingPairsInput, *windowSizeInput)
	flag.Parse()

	if !utils.IsValidTradingPairs(*tradingPairsInput) {
		log.Fatal("invalid trading pair list")
	}
	pairNames := strings.Split(*tradingPairsInput, ",")

	// create an aggregator for each pair and a channel to send incoming sizedPrices to it
	pairs := make(map[string]chan *types.ChannelMessage)
	for _, name := range pairNames {
		aggregator := coinbase.NewPairAggregator(name, *windowSizeInput)
		incomingMatches := make(chan *types.ChannelMessage)
		go aggregator.ListenToMatches(incomingMatches)
		pairs[name] = incomingMatches
	}

	// subscribe to "matches" for the list of trading pairs
	incomingMessages := make(chan *types.ChannelMessage)
	coinbase.WebsocketSubscribe(incomingMessages, pairNames)

	for msg := range incomingMessages {
		if msg.Type != "match" {
			continue
		}

		aggregator, ok := pairs[msg.ProductId]
		if !ok {
			log.Printf("received a message for a pair we don't monitor: %s", msg.ProductId)
			continue
		}

		aggregator <- msg

	}
}
