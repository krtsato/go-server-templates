package skeletondb

import (
	"database/sql"

	"github.com/krtsato/go-server-templates/2021-03-skeleton/pkg/dater"
	"github.com/krtsato/go-server-templates/2021-03-skeleton/pkg/infra"
)

// Account accounts table struct
type Account struct {
	ID        int `gorm:"primary_key;auto_increment:true"`
	Name      string
	Note      sql.NullString
	ThxCount  int
	CreatedAt dater.LocalDatetime // gorm マネージドのため値セット不要
	UpdatedAt dater.LocalDatetime // gorm マネージドのため値セット不要
	DeletedAt dater.LocalDatetime // gorm マネージドのため値セット不要
}

// Accounts Account スライス
type Accounts []*Account

// TableAccount テーブル名
const TableAccount = "accounts"

// TableName DB名.テーブル名を返却
// gorm で使用するためダックタイピング
func (Account) TableName() string {
	return infra.SkeletonDB + "." + TableAccount
}

// IsEmpty empty check
func (a *Account) IsEmpty() bool {
	return *a == Account{}
}
