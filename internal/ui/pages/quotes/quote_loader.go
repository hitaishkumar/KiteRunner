package quotes

import (
	"encoding/json"
	"os"

	"KiteRunner/internal/model"
)

func LoadQuote(path string) (*model.QuoteResponse, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var q model.QuoteResponse
	if err := json.Unmarshal(data, &q); err != nil {
		return nil, err
	}

	return &q, nil
}
