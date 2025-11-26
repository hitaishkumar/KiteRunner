package model

type FullProfile struct {
	Status string `json:"status"`
	Data   struct {
		UserID       string `json:"user_id"`
		TwofaType    string `json:"twofa_type"`
		UserName     string `json:"user_name"`
		UserType     string `json:"user_type"`
		Email        string `json:"email"`
		Phone        string `json:"phone"`
		Broker       string `json:"broker"`
		BankAccounts []struct {
			Name    string `json:"name"`
			Branch  string `json:"branch"`
			Account string `json:"account"`
		} `json:"bank_accounts"`
		DpIds             []string `json:"dp_ids"`
		Products          []string `json:"products"`
		OrderTypes        []string `json:"order_types"`
		Exchanges         []string `json:"exchanges"`
		Pan               string   `json:"pan"`
		UserShortname     string   `json:"user_shortname"`
		AvatarURL         string   `json:"avatar_url"`
		Tags              []string `json:"tags"`
		PasswordTimestamp string   `json:"password_timestamp"`
		TwofaTimestamp    string   `json:"twofa_timestamp"`
		Meta              struct {
			Poa           string   `json:"poa"`
			Silo          string   `json:"silo"`
			AccountBlocks []string `json:"account_blocks"`
		} `json:"meta"`
	} `json:"data"`
}
