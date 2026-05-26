package router

import (
	"api-go/controller"
	"api-go/db"
	"api-go/repository"
	"api-go/usecase"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	dbConection, err := db.ConectDB()
	if err != nil {
		panic("Failed to connect to database")
	}

	ProductRepository := repository.NewProductRepository(dbConection)
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	ProductController := controller.NewProductController(ProductUseCase)

	server := gin.Default()
	server.GET("/products", ProductController.GetProducts)
	server.GET("/products/:id", ProductController.GetProductById)
	server.DELETE("/products/:id", ProductController.DeleteProduct)
	server.PUT("/products/:id", ProductController.UpdateProduct)
	server.POST("/products", ProductController.CreateProduct)

	server.Run(":8080")
}
