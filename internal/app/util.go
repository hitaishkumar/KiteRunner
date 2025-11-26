package app

import (
	"KiteRunner/internal/model"
	"encoding/json"
	"os"
)

func loadFullProfile() (model.FullProfile, error) {
	var fp model.FullProfile

	data, err := os.ReadFile("config/mockresponses/full_profile.json")
	if err != nil {
		return fp, err
	}

	err = json.Unmarshal(data, &fp)
	return fp, err
}
