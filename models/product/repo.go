package product

import (
	"database/sql"
	"encoding/json"
	"gadgetify/models/category"
)

type Product struct {
	ID          string            `json:"id"  gorm:"type:uuid;default:uuid_generate_v4();primaryKey;" `
	Name        string            `json:"name" form:"name" validate:"required,min=5,max=150"`
	SKU         string            `json:"sku" form:"sku" validate:"required,min=5,max=50"`
	ImageURL    *string           `json:"image_url" gorm:"type:varchar(255);"`
	Price       uint32            `json:"price" form:"price" validate:"required,min=0"`
	Description string            `json:"description" gorm:"type:text;" form:"description" validate:"max=5000"`
	Stock       uint32            `json:"stock" gorm:"default:0;" form:"stock" validate:"required,min=0"`
	CategoryID  sql.NullString    `json:"category_id"`
	Category    category.Category `json:"category"`
}

func (p Product) MarshalJSON() ([]byte, error) {
	type Alias Product
	if p.CategoryID.Valid {
		return json.Marshal(&struct {
			Alias
			CategoryID string `json:"category_id"`
		}{
			Alias:      (Alias)(p),
			CategoryID: p.CategoryID.String,
		})
	} else {
		return json.Marshal(&struct {
			Alias
			CategoryID interface{} `json:"category_id"`
			Category   interface{} `json:"category"`
		}{
			Alias:      (Alias)(p),
			CategoryID: nil,
			Category:   nil,
		})
	}
}
