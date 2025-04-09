// @title           Product Service API
// @version         1.0
// @description     API for managing products.
// @termsOfService  http://example.com/terms/

// @contact.name   API Support
// @contact.email  support@example.com

// @host      localhost:8081
// @BasePath  /api/v1
package main

import (
	"fmt"
	"log"
	"product-service/controllers"
	"product-service/database"
	"product-service/models"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "product-service/docs"
)

// @title Product Service API
// @version 1.0
// @description This is a sample API for managing products.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8081
// @BasePath /api/v1
func main() {
	database.ConnectDatabase()
	database.DB.AutoMigrate(&models.Product{})

	r := gin.Default()

	// Serve Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// CRUD routes
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products", controllers.GetAllProducts)
	r.GET("/products/:id", controllers.GetProductById)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)

	port := ":8081"
	fmt.Println("Product Service is running: http://localhost" + port)
	if err := r.Run(port); err != nil {
		log.Fatal("Server could not be started:", err)
	}
}
