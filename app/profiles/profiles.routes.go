package profiles

import (
	"resume-sys/core/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	ctrl := NewProfileController()

	authGroup := router.Group("/profiles").Use(middlewares.Auth)
	{
		authGroup.GET("", ctrl.Show)
		authGroup.GET("/:id", ctrl.Show)
		authGroup.POST("", ctrl.Create)

	}
}
