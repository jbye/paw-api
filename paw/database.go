package paw

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // ??
)

// InitializeDatabase -
func InitializeDatabase(app *App) gorm.DB {
	db, err := gorm.Open("postgres", app.Config.Database.ConnectionString)

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
