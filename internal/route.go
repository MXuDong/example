package internal

import (
	"github.com/gin-gonic/gin"
	"github.io/MXuDong/example/internal/server"
)

func Route(r *gin.Engine) {
	// ========================== base controller here
	BaseGroup := r.Group("/base")
	{
		BaseGroup.GET("/hello", server.HelloHandler)   // hello handler
		BaseGroup.GET("/config", server.ConfigHandler) // get the application runtime config
	}
}
