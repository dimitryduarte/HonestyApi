package models

type Users struct {
	IdUser   uint64  `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Email    string  `json:"email" gorm:"column:email"`
	Password string  `json:"password" gorm:"column:password"`
	Name     string  `json:"name" gorm:"column:name"`
	Company  string  `json:"company" gorm:"column:company"`
	Sector   string  `json:"sector" gorm:"column:sector"`
	Wallet   float64 `json:"wallet" gorm:"column:wallet"`
}
