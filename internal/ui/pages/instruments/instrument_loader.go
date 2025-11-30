package instruments

import (
	"encoding/csv"
	"os"
	"strconv"

	"KiteRunner/internal/model"
)

func LoadInstrumentsCSV(path string) ([]model.Instrument, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// skip header
	var list []model.Instrument
	for i := 1; i < len(rows); i++ {
		row := normalizeRow(rows[i], 12) // ensure row has 12 columns, fill missing

		inst := model.Instrument{
			InstrumentToken: atoi(row[0]),
			ExchangeToken:   atoi(row[1]),
			TradingSymbol:   fallback(row[2]),
			Name:            fallback(row[3]),
			LastPrice:       atof(row[4]),
			Expiry:          fallback(row[5]),
			Strike:          atof(row[6]),
			TickSize:        atof(row[7]),
			LotSize:         atoi(row[8]),
			InstrumentType:  fallback(row[9]),
			Segment:         fallback(row[10]),
			Exchange:        fallback(row[11]),
		}

		list = append(list, inst)
	}

	return list, nil
}
func LoadInstrumentsData(path string) ([]model.Instrument, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// skip header
	var list []model.Instrument
	for i := 1; i < len(rows); i++ {
		row := normalizeRow(rows[i], 12) // ensure row has 12 columns, fill missing

		inst := model.Instrument{
			InstrumentToken: atoi(row[0]),
			ExchangeToken:   atoi(row[1]),
			TradingSymbol:   fallback(row[2]),
			Name:            fallback(row[3]),
			LastPrice:       atof(row[4]),
			Expiry:          fallback(row[5]),
			Strike:          atof(row[6]),
			TickSize:        atof(row[7]),
			LotSize:         atoi(row[8]),
			InstrumentType:  fallback(row[9]),
			Segment:         fallback(row[10]),
			Exchange:        fallback(row[11]),
		}

		list = append(list, inst)
	}

	return list, nil
}

// If empty → replace with "NOT AVAILABLE"
func fallback(s string) string {
	if s == "" {
		return "NOT AVAILABLE"
	}
	return s
}

func normalizeRow(row []string, expected int) []string {
	// If row has fewer columns → add empty strings
	if len(row) < expected {
		diff := expected - len(row)
		for i := 0; i < diff; i++ {
			row = append(row, "")
		}
	}
	return row
}

func atoi(s string) int {
	if s == "" {
		return 0
	}
	n, _ := strconv.Atoi(s)
	return n
}

func atof(s string) float64 {
	if s == "" {
		return 0.0
	}
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
