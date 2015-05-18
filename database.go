package main

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// InitializeDatabase -
func InitializeDatabase() gorm.DB {
	db, err := gorm.Open("postgres", "user=paw dbname=paw sslmode=disable")

	if err != nil {
		log.Error("Cannot open database connection", err)
	}

	return db
}

// User -
type User struct {
	ID                int
	FirstName         string `sql:"size:80"`
	MiddleName        string `sql:"size:80"`
	LastName          string `sql:"size:80"`
	EncryptedPassword string `sql:"size:255"`
	FacebookUserID    string `sql:"size:255"`
	GoogleUserID      string `sql:"size:255"`
	UpdatedAt         time.Time
	CreatedAt         time.Time
}
