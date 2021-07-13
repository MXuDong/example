package server

import (
	"github.com/gin-gonic/gin"
	"github.io/MXuDong/example/config"
	"github.io/MXuDong/example/internal/controller"
)

func Route(r *gin.Engine) {
	// ========================== base controller here
	BaseGroup := r.Group("/base")
	{
		BaseGroup.GET("/hello", controller.HelloHandler)   // hello handler
		BaseGroup.GET("/config", controller.ConfigHandler) // get the application runtime config
	}
	// ========================== invoke inner
	InnerGroup := r.Group("/inner")
	{
		InnerGroup.POST("/post", controller.TracePost)
	}

	// ========================== Protocol mock
	ProtocolMockGroup := r.Group("/protocol")
	{
		ProtocolMockGroup.POST("/tcp", controller.MockTcpRequest)
	}

	// ========================== Kubernetes feature
	KubernetesFeatureGroup := r.Group("/kubernetes")
	{
		KubernetesFeatureGroup.GET("/config", controller.Config)
		if config.Ctl.Config.KubernetesConfig.Enable() {
			// do some thing
		}
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
