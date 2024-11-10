package repository

import "ginexample.com/pkg/models"

var products = []models.Product{
	{
		Id:                "1",
		ManufacturerId:    1,
		ProductCategoryId: 1,
		Name:              "Intel Core i7-12700K",
		Price:             399.99,
		Quantity:          50,
		ArticleNumber:     1001,
		Description:       "Процессор Intel 12-го поколения с 12 ядрами и 20 потоками",
		ImageUrl:          "https://example.com/intel-i7.jpg",
	},
	{
		Id:                "2",
		ManufacturerId:    2,
		ProductCategoryId: 2,
		Name:              "NVIDIA GeForce RTX 4070",
		Price:             599.99,
		Quantity:          30,
		ArticleNumber:     1002,
		Description:       "Видеокарта с поддержкой Ray Tracing и DLSS 3.0",
		ImageUrl:          "https://example.com/rtx4070.jpg",
	},
}

func GetProducts() *[]models.Product {
	return &products
}
