package infra

import (
	"errors"
	"reflect"
)

// QueryClient raw SQL client
type QueryClient interface {
	// Find
	// raw SQL を使用してデータを取得
	// WHERE 句用の引数には slice or value 型を指定
	Find(query string, result interface{}, params interface{}) error

	// FindByVariadicParams
	// raw SQL を使用してデータを取得
	// WHERE 句用に可変長引数を指定
	FindByVariadicParams(query string, result interface{}, params ...interface{}) error
}

type queryClientImpl struct {
	dbm *DBManager
}

// DefaultQueryClientImpl init DB managerを用いたclient生成
func DefaultQueryClientImpl() QueryClient {
	return &queryClientImpl{dbm: GetDBManager()}
}

// NewQueryClientImpl DB managerを渡してclient生成
func NewQueryClientImpl(dbm *DBManager) QueryClient {
	return &queryClientImpl{dbm: dbm}
}

// Find
// raw SQL を使用してデータを取得
// WHERE 句用の引数には slice or value 型を指定
//
// ex1) query := SELECT * FROM schema.dummy WHERE user_id = ? AND status = ?
//     var result Result
//     client.Find(query, &result, []string{"100", "valid"})
// ex2) query := SELECT * FROM schema.dummy WHERE user_id = ?
//     var result Result
//     client.Find(query, &result, "valid")
func (c *queryClientImpl) Find(query string, result interface{}, params interface{}) error {
	if validErr := validateFindArgs(result, params); validErr != nil {
		return validErr
	}
	vp := reflect.ValueOf(params)
	if vp.Kind() == reflect.Slice || vp.Kind() == reflect.Array {
		vps := make([]interface{}, vp.Len())
		for i := 0; i < vp.Len(); i++ {
			vps[i] = vp.Index(i).Interface()
		}
		return c.dbm.Reader.Raw(query, vps...).Scan(result).Error
	}
	return c.dbm.Reader.Raw(query, params).Scan(result).Error
}

// FindByVariadicParams
// raw SQL を使用してデータを取得
// WHERE 句用に可変長引数を指定
//
// ex) query := SELECT * FROM schema.dummy WHERE user_id = ? AND status = ?
//     var result Result
//     client.Find(query, &result, 100, "valid")
func (c *queryClientImpl) FindByVariadicParams(query string, result interface{}, params ...interface{}) error {
	if validErr := validateFindArgs(result, params); validErr != nil {
		return validErr
	}
	return c.dbm.Reader.Raw(query, params).Scan(result).Error
}

// validateFindArgs 必要最低限のバリデーション
func validateFindArgs(result interface{}, params interface{}) error {
	if reflect.TypeOf(result).Kind() != reflect.Ptr {
		return errors.New(" result should be pointer")
	}
	if reflect.TypeOf(params).Kind() == reflect.Ptr {
		return errors.New(" params should not be pointer")
	}
	return nil
}
