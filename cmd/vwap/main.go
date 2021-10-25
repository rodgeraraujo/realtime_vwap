package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/rodgeraraujo/vwap/pkg/coinbase"
	"github.com/rodgeraraujo/vwap/pkg/utils"
)

func main() {
	tradingPairsInput := flag.String("pairs", "BTC-USD,ETH-USD,ETH-BTC", "trading pairs to monitor")

	flag.Parse()

	if !utils.IsValidTradingPairs(*tradingPairsInput) {
		log.Fatal("invalid trading pair list")
	}
	pairNames := strings.Split(*tradingPairsInput, ",")

	// subscribe to "matches" for the list of trading pairs
	incomingMessages := make(chan *coinbase.ChannelMessage)
	coinbase.WebsocketSubscribe(incomingMessages, pairNames)

	for msg := range incomingMessages {
		if msg.Type != "match" {
			continue
		}

		fmt.Println(msg)

	}
}
