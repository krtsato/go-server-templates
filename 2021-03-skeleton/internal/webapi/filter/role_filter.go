package filter

import (
	"github.com/krtsato/go-server-templates/2021-03-skeleton/internal/config"
	"github.com/krtsato/go-server-templates/2021-03-skeleton/internal/logger/access"
	"github.com/krtsato/go-server-templates/2021-03-skeleton/internal/webapi/apierr"
	"github.com/krtsato/go-server-templates/2021-03-skeleton/internal/webapi/apityp"
	"github.com/krtsato/go-server-templates/2021-03-skeleton/internal/webapi/auth"

	"github.com/gin-gonic/gin"
)

// RoleFilter HTTP リクエストが持つ Role
type RoleFilter struct {
	Roles  auth.Roles
	webCfg config.Web
}

// NewRoleFilter middleware の生成
func NewRoleFilter(roles auth.Roles, webCfg config.Web) *RoleFilter {
	return &RoleFilter{Roles: roles, webCfg: webCfg}
}

// NewRoleFilterReader ReaderRole 以上の権限を許容
func NewRoleFilterReader(webCfg config.Web) *RoleFilter {
	return NewRoleFilter(auth.Roles{auth.AdminRole, auth.WriterRole, auth.ReaderRole}, webCfg)
}

// NewRoleFilterWriter WriterRole 以上の権限を許容
func NewRoleFilterWriter(webCfg config.Web) *RoleFilter {
	return NewRoleFilter(auth.Roles{auth.AdminRole, auth.WriterRole}, webCfg)
}

// NewRoleFilterAdmin AdminRole のみ許容
func NewRoleFilterAdmin(webCfg config.Web) *RoleFilter {
	return NewRoleFilter(auth.Roles{auth.AdminRole}, webCfg)
}

// 権限エラーは 403 番で返却
func abortForbiddenRequest(c *gin.Context, err error) {
	access.Log.Error(err)
	c.AbortWithStatusJSON(403, apityp.ResultJSON{Error: err.Error()})
}

// Execute GinFilter の実装
func (f RoleFilter) Execute(c *gin.Context) {
	if !f.webCfg.AuthCheck {
		c.Next()
		return
	}

	account, err := auth.UnmarshalAccount(c)
	if err != nil {
		abortForbiddenRequest(c, err)
		return
	}
	if f.Roles.HasRole(account.Roles...) {
		c.Next()
		return
	}
	abortForbiddenRequest(c, apierr.NewForbiddenF(apierr.PublicErrCode, "forbidden roles: %v", account.Roles))
}
