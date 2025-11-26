package model

type QuoteResponse struct {
	Status string                     `json:"status"`
	Data   map[string]InstrumentQuote `json:"data"`
}

type InstrumentQuote struct {
	InstrumentToken   int     `json:"instrument_token"`
	Timestamp         string  `json:"timestamp"`
	LastTradeTime     string  `json:"last_trade_time"`
	LastPrice         float64 `json:"last_price"`
	LastQuantity      int     `json:"last_quantity"`
	BuyQuantity       int     `json:"buy_quantity"`
	SellQuantity      int     `json:"sell_quantity"`
	Volume            int     `json:"volume"`
	AveragePrice      float64 `json:"average_price"`
	OI                int     `json:"oi"`
	OIDayHigh         int     `json:"oi_day_high"`
	OIDayLow          int     `json:"oi_day_low"`
	NetChange         float64 `json:"net_change"`
	LowerCircuitLimit float64 `json:"lower_circuit_limit"`
	UpperCircuitLimit float64 `json:"upper_circuit_limit"`

	OHLC  OHLCData    `json:"ohlc"`
	Depth MarketDepth `json:"depth"`
}

type OHLCData struct {
	Open  float64 `json:"open"`
	High  float64 `json:"high"`
	Low   float64 `json:"low"`
	Close float64 `json:"close"`
}

type MarketDepth struct {
	Buy  []DepthEntry `json:"buy"`
	Sell []DepthEntry `json:"sell"`
}

type DepthEntry struct {
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Orders   int     `json:"orders"`
}
