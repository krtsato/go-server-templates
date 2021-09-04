package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/krtsato/go-server-templates/202103-skeleton/internal/webapi/apityp"
)

// System システム関連・開発者都合のエンドポイント
type System struct{}

// ApplyEndpoints *gin.RouterGroup にエンドポイントを追加
func (s *System) ApplyEndpoints(r *gin.RouterGroup) {
	rg := r.Group("/system")
	rg.GET("/health-check", s.healthCheck)
}

func (s *System) healthCheck(c *gin.Context) {
	c.JSON(200, apityp.ResultJSON{Result: "Healthy"})
}

// InjectSystem SystemController を生成
func InjectSystem() *System {
	return &System{}
}
