package filter

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/internal/logger/access"

	"github.com/gin-gonic/gin"
)

// AccessLogFilter アクセスログの出力
type AccessLogFilter struct{}

// NewAccessLogFilter middleware を生成
func NewAccessLogFilter() *AccessLogFilter {
	return &AccessLogFilter{}
}

// Execute GinFilter interface の実装
func (AccessLogFilter) Execute(c *gin.Context) {
	start := time.Now()
	body := extractRequestBody(c.Request)
	c.Next()
	end := time.Now()

	// ログ出力
	access.Log.Info(&access.Attributes{
		Request:    c.Request,
		Start:      start,
		End:        end,
		Body:       body,
		StatusCode: c.Writer.Status(),
	})
}

// Request Body を文字列で返却
func extractRequestBody(req *http.Request) string {
	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}
	body := string(buf)

	// NopCloser により Closer インタフェースを付与することで
	// io.ReaderCloser が実装される
	// 後続の処理で Request Body を読み込めるようになる
	req.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
	return body
}
