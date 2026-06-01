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

	PostRepository := repository.NewPostRepository(dbConection)
	PostUseCase := usecase.NewPostUseCase(PostRepository)
	PostController := controller.NewPostController(PostUseCase)

	CommentRepository := repository.NewCommentRepository(dbConection)
	CommentUseCase := usecase.NewCommentUseCase(CommentRepository, PostRepository)
	CommentController := controller.NewCommentController(CommentUseCase)

	server := gin.Default()

	// CORS Middleware
	server.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Product Routes
	server.GET("/products", ProductController.GetProducts)
	server.GET("/products/:id", ProductController.GetProductById)
	server.DELETE("/products/:id", ProductController.DeleteProduct)
	server.PUT("/products/:id", ProductController.UpdateProduct)
	server.POST("/products", ProductController.CreateProduct)

	// Post Routes
	server.GET("/posts", PostController.GetPosts)
	server.POST("/posts", PostController.CreatePost)
	server.POST("/posts/:id/like", PostController.ToggleLike)
	server.POST("/posts/:id/repost", PostController.ToggleRepost)

	// Comment Routes
	server.GET("/posts/:id/comments", CommentController.GetComments)
	server.POST("/posts/:id/comments", CommentController.CreateComment)

	server.Run(":8080")
}
