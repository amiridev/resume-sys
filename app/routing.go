package app

import (
	"resume-sys/app/auth"
	"resume-sys/app/courses"
	"resume-sys/app/documents"
	"resume-sys/app/experiences"
	"resume-sys/app/projects"
	"resume-sys/app/skills"
	"resume-sys/app/users"
	"resume-sys/core"

	"github.com/gin-gonic/gin"
)

// @tags    App
// @router	/api [get]
// @summary	app route, get heathy status
func HomeRoute(c *gin.Context) {
	ctrl := core.GetController()
	ctrl.Success(c, map[string]string{
		"heathy": "I'm OK...",
	})
}

func RegisterRoutes(router *gin.RouterGroup) {
	core.Initialize()

	router.GET("", HomeRoute)

	v1Group := router.Group("/v1")
	{
		auth.RegisterRoutes(v1Group)
		users.RegisterRoutes(v1Group)
		projects.RegisterRoutes(v1Group)
		courses.RegisterRoutes(v1Group)
		documents.RegisterRoutes(v1Group)
		skills.RegisterRoutes(v1Group)
		experiences.RegisterRoutes(v1Group)
	}
}
