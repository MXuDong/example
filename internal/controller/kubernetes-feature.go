package controller

import (
	"github.com/gin-gonic/gin"
	"github.io/MXuDong/example/config"
	"github.io/MXuDong/example/pkg/util/ctr"
)

func Config(ctx *gin.Context) {
	ctr.SuccessSingleObject(ctx, config.Ctl.Config.KubernetesConfig)
}
