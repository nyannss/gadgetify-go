package configs

import "gadgetify/models/products"

func initMigrate() {
	DB.AutoMigrate(&products.Products{})
}
