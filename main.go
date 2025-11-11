package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/k-sub1995/gin-swagger-template/docs"
	"github.com/k-sub1995/gin-swagger-template/handlers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin Template API
// @version 1.0
// @description Sample server for Gin with Swagger.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()

	r.GET("/ping", handlers.Ping)

	// use ginSwagger middleware to serve the API docs
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	// Add ReDoc
	r.GET("/redoc", handlers.ReDoc)

	r.Run(":8080")
}
