package paw

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserResource -
type UserResource struct {
	App *App
}

// CreateUser -
func (r *UserResource) CreateUser(c *gin.Context) {
	var user User

	if !c.Bind(&user) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to decode request"})
		return
	}

	r.App.Database.Save(&user)
	c.JSON(http.StatusCreated, user)
}

// ListUsers -
func (r *UserResource) ListUsers(c *gin.Context) {
	var users []User

	r.App.Database.Order("id desc").Find(&users)

	c.JSON(http.StatusOK, users)
}

// ShowUser -
func (r *UserResource) ShowUser(c *gin.Context) {
	var userID = c.Params.ByName("id")
	var existing User

	if r.App.Database.First(&existing, userID).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
	} else {
		c.JSON(http.StatusOK, existing)
	}
}

// UpdateUser -
func (r *UserResource) UpdateUser(c *gin.Context) {
	var userID = c.Params.ByName("id")
	var user User

	if !c.Bind(&user) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to decode request"})
		return
	}

	var existing User
	if r.App.Database.First(&existing, userID).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
	} else {

		// Transfer attributes explicitly
		existing.FirstName = user.FirstName
		existing.MiddleName = user.MiddleName
		existing.LastName = user.LastName

		r.App.Database.Save(&existing)
		c.JSON(http.StatusOK, existing)
	}
}
