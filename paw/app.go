package paw

import (
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"gopkg.in/yaml.v2"
)

// App -
type App struct {
	Engine   *gin.Engine
	Database gorm.DB

	Config struct {
		Service struct {
			Host string
		}

		Database struct {
			User             string
			Password         string
			Host             string
			Name             string
			ConnectionString string
		}
	}
}

// New -
func New() *App {
	app := &App{}

	loadConfig(app)

	return app
}

func loadConfig(app *App) {
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
	err = yaml.Unmarshal([]byte(ymlData), &app.Config)
	if err != nil {
		log.Error("Unable to unmarshal config", err)
		panic(err)
	}

	// ENV overrides

	servicePort := os.Getenv("PORT")
	if servicePort != "" {
		app.Config.Service.Host = ":" + servicePort
	}

	databaseURL := os.Getenv("DATABASE_URL")
	var connectionString string
	if databaseURL != "" {
		connectionString, _ := pq.ParseURL(databaseURL)
		connectionString += " sslmode=require"
	} else {
		connectionString = "postgres://" + app.Config.Database.User + ":" + app.Config.Database.Password +
			"@" + app.Config.Database.Host + ":5432/" + app.Config.Database.Name +
			"?sslmode=disable"
	}
	app.Config.Database.ConnectionString = connectionString
}
