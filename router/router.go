package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	initializeRoutes(gin.Default())
}
