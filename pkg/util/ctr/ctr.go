package ctr

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SuccessEmpty will set response code with 204, and without any content
// Is the common func to return directly for gin handler
func SuccessEmpty(c *gin.Context) {
	SuccessWithCode(c, http.StatusNoContent, nil)
}

// SuccessSingleObject will set response code with 200, and set content from target object to json.
// If input obj is empty, it will down level to use SuccessEmpty
func SuccessSingleObject(c *gin.Context, obj interface{}) {
	if obj == nil {
		SuccessEmpty(c)
		return
	}

	SuccessWithCode(c, http.StatusOK, obj)
}

func SuccessWithCode(c *gin.Context, code int, obj interface{}) {
	c.JSON(code, obj)
}

func Error(c *gin.Context, err error) {
	_ = c.Error(err)
	return
}
