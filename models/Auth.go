package models

type Logins struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenDetails struct {
	AccessToken  string  `json:"AccessToken"`
	RefreshToken string  `json:"RefreshToken"`
	AccessUuid   string  `json:"AccessUuid"`
	RefreshUuid  string  `json:"RefreshUuid"`
	AtExpires    int64   `json:"AtExpires"`
	RtExpires    int64   `json:"RtExpires"`
	User         string  `json:"User"`
	Wallet       float32 `json:"Wallet"`
}

type Todo struct {
	UserID uint64 `json:"user_id"`
	Title  string `json:"title"`
}
