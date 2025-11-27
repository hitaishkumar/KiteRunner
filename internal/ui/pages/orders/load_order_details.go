package orders

import (
	"KiteRunner/internal/model"
	"encoding/json"
	"os"
)

func LoadOrderDetails() ([]model.OrderHistory, error) {
	data, err := os.ReadFile("config/mockresponses/order_info.json")
	if err != nil {
		return nil, err
	}

	var resp model.OrderHistoryResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return resp.Data, nil
}
