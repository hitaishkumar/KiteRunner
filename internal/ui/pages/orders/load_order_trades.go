package orders

import (
	"KiteRunner/internal/model"
	"encoding/json"
	"os"
)

func LoadOrderTrades() ([]model.OrderTrade, error) {
	b, err := os.ReadFile("config/mockresponses/order_trades.json")
	if err != nil {
		return nil, err
	}

	var resp model.OrderTradeResponse
	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}
