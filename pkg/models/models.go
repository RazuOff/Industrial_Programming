package models

type Product struct {
	ID                int     `gorm:"primaryKey"`
	ManufacturerId    int     `json:"manufacturer_id"`
	ProductCategoryId int     `json:"productCategory_id"`
	Name              string  `json:"name"`
	Price             float64 `json:"price"`
	Quantity          int     `json:"quantity"`
	ArticleNumber     int     `json:"articleNumber"`
	Description       string  `json:"description"`
	ImageUrl          string  `json:"image_url"`
}

type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `gorm:"default:'User'"`
}
