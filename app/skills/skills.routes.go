package skills

import (
	"resume-sys/core/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	ctrl := NewSkillsController()

	authGroup := router.Group("/skills").Use(middlewares.Auth)
	{
		authGroup.GET("", ctrl.List)
		authGroup.GET("/:id", ctrl.Show)
		authGroup.POST("", ctrl.Create)
		authGroup.PUT("/:id", ctrl.Update)
		authGroup.DELETE("/:id", ctrl.Delete)
	}
}
