package orders

import (
	"KiteRunner/internal/model"
	"encoding/json"
	"os"
)

func LoadOrders() ([]model.Order, error) {
	b, err := os.ReadFile("config/mockresponses/orders.json")
	if err != nil {
		return nil, err
	}

	var resp model.OrdersResponse
	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}
