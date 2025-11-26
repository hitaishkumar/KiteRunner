package model

type Instrument struct {
	InstrumentToken int     `csv:"instrument_token"`
	ExchangeToken   int     `csv:"exchange_token"`
	TradingSymbol   string  `csv:"tradingsymbol"`
	Name            string  `csv:"name"`
	LastPrice       float64 `csv:"last_price"`
	Expiry          string  `csv:"expiry"`
	Strike          float64 `csv:"strike"`
	TickSize        float64 `csv:"tick_size"`
	LotSize         int     `csv:"lot_size"`
	InstrumentType  string  `csv:"instrument_type"`
	Segment         string  `csv:"segment"`
	Exchange        string  `csv:"exchange"`
}
