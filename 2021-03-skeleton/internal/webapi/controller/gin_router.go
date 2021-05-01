package controller

import (
	"github.com/gin-gonic/gin"
)

// GinRouterGroup この interface の実装により各エンドポイントを作成
type GinRouterGroup interface {
	ApplyEndpoints(r *gin.RouterGroup)
}
