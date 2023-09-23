package routes

import (
	"github.com/gin-gonic/gin"
)

func Routes(group *gin.RouterGroup) {
	group = group.Group("/auth")

	group.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{})
	})
}
