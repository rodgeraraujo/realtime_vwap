package coinbase

import (
	"encoding/json"
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

const (
	coinbaseUrl = "wss://ws-feed.exchange.coinbase.com"
	origin      = "http://localhost/"
)

// SubscribeMsg is the message received from the Coinbase Websocket
type SubscribeMsg struct {
	Type       string   `json:"type"`
	ProductIds []string `json:"product_ids"`
	Channels   []string `json:"channels"`
}

type ChannelMessage struct {
	Type      string `json:"type"`
	ProductId string `json:"product_id"`
	Sequence  int64  `json:"sequence"`
	Time      string `json:"time"`
	TradeId   int64  `json:"trade_id"`
	Price     string `json:"price"`
	Size      string `json:"size"`
}

// WebsocketSubscribe is the message sent to the Coinbase Websocket
func WebsocketSubscribe(ch chan *ChannelMessage, pairNames []string) {
	ws, err := websocket.Dial(coinbaseUrl, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	msg := SubscribeMsg{
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
			var msg *ChannelMessage
			err := websocket.JSON.Receive(ws, &msg)
			if err != nil {
				log.Fatal(fmt.Errorf("fail to read message: %v", err))
			}
			ch <- msg
		}
	}()
}
