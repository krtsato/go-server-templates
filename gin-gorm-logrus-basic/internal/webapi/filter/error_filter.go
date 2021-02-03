package filter

import (
	"net/http"

	"github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/internal/logger/app"
	"github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/internal/webapi/apierr"
	"github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/internal/webapi/apityp"

	"github.com/gin-gonic/gin"
)

// ErrorFilter リクエスト処理後のエラーハンドリング
type ErrorFilter struct{}

// NewErrorFilter middleware を生成
func NewErrorFilter() *ErrorFilter {
	return &ErrorFilter{}
}

// Execute Filter interface の実装
func (ErrorFilter) Execute(c *gin.Context) {
	c.Next()
	lastElm := c.Errors.Last()
	if lastElm == nil {
		return
	}

	switch err := lastElm.Err.(type) {
	case apierr.APIError: // Web API のエラーハンドリング
		handleAPIError(c, err)
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, apityp.ResultJSON{Error: err.Error()})
		app.Log.Errorf("Undefined Error: %v", err)
	}
}

func handleAPIError(c *gin.Context, err apierr.APIError) {
	errType := err.ErrorType()

	switch errType {
	case apierr.PublicErrType: // publicErrorType
		c.AbortWithStatusJSON(err.StatusCode(), api.ResultJSON{Error: err.Error()})
	case apierr.PrivateErrType: // privateErrorType
		c.AbortWithStatus(err.StatusCode())
	default: // unknownErrorType
		c.AbortWithStatus(err.StatusCode())
	}
	app.Log.Errorf("Web API %s: %v", errType.String(), err)
}
