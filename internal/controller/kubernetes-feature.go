package controller

import (
	"github.com/gin-gonic/gin"
	"github.io/MXuDong/example/config"
	"net/http"
)

func Config(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, config.Ctl.Config.KubernetesConfig)
}
