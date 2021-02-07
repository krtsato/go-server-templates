package auth

import (
	"encoding/json"

	"github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/internal/webapi/apierr"

	"github.com/gin-gonic/gin"
)

// Account アカウント
type Account struct {
	Mail  string `json:"mail"`
	Name  string `json:"name"`
	Roles Roles  `json:"roles"`
}

// UnmarshalAccount Request Header から Account を取得
func UnmarshalAccount(c *gin.Context) (*Account, error) {
	userJSON := c.Request.Header.Get(AuthAccount)
	if userJSON == "" {
		return &Account{}, apierr.NewForbiddenF(apierr.PublicErrCode, "request header doesn't have %s", AuthAccount)
	}
	var account Account
	if err := json.Unmarshal([]byte(userJSON), &account); err != nil {
		return &Account{}, apierr.NewForbiddenM(apierr.PublicErrCode, "failed to unmarshal Account")
	}
	return &account, nil
}
