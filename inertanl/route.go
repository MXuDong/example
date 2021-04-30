package inertanl

import (
	"github.com/gin-gonic/gin"
	"github.io/MXuDong/example/inertanl/server"
)

func Route(r *gin.Engine) {
	// ========================== base controller here
	BaseGroup := r.Group("/base")
	{
		BaseGroup.GET("/hello", server.HelloHandler)
	}
}
