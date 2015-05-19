package main

import (
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stephenmuss/ginerus"
	"gopkg.in/yaml.v2"
)

// Config -
var config struct {
	Service struct {
		Host string
	}

	Database struct {
		User     string
		Password string
		Host     string
		Name     string
	}
}

var app struct {
	Database gorm.DB
	Engine   *gin.Engine
}

func loadConfig() {
	yamlPath := "config.yaml"
	if _, err := os.Stat(yamlPath); err != nil {
		log.Error("Config path is not valid")
		panic(err)
	}
	ymlData, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		log.Error("Unable to parse config", err)
		panic(err)
	}
	err = yaml.Unmarshal([]byte(ymlData), &config)
	if err != nil {
		log.Error("Unable to unmarshal config", err)
		panic(err)
	}
}

func main() {
	loadConfig()

	app.Engine = gin.New()
	app.Engine.Use(ginerus.Ginerus())
	app.Database = InitializeDatabase()

	userResource := &UserResource{db: app.Database}

	apiGroup := app.Engine.Group("/v1")
	{
		apiGroup.GET("/users", userResource.ListUsers)
		apiGroup.POST("/users", userResource.CreateUser)
		apiGroup.PUT("/users/:id", userResource.UpdateUser)
		apiGroup.GET("/users/:id", userResource.ShowUser)
	}

	app.Engine.Run(config.Service.Host)
}
