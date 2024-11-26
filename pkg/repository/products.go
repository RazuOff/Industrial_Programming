package repository

import "ginexample.com/pkg/models"

func GetProducts() []models.Product {
	var products []models.Product
	db.Find(&products)
	return products
}
