package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// UserResource -
type UserResource struct {
	db gorm.DB
}

// CreateUser -
func (r *UserResource) CreateUser(c *gin.Context) {
	var user User

	if !c.Bind(&user) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to decode request"})
		return
	}

	r.db.Save(&user)
	c.JSON(http.StatusCreated, user)
}

// ListUsers -
func (r *UserResource) ListUsers(c *gin.Context) {
	c.String(http.StatusOK, "ListUsers")
}

// UpdateUser -
func (r *UserResource) UpdateUser(c *gin.Context) {
	c.String(http.StatusOK, "UpdateUser")
}
