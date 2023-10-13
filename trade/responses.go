package trade

// TradeRequestDoc is used to return the response which contains only one message field
type TradeRequestDoc struct {
	Key string        `json:"key"`
	Doc *TradeRequest `json:"doc"`
}
