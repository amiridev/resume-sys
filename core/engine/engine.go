package engine

import (
	"resume-sys/app"
	"resume-sys/core/config"
	"resume-sys/core/middlewares"
	"resume-sys/core/swagger"
	"resume-sys/pkg/handlers"

	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func GetEngine() *gin.Engine {
	return engine
}

func Initialize() {
	gin.SetMode(config.GetEnv("GIN_MODE"))

	engine = gin.Default()
	engine.SetTrustedProxies(nil)
	engine.RedirectTrailingSlash = true
	engine.RedirectFixedPath = true

	// middlewares
	engine.Use(middlewares.Cors())
	engine.Use(gin.CustomRecovery(handlers.InternalErrorHandler))
}

func Serve() {
	engine.Run(config.GetRunAddress())
}

func RegisterRoutes() {
	routerGroup := GetEngine().Group("api")

	app.RegisterRoutes(routerGroup)

	swagger.RegisterSwagger(routerGroup)
}
