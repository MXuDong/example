package controller

import (
	"github.com/gin-gonic/gin"
	"github.io/MXuDong/example/config"
	"github.io/MXuDong/example/internal/model"
	"github.io/MXuDong/example/pkg/util/ctr"
)

// HelloHandler is the gin handler, for return the hello.
func HelloHandler(ctx *gin.Context) {
	h := model.HelloStruct{}
	h.Value = "Hello Example"

	ctr.SuccessSingleObject(ctx, h)
}

// ConfigHandler will return application run time config value.
// The config will auto convert to json.
func ConfigHandler(ctx *gin.Context) {
	ctr.SuccessSingleObject(ctx, config.Ctl.Config)
}
