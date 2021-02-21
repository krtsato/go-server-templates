package gglbdb

// Pagination pagination struct
type Pagination struct {
	Page   int
	Offset int
	Limit  int
}

// Pagination デフォルト値
const (
	defaultPage  = 1
	defaultLimit = 500
)

// DefaultPagination default pagination struct
func DefaultPagination() *Pagination {
	return &Pagination{Page: defaultPage, Offset: 0, Limit: defaultLimit}
}

// NewPagination pagination 生成
// page: 取得開始ページ, 1 始まり
// limit: 1ページあたりの取得サイズ, 0 件のとき defaultLimit
func NewPagination(page, limit int) *Pagination {
	if page < 1 {
		page = defaultPage
	}
	if limit == 0 {
		limit = defaultLimit
	}
	offset := (page - 1) * limit
	return &Pagination{Page: page, Offset: offset, Limit: limit}
}

// Next next pagination
func (p *Pagination) Next() {
	p.Page += 1
	p.Offset += p.Limit
}
