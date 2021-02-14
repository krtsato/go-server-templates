package gglbdb

import (
	"github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/pkg/dater"
	"github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/pkg/infra"
)

// Account accounts table struct
type Account struct {
	ID        int `gorm:"primary_key;auto_increment:false"`
	Name      string
	Note      string
	ThxCount  int
	CreatedAt dater.LocalDatetime // gorm マネージドのため値セット不要
	UpdatedAt dater.LocalDatetime // gorm マネージドのため値セット不要
	DeletedAt dater.LocalDatetime // gorm マネージドのため値セット不要
}

// Accounts Account Slice
type Accounts []Account

// TableAccount テーブル名
const TableAccount = "accounts"

// TableName DB名.テーブル名を返却
// gorm で使用するためダックタイピング
func (Account) TableName() string {
	return infra.GglbDB + "." + TableAccount
}

// IsEmpty empty check
func (a Account) IsEmpty() bool {
	return a == Account{}
}
