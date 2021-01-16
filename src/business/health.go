package business

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health returns the status of the application
func Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
	})
}
