package models

type Product struct {
	Id                string  `gorm:"primaryKey" json:"id"`
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
	Id       string `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `gorm:"default:'User'" json:"role"`
}
