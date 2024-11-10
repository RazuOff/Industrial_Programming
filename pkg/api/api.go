package api

import "github.com/gin-gonic/gin"

func StartServer(port string) {
	router := gin.Default()
	fillEndpoints(router)
	router.Run(port)
}

func fillEndpoints(router *gin.Engine) {

	router.GET("/products", GetProducts)

	router.GET("/products/:id", GetProductsById)

	router.POST("/products", createProduct)

	router.PUT("/products/:id", updateProduct)

	router.DELETE("/products/:id", deleteProduct)

	router.Run(":8080")
}
