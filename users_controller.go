package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UsersIndex -
func UsersIndex(c *gin.Context) {
	c.String(http.StatusOK, "List users")
}

// UsersCreate -
func UsersCreate(c *gin.Context) {
	c.String(http.StatusOK, "Create user")
}

// UsersUpdate -
func UsersUpdate(c *gin.Context) {
	c.String(http.StatusOK, "Update users")
}
