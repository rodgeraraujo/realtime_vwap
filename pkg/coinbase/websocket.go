package coinbase

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/rodgeraraujo/vwap/pkg/types"
	"golang.org/x/net/websocket"
)

const (
	coinbaseUrl = "wss://ws-feed.exchange.coinbase.com"
	origin      = "http://localhost/"
)

// WebsocketSubscribe is the message sent to the Coinbase Websocket
func WebsocketSubscribe(ch chan *types.ChannelMessage, pairNames []string) {
	ws, err := websocket.Dial(coinbaseUrl, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	msg := types.SubscribeMsg{
		Type:       "subscribe",
		ProductIds: pairNames,
		Channels:   []string{"matches"},
	}

	// Send the message
	err = websocket.JSON.Send(ws, msg)
	if err != nil {
		log.Fatal(fmt.Errorf("failt to send message: %v", err))
	}

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(fmt.Errorf("failt to marshal message: %v", err))
	}

	if _, err := ws.Write(jsonMsg); err != nil {
		log.Fatal(fmt.Errorf("failt to write subscribe message: %v", err))
	}

	// Read messages from the websocket
	go func() {
		for {
			var msg *types.ChannelMessage
			err := websocket.JSON.Receive(ws, &msg)
			if err != nil {
				log.Fatal(fmt.Errorf("fail to read message: %v", err))
			}
			ch <- msg
		}
	}()
}
