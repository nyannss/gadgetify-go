package products

type Products struct {
	ID   string `json:"id"  gorm:"primaryKey"`
	Name string `json:"name"`
}
