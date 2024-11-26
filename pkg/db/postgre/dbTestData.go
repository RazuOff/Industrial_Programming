package postgre

import "ginexample.com/pkg/models"

var testProducts = []models.Product{
	{
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

var testUsers = []models.User{
	{
		Username: "Admin",
		Password: "1234",
		Role:     "Admin",
	},
	{
		Username: "user",
		Password: "1234",
	},
}
