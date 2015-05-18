package main

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// InitializeDatabase -
func InitializeDatabase() gorm.DB {
	connectionString := "postgres://" + config.Database.User + ":" + config.Database.Password +
		"@" + config.Database.Host + ":5432/" + config.Database.Name +
		"?sslmode=disable"

	log.Debug("ConnectionString: %s", connectionString)

	db, err := gorm.Open("postgres", connectionString)

	if err != nil {
		log.Error("Cannot open database connection", err)
	}

	return db
}

// User -
type User struct {
	ID                int        `gorm:"primary_key" json:"id"`
	FirstName         string     `sql:"size:80" json:"firstName"`
	MiddleName        string     `sql:"size:80" json:"middleName"`
	LastName          string     `sql:"size:80" json:"lastName"`
	EncryptedPassword string     `sql:"size:255" json:"-"`
	FacebookUserID    string     `sql:"size:255" json:"-"`
	GoogleUserID      string     `sql:"size:255" json:"-"`
	UpdatedAt         time.Time  `json:"updatedAt"`
	CreatedAt         time.Time  `json:"createdAt"`
	DeletedAt         *time.Time `json:"deletedAt"`
}
