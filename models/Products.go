package models

import (
	"html"
	"strings"

	"gorm.io/gorm"
)

type Product struct {
	IdProduct   int64   `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	ProductName string  `json:"productName" gorm:"column:productName"`
	Brand       string  `json:"brand" gorm:"column:brand"`
	Description string  `json:"description" gorm:"column:description"`
	Photo1      string  `json:"photo1" gorm:"column:photo1"`
	Price       float32 `json:"price" gorm:"column:price"`
}

func (p *Product) GetProduct(db *gorm.DB) (*[]Product, error) {
	var err error
	products := []Product{}

	err = db.Find(&products).Error
	if err != nil {
		return &[]Product{}, err
	}
	return &products, nil
}

func (p *Product) GetProductId(db *gorm.DB, idProduct int64) (*Product, error) {
	var err error
	err = db.Find(&p, idProduct).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}

func (p *Product) SaveProduct(db *gorm.DB) (*Product, error) {
	var err error
	err = db.Debug().Model(&Product{}).Create(&p).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}

func (p *Product) UpdateProduct(db *gorm.DB) (*Product, error) {

	var err error
	err = db.Debug().Model(&Product{}).
		Where("id = ?", p.IdProduct).
		Updates(Product{
			ProductName: p.ProductName,
			Brand:       p.Brand,
			Description: p.Description,
			Photo1:      p.Photo1,
			Price:       p.Price,
		}).Error

	if err != nil {
		return &Product{}, err
	}
	return p, nil
}

func (p *Product) DeleteAPost(db *gorm.DB, productId int64) (int64, error) {

	db = db.Debug().Model(&Product{}).Where("id = ?", productId).Take(&Product{}).Delete(&Product{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (p *Product) Prepare() {
	p.ProductName = html.EscapeString(strings.TrimSpace(p.ProductName))
	p.Description = html.EscapeString(strings.TrimSpace(p.Description))
	p.Brand = html.EscapeString(strings.TrimSpace(p.Brand))
	p.Photo1 = html.EscapeString(strings.TrimSpace(p.Photo1))
}
