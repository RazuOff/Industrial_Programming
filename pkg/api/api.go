package api

import (
	_ "ginexample.com/docs"
	"ginexample.com/pkg/auth"
	"ginexample.com/pkg/handler"
	"ginexample.com/pkg/handler/products"
	taskhandler "ginexample.com/pkg/handler/taskHandler"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartServer(port string) {
	router := gin.Default()
	fillEndpoints(router)
	router.Run(port)
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func fillEndpoints(router *gin.Engine) {

	router.POST("/registrate", handler.Registrate)
	router.POST("/login", handler.Login)
	router.GET("/users", auth.AuthMiddleware(), handler.GetUsers)

	router.GET("/products", products.GetProducts)
	router.GET("/products/:id", products.GetProductsById)
	router.GET("/products/timeout", products.GetProductsWithTimeout)
	router.POST("/products", auth.AuthMiddleware(), auth.AdminCheck(), products.CreateProduct)
	router.PUT("/products", auth.AuthMiddleware(), auth.AdminCheck(), products.UpdateProduct)
	router.DELETE("/products/:id", auth.AuthMiddleware(), auth.AdminCheck(), products.DeleteProduct)

	router.POST("/tasks", taskhandler.CreateTask)
	router.GET("/tasks/:id", taskhandler.GetTask)
	router.DELETE("/tasks/:id", taskhandler.CancelTask)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
}
