package router

import (
	"api-go/controller"
	"api-go/db"
	"api-go/repository"
	"api-go/usecase"
	"os"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	dbConection, err := db.ConectDB()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
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

	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:4200"
	}

	// CORS Middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", frontendURL)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Session-ID")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Product Routes
	router.GET("/products", ProductController.GetProducts)
	router.GET("/products/:id", ProductController.GetProductById)
	router.DELETE("/products/:id", ProductController.DeleteProduct)
	router.PUT("/products/:id", ProductController.UpdateProduct)
	router.POST("/products", ProductController.CreateProduct)

	// Post Routes
	router.GET("/posts", PostController.GetPosts)
	router.GET("/posts/since", PostController.GetPostsSince)
	router.POST("/posts", PostController.CreatePost)
	router.POST("/posts/:id/like", PostController.ToggleLike)
	router.POST("/posts/:id/repost", PostController.ToggleRepost)

	// Comment Routes
	router.GET("/posts/:id/comments", CommentController.GetComments)
	router.POST("/posts/:id/comments", CommentController.CreateComment)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
