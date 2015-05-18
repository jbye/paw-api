package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stephenmuss/ginerus"
)

var app struct {
	Database gorm.DB
	Engine   *gin.Engine
}

func main() {
	app.Engine = gin.New()
	app.Engine.Use(ginerus.Ginerus())

	app.Database = InitializeDatabase()

	initRouting()

	app.Engine.Run(":3001")
}

func initRouting() {

	apiGroup := app.Engine.Group("/v1")
	{
		apiGroup.GET("/users", UsersIndex)
		apiGroup.POST("/users", UsersCreate)
		apiGroup.PUT("/users/:id", UsersUpdate)
	}
}
