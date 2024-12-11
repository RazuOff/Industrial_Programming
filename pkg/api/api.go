package api

import (
	"ginexample.com/pkg/auth"
	"ginexample.com/pkg/handler"
	"ginexample.com/pkg/handler/products"
	taskhandler "ginexample.com/pkg/handler/taskHandler"
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
	router.GET("/users", auth.AuthMiddleware(), handler.GetUsers)

	router.GET("/products", products.GetProducts)
	router.GET("/products/:id", products.GetProductsById)
	router.POST("/products", auth.AuthMiddleware(), auth.AdminCheck(), products.CreateProduct)
	router.PUT("/products/:id", auth.AuthMiddleware(), auth.AdminCheck(), products.UpdateProduct)
	router.DELETE("/products/:id", auth.AuthMiddleware(), auth.AdminCheck(), products.DeleteProduct)

	router.POST("/tasks", taskhandler.CreateTask)
	router.GET("/tasks/:id", taskhandler.GetTask)
	router.DELETE("/tasks/:id", taskhandler.CancelTask)
}
