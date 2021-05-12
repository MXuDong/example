package server

import (
	"github.com/gin-gonic/gin"
	"github.io/MXuDong/example/config"
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

func Run() {
	// set gin mode
	ginMode := gin.DebugMode
	if config.Config.ServerConfig.Mod == config.ServerMod_Run {
		ginMode = gin.ReleaseMode
	}
	gin.SetMode(ginMode)
	g := gin.Default()
	// set route
	Route(g)
	_ = g.Run(config.Config.ServerConfig.Port)
}
