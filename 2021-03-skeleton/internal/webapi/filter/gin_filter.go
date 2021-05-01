package filter

import "github.com/gin-gonic/gin"

// GinFilter Web API middleware
// フィルタ処理前の共通処理
// c.Next()
// フィルタ処理後の共通処理
type GinFilter interface {
	Execute(c *gin.Context)
}
