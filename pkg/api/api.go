package api

import (
	"ginexample.com/pkg/auth"
	"ginexample.com/pkg/handler"
	"github.com/gin-gonic/gin"
)

func StartServer(port string) {
	router := gin.Default()
	fillEndpoints(router)
	router.Run(port)
}

func fillEndpoints(router *gin.Engine) {

	router.POST("/registrate", handler.Registrate)

	router.POST("/login", handler.Login)

	router.GET("/products", handler.GetProducts)

	router.GET("/products/:id", handler.GetProductsById)

	router.GET("/users", auth.AuthMiddleware(), handler.GetUsers)

	router.POST("/products", auth.AuthMiddleware(), auth.AdminCheck(), handler.CreateProduct)

	router.PUT("/products/:id", auth.AuthMiddleware(), auth.AdminCheck(), handler.UpdateProduct)

	router.DELETE("/products/:id", auth.AuthMiddleware(), auth.AdminCheck(), handler.DeleteProduct)

}
