package ctr

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SuccessEmpty will set response code with 204, and without any content
// Is the common func to return directly for gin handler
func SuccessEmpty(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}
