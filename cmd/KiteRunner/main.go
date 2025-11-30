package main

import (
	"KiteRunner/internal/app"
	"fmt"

	kiteconnect "github.com/zerodha/gokiteconnect/v4"
	kiteticker "github.com/zerodha/gokiteconnect/v4/ticker"
)

var (
	Ticker *kiteticker.Ticker
)

var (
	InstToken = []uint32{408065, 112129}
)

const (
	ApiKey    string = "my_api_key"
	ApiSecret string = "my_api_secret"
)

func main() {
	a := app.New()

	// Create a new Kite connect instance
	kc := kiteconnect.New(ApiKey)

	// Login URL from which request token can be obtained
	fmt.Println(kc.GetLoginURL())

	// Obtained request token after Kite Connect login flow
	requestToken := "request_token_obtained"

	// Get user details and access token
	data, err := kc.GenerateSession(requestToken, ApiSecret)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	// Set access token
	kc.SetAccessToken(data.AccessToken)

	// Get margins
	margins, err := kc.GetUserMargins()
	if err != nil {
		fmt.Printf("Error getting margins: %v", err)
	}
	fmt.Println("margins: ", margins)

	if err := a.Run(); err != nil {
		panic(err)
	}
}
