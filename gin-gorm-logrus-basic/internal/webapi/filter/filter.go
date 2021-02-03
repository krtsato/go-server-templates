package filter

import "github.com/gin-gonic/gin"

// Filter Web API middleware
// フィルタ処理前の共通処理
// c.Next()
// フィルタ処理後の共通処理
type Filter interface {
	Execute(c *gin.Context)
}
