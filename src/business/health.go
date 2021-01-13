package business

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Health returns the status of the application
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
	})
}
