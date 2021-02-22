package gglbdb

import (
	"context"
	"errors"

	"github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/pkg/infra"
)

// AccountClient Account Client interface
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

// DefaultAccountClientImpl 初期 DI 以外のタイミングで DB manager を渡して client を生成
func DefaultAccountClientImpl() AccountClient {
	return &accountClientImpl{dbm: infra.GetDBManager()}
}

// FindByAccountID AccountID (PK) からレコードを取得
func (a *accountClientImpl) FindByAccountID(ctx context.Context, accountID int) (*Account, error) {
	if a == nil {
		return &Account{}, errors.New("nil receiver of accountClientImpl is invalid")
	}
	var account Account
	err := a.dbm.Reader.First(&account, accountID).Error
	return &account, err
}

// FindByParam ページネーション対応の検索用パラメタからレコードを取得
func (a *accountClientImpl) FindByParam(ctx context.Context, param *FindAccountParam, pagination *Pagination) (Accounts, *Page, error) {
	if a == nil {
		return Accounts{}, &Page{}, errors.New("nil receiver of accountClientImpl is invalid")
	}

	reader := a.dbm.Reader
	// TODO: Null Only
	//if param.MediaNullOnly {
	//	reader = reader.Where("media is null")
	//}
	if len(param.Name) > 0 {
		reader = reader.Where("name = ?", param.Name)
	}
	// TODO: ThxCount Condition
	//if len(param.MediaCategory) > 0 {
	//	reader = reader.Where("media_category = ?", param.MediaCategory)
	//}

	var count int64
	if err := reader.Model(&Account{}).Count(&count).Error; err != nil {
		return nil, &Page{}, err
	}

	var accounts Accounts
	if err := reader.Limit(pagination.Limit).Offset(pagination.Offset).Find(&accounts).Error; err != nil {
		return nil, &Page{}, err
	}
	return accounts, NewPage(pagination, count), nil
}

// UpdateByParam null skip の更新用パラメタでレコードを置換
func (a *accountClientImpl) UpdateByParam(ctx context.Context, accountID int, param *UpdateAccountParam) (*Account, error) {
	if a == nil {
		return &Account{}, errors.New("nil receiver of accountClientImpl is invalid")
	}

	account, err := a.FindByAccountID(ctx, accountID)
	if err != nil {
		return &Account{}, err
	}
	if account.IsEmpty() {
		return &Account{}, nil
	}

	updAccount := account.NewWithUpdateParam(param)
	if err := a.dbm.Writer.Save(updAccount).Error; err != nil {
		return &Account{}, err
	}
	return updAccount, nil
}

// Save null skip しない Account モデルでレコードを置換
func (a *accountClientImpl) Save(ctx context.Context, account *Account) (*Account, error) {
	if a == nil {
		return &Account{}, errors.New("nil receiver of accountClientImpl is invalid")
	}
	err := a.dbm.Writer.Save(account).Error
	return account, err
}
