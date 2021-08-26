package models

type TokenDetails struct {
	AccessToken string  `json:"accessToken"`
	UserName    string  `json:"userName"`
	Company     string  `json:"company"`
	Sector      string  `json:"sector"`
	Wallet      float32 `json:"wallet"`
}

type Todo struct {
	UserID uint64 `json:"user_id"`
	Title  string `json:"title"`
}
