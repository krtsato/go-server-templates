package filter

import (
	"net/http"

	"github.com/krtsato/go-server-templates/2021-03-skeleton/internal/logger/app"
	"github.com/krtsato/go-server-templates/2021-03-skeleton/internal/webapi/apierr"
	"github.com/krtsato/go-server-templates/2021-03-skeleton/internal/webapi/apityp"

	"github.com/gin-gonic/gin"
)

// ErrorFilter リクエスト処理後のエラーハンドリング
type ErrorFilter struct{}

// NewErrorFilter middleware を生成
func NewErrorFilter() *ErrorFilter {
	return &ErrorFilter{}
}

// Execute GinFilter interface の実装
func (ErrorFilter) Execute(c *gin.Context) {
	c.Next()
	lastElm := c.Errors.Last()
	if lastElm == nil {
		return
	}

	switch err := lastElm.Err.(type) {
	case apierr.APIErr:
		handleAPIError(c, err)
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, apityp.ResultJSON{Error: err.Error()})
		app.Log.Errorf("Undefined Error: %s", err.Error())
	}
}

func handleAPIError(c *gin.Context, err apierr.APIErr) {
	errCode := err.ErrorCode()

	switch errCode {
	case apierr.PublicErrCode:
		c.AbortWithStatusJSON(err.StatusCode(), apityp.ResultJSON{Error: err.Error()})
	case apierr.PrivateErrCode:
		c.AbortWithStatus(err.StatusCode())
	default: // unknownErrCode
		c.AbortWithStatusJSON(err.StatusCode(), apityp.ResultJSON{Error: err.Error()})
	}
	app.Log.Errorf("Web API %s: %s", errCode.String(), err.Error())
}
