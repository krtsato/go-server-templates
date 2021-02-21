package gglbdb

import "github.com/shopspring/decimal"

// Page 取得したレコードの Page
type Page struct {
	Page         int
	TotalPages   int
	TotalRecords int64
}

// NewPage Pagination, TotalRecord から Page を生成
func NewPage(pagination Pagination, totalRecords int64) *Page {
	var totalPage int64 = 1
	if pagination.Limit > 0 {
		totalPage = decimal.NewFromInt(totalRecords).Div(decimal.NewFromInt(int64(pagination.Limit))).Ceil().IntPart()
	}
	return &Page{
		Page:         pagination.Page,
		TotalPages:   int(totalPage),
		TotalRecords: totalRecords,
	}
}
