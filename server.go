package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jbye/paw-api/paw"
	"github.com/stephenmuss/ginerus"
)

func main() {
	app := paw.New()

	app.Engine = gin.New()
	app.Engine.Use(ginerus.Ginerus())
	app.Database = paw.InitializeDatabase(app)

	userResource := &paw.UserResource{App: app}

	apiGroup := app.Engine.Group("/v1")
	{
		apiGroup.GET("/users", userResource.ListUsers)
		apiGroup.POST("/users", userResource.CreateUser)
		apiGroup.PUT("/users/:id", userResource.UpdateUser)
		apiGroup.GET("/users/:id", userResource.ShowUser)
	}

	app.Engine.Run(app.Config.Service.Host)
}
