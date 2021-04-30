package server

import (
	"github.com/gin-gonic/gin"
	"github.io/MXuDong/example/inertanl/model"
	"github.io/MXuDong/example/pkg/util/ctr"
)

// HelloHandler is the gin handler, for return the hello
func HelloHandler(ctx *gin.Context) {
	h := model.HelloStruct{}
	h.Value = "Hello Example"

	ctr.SuccessSingleObject(ctx, h)
}
