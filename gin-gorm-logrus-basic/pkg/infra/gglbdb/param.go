package gglbdb

import "database/sql"

// FindAccountParam 検索用パラメタ
// オプショナルなパラメタのため使わない場合はゼロ値を設定
// WHERE 句で絞り込み
type FindAccountParam struct {
	Name         string
	ThxCount     string
	ThxCountCond int8
}

// UpdateAccountParam 更新用パラメタ
type UpdateAccountParam struct {
	Name     string
	Note     sql.NullString
	ThxCount int
}

// NewWithUpdateParam 更新用パラメタから null skip する Account を生成
func (a *Account) NewWithUpdateParam(param *UpdateAccountParam) *Account {
	if a == nil {
		return &Account{}
	}
	if param.Note.Valid {
		a.Note = param.Note
	}
	return a
}
