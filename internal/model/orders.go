package model

type OrdersResponse struct {
	Status string  `json:"status"`
	Data   []Order `json:"data"`
}

type Order struct {
	PlacedBy           string                 `json:"placed_by"`
	OrderID            string                 `json:"order_id"`
	ExchangeOrderID    *string                `json:"exchange_order_id"`
	ParentOrderID      *string                `json:"parent_order_id"`
	Status             string                 `json:"status"`
	StatusMessage      *string                `json:"status_message"`
	StatusMessageRaw   *string                `json:"status_message_raw"`
	OrderTimestamp     string                 `json:"order_timestamp"`
	ExchangeUpdateTime *string                `json:"exchange_update_timestamp"`
	ExchangeTimestamp  *string                `json:"exchange_timestamp"`
	Variety            string                 `json:"variety"`
	Modified           bool                   `json:"modified"`
	Exchange           string                 `json:"exchange"`
	TradingSymbol      string                 `json:"tradingsymbol"`
	InstrumentToken    int64                  `json:"instrument_token"`
	OrderType          string                 `json:"order_type"`
	TransactionType    string                 `json:"transaction_type"`
	Validity           string                 `json:"validity"`
	ValidityTTL        int64                  `json:"validity_ttl"`
	Product            string                 `json:"product"`
	Quantity           int                    `json:"quantity"`
	DisclosedQuantity  int                    `json:"disclosed_quantity"`
	Price              float64                `json:"price"`
	TriggerPrice       float64                `json:"trigger_price"`
	AveragePrice       float64                `json:"average_price"`
	FilledQuantity     int                    `json:"filled_quantity"`
	PendingQuantity    int                    `json:"pending_quantity"`
	CancelledQuantity  int64                  `json:"cancelled_quantity"`
	AuctionNumber      *string                `json:"auction_number"`
	MarketProtection   int64                  `json:"market_protection"`
	Meta               map[string]interface{} `json:"meta"`
	Tag                *string                `json:"tag"`
	GUID               *string                `json:"guid"`
	Tags               []string               `json:"tags"`
}

type OrderHistoryResponse struct {
	Status string         `json:"status"`
	Data   []OrderHistory `json:"data"`
}

type OrderHistory struct {
	AveragePrice      float64 `json:"average_price"`
	CancelledQuantity int64   `json:"cancelled_quantity"`
	DisclosedQuantity int64   `json:"disclosed_quantity"`
	Exchange          string  `json:"exchange"`
	ExchangeOrderID   string  `json:"exchange_order_id"`
	ExchangeTimestamp string  `json:"exchange_timestamp"`
	FilledQuantity    int64   `json:"filled_quantity"`
	InstrumentToken   int64   `json:"instrument_token"`
	OrderID           string  `json:"order_id"`
	OrderTimestamp    string  `json:"order_timestamp"`
	OrderType         string  `json:"order_type"`
	ParentOrderID     string  `json:"parent_order_id"`
	PendingQuantity   int64   `json:"pending_quantity"`
	PlacedBy          string  `json:"placed_by"`
	Price             float64 `json:"price"`
	Product           string  `json:"product"`
	Quantity          int64   `json:"quantity"`
	Status            string  `json:"status"`
	StatusMessage     string  `json:"status_message"`
	Tag               string  `json:"tag"`
	TradingSymbol     string  `json:"tradingsymbol"`
	TransactionType   string  `json:"transaction_type"`
	TriggerPrice      float64 `json:"trigger_price"`
	Validity          string  `json:"validity"`
	Variety           string  `json:"variety"`
	Modified          bool    `json:"modified"`
}

type OrderTrade struct {
	TradeID           string  `json:"trade_id"`
	OrderID           string  `json:"order_id"`
	Exchange          string  `json:"exchange"`
	TradingSymbol     string  `json:"tradingsymbol"`
	InstrumentToken   int64   `json:"instrument_token"`
	Product           string  `json:"product"`
	AveragePrice      float64 `json:"average_price"`
	Quantity          int     `json:"quantity"`
	ExchangeOrderID   string  `json:"exchange_order_id"`
	TransactionType   string  `json:"transaction_type"`
	FillTimestamp     string  `json:"fill_timestamp"`
	OrderTimestamp    string  `json:"order_timestamp"`
	ExchangeTimestamp string  `json:"exchange_timestamp"`
}

type OrderTradeResponse struct {
	Status string       `json:"status"`
	Data   []OrderTrade `json:"data"`
}
