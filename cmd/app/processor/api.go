package processor

import (
	"github.com/folklinoff/fitness-app/internal/handler"
	"github.com/gin-gonic/gin"
)

func api() *gin.Engine {
	e := gin.Default()
	e.POST("/login", handler.Login)
	e.POST("/register", handler.Register)
	return e
}
