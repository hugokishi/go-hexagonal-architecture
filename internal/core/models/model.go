package models

import "github.com/hugokishi/hexagonal-go/internal/core/common"

type Product struct {
	common.Model
	Title       string             ` gorm:"column:title;NOT NULL" json:"title"`
	Description string             ` gorm:"column:description;NOT NULL" json:"description"`
	Type        common.ProductType ` gorm:"column:type;NOT NULL;type:product_type_enum" json:"type"`
	Price       float64            ` gorm:"column:price;NOT NULL" json:"price"`
}
