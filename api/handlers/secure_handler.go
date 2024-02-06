package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SecuredHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "You have entered a secured route"})
}
