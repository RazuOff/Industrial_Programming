package repository

import (
	"ginexample.com/pkg/db/postgre"
	"ginexample.com/pkg/models"
)

func GetProducts() ([]models.Product, error) {
	var products []models.Product
	err := postgre.DB.Find(&products).Error
	return products, err
}

func AddProducts(products []models.Product) error {

	return postgre.DB.Create(products).Error
}

func GetProductById(id string) (models.Product, error) {
	var product models.Product
	err := postgre.DB.Where("id = ?", id).First(&product).Error
	return product, err
}

func DeleteProduct(prod models.Product) error {

	return postgre.DB.Delete(&prod).Error
}

func UpdateProduct(prod models.Product) error {

	return postgre.DB.Save(prod).Error
}
