package category

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID    string `gorm:"type:uuid;default:uuid_generate_v4();primarykey;"`
	Name  string `json:"name"`
	Label string `json:"label"`
}
