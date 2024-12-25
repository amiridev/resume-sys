package bootstrap

import (
	"resume-sys/core/config"
	"resume-sys/core/engine"
	"resume-sys/database"
)

func Handle() {
	// init envs/configs
	config.InitEnvs()

	// db connections
	database.InitAndConnect()

	// setup engine & middlewares
	engine.Initialize()

	// register routes
	engine.RegisterRoutes()

	// init swagger

	engine.Serve()
}
