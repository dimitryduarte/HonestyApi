package models

type Product struct {
	IdProduct   int64   `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	ProductName string  `json:"productName" gorm:"column:productName"`
	Brand       string  `json:"brand" gorm:"column:brand"`
	Photo1      string  `json:"photo1" gorm:"column:photo1"`
	Price       float32 `json:"price" gorm:"column:price"`
}
