package models

type Product struct {
	Id                string
	ManufacturerId    int
	ProductCategoryId int
	Name              string
	Price             float64
	Quantity          int
	ArticleNumber     int
	Description       string
	ImageUrl          string
}

type User struct {
	Id       string
	Username string
	Password string
	Role     string
}
