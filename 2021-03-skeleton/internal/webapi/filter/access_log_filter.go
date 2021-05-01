package filter

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/krtsato/go-server-templates/2021-03-skeleton/internal/logger/access"

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
// 引数が c.Request.Body (io.ReadCloser) でない理由
// ポインタ (*http.Request) を経由して Body を上書きするため
func extractRequestBody(req *http.Request) string {
	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}
	body := string(buf)

	// NopCloser により Close() メソッドを付与することで Closer インターフェースを満たす
	// bytes.NewBuffer(buf) の結果に io.ReaderCloser が実装されるため
	// 後続の処理で Request Body を読み込める
	req.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
	return body
}
