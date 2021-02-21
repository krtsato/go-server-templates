package gglbdb

import (
	"context"
	"github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/pkg/infra"
)

type AccountClient interface {
	FindByAccountID(ctx context.Context, accountID int) (*Account, error)
	FindByParam(ctx context.Context, param *FindAccountParam, pagination *Pagination) (Accounts, *Page, error)
	UpdateByParam(ctx context.Context, accountID int, param *UpdateAccountParam) (*Account, error)
	Save(ctx context.Context, account *Account) (*Account, error)
}

type accountClientImpl struct {
	dbm *infra.DBManager
}

// NewAccountClientImpl DB manager を渡して client を生成
func NewAccountClientImpl(dbm *infra.DBManager) AccountClient {
	return &accountClientImpl{dbm: dbm}
}

// 初期 DI 以外のタイミングで DB manager を渡して client を生成
func DefaultAccountClientImpl() AccountClient {
	return &accountClientImpl{dbm: infra.GetDBManager()}
}

func (a *accountClientImpl) FindByAccountID(ctx context.Context, accountID int) (*Account, error) {
	return &Account{}, nil
}

func (a *accountClientImpl) FindByParam(ctx context.Context, param *FindAccountParam, pagination *Pagination) (Accounts, *Page, error) {
	return Accounts{}, nil, nil
}
func (a *accountClientImpl) UpdateByParam(ctx context.Context, accountID int, param *UpdateAccountParam) (*Account, error) {
	return &Account{}, nil
}

func (a *accountClientImpl) Save(ctx context.Context, account *Account) (*Account, error) {
	return &Account{}, nil
}
