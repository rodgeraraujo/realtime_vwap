package types

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
