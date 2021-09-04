package infra

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"sync"
	"time"
)

// Datasource 接続設定
type Datasource struct {
	Driver             string `yaml:"driver"`
	UseConnPool        bool   `yaml:"useConnPool"`
	MaxIdleConnSize    int    `yaml:"maxIdleConnSize"`
	MaxOpenConnSize    int    `yaml:"maxOpenConnSize"`
	ConnMaxLifetimeMin int    `yaml:"connMaxLifetimeMin"`
	Reader             Reader `yaml:"reader"`
	Writer             Writer `yaml:"writer"`
}

// Reader 読み出し設定
type Reader struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	URLQuery string `yaml:"urlQuery"`
	User     string `yaml:"user"`
	Pass     string `yaml:"pass"`
}

// Writer 書き込み設定
type Writer struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	URLQuery string `yaml:"urlQuery"`
	User     string `yaml:"user"`
	Pass     string `yaml:"pass"`
}

// DBManager 接続管理
type DBManager struct {
	Reader *gorm.DB
	Writer *gorm.DB
}

// singleton DB インスタンス
var once sync.Once
var dbm *DBManager

// GetDBManager アプリケーション終了時に必ず Close()
func GetDBManager() *DBManager {
	return dbm
}

// InitDB デフォルト Datasource で singleton DB インスタンスを生成
func InitDB(e DBEnv) (*DBManager, error) {
	for _, d := range defaultDataSources {
		if d.dbEnv == e {
			return InitDBByDataSource(d.datasource)
		}
	}
	return nil, fmt.Errorf("unknown DBEnv: %s", e.String())
}

// InitDBByDataSource カスタム Datasource で singleton DB インスタンスを生成
func InitDBByDataSource(ds Datasource) (*DBManager, error) {
	var err error
	once.Do(func() {
		conn, oErr := OpenDB(ds)
		if oErr != nil {
			err = oErr
		}
		dbm = conn
	})
	return dbm, err
}

// OpenDB 各種設定値を DBManager にセット
func OpenDB(d Datasource) (*DBManager, error) {
	// Reader
	r := d.Reader
	rDSN := r.User + ":" + r.Pass + "@tcp(" + r.Host + ":" + r.Port + ")/" + SkeletonDB + "?" + r.URLQuery
	rConn, rOpenErr := gorm.Open(mysql.Open(rDSN), &gorm.Config{})
	if rOpenErr != nil {
		return nil, rOpenErr
	}
	rDB, rDBErr := rConn.DB()
	if rDBErr != nil {
		defer closeSQL(rDB)
		return nil, rDBErr
	}
	if d.UseConnPool {
		rDB.SetMaxIdleConns(d.MaxIdleConnSize)
		rDB.SetMaxOpenConns(d.MaxOpenConnSize)
		rDB.SetConnMaxLifetime(time.Duration(d.ConnMaxLifetimeMin) * time.Minute)
	}

	// Writer
	w := d.Writer
	wDSN := w.User + ":" + w.Pass + "@tcp(" + w.Host + ":" + w.Port + ")/" + SkeletonDB + "?" + w.URLQuery
	wConn, wOpenErr := gorm.Open(mysql.Open(wDSN), &gorm.Config{})
	if wOpenErr != nil {
		defer closeSQL(rDB)
		return nil, wOpenErr
	}
	wDB, wDBErr := wConn.DB()
	if wDBErr != nil {
		defer closeSQL(wDB)
		return nil, wDBErr
	}
	if d.UseConnPool {
		wDB.SetMaxIdleConns(d.MaxIdleConnSize)
		wDB.SetMaxOpenConns(d.MaxOpenConnSize)
		wDB.SetConnMaxLifetime(time.Duration(d.ConnMaxLifetimeMin) * time.Minute)
	}
	return &DBManager{Reader: rConn, Writer: wConn}, nil
}

// closeSQL *sql.DB コネクションを終了
// infra パッケージ内部からの呼び出しを想定
func closeSQL(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Fatalf("closing DB error: %v", err)
	}
}

// CloseDB コネクションの終了
// infra パッケージ外部からの呼び出しを想定
func (dbm *DBManager) CloseDB() []error {
	if dbm == nil {
		return nil
	}

	// Reader
	var closeErrors []error
	rDB, rErr := dbm.Reader.DB()
	if rErr != nil {
		closeErrors = append(closeErrors, rErr)
	}
	if err := rDB.Close(); err != nil {
		closeErrors = append(closeErrors, err)
	}

	// Writer
	wDB, wErr := dbm.Writer.DB()
	if wErr != nil {
		closeErrors = append(closeErrors, wErr)
	}
	if err := wDB.Close(); err != nil {
		closeErrors = append(closeErrors, err)
	}
	return closeErrors
}

// DefaultDataSource デフォルト Datasource
type DefaultDataSource struct {
	dbEnv      DBEnv
	datasource Datasource
}

// 例: Aurora MySQL を想定
// DB にとってのデフォルト値をメモリ内部に確保する
// yml 設定ファイルの default フィールドはアプリにとってのデフォルト値
const (
	auroraDriver             = "mysql"
	auroraMaxIdleConnSize    = 10
	auroraMaxOpenConnSize    = 500
	auroraConnMaxLifetimeMin = 10
	auroraURLQuery           = "charset=utf8mb4&parseTime=True&loc=Asia%2FTokyo"
)

var defaultDataSources = []DefaultDataSource{
	{
		dbEnv: Local,
		datasource: Datasource{
			Driver:             auroraDriver,
			UseConnPool:        true,
			MaxIdleConnSize:    auroraMaxIdleConnSize,
			MaxOpenConnSize:    auroraMaxOpenConnSize,
			ConnMaxLifetimeMin: auroraConnMaxLifetimeMin,
			Reader: Reader{
				Host:     "mysql",
				Port:     "3306",
				URLQuery: auroraURLQuery,
				User:     "root",
				Pass:     "",
			},
			Writer: Writer{
				Host:     "mysql",
				Port:     "3306",
				URLQuery: auroraURLQuery,
				User:     "root",
				Pass:     "",
			},
		},
	},
	{
		dbEnv: Dev,
		datasource: Datasource{
			Driver:             auroraDriver,
			UseConnPool:        true,
			MaxIdleConnSize:    auroraMaxIdleConnSize,
			MaxOpenConnSize:    auroraMaxOpenConnSize,
			ConnMaxLifetimeMin: auroraConnMaxLifetimeMin,
			Reader: Reader{
				Host:     "", // Aurora Endpoint
				Port:     "3306",
				URLQuery: auroraURLQuery,
				User:     "user_r",
				Pass:     "pass_r",
			},
			Writer: Writer{
				Host:     "", // Aurora Endpoint
				Port:     "3306",
				URLQuery: auroraURLQuery,
				User:     "user_w",
				Pass:     "pass_w",
			},
		},
	},
	{
		dbEnv: Prd,
		datasource: Datasource{
			Driver:             auroraDriver,
			UseConnPool:        true,
			MaxIdleConnSize:    auroraMaxIdleConnSize,
			MaxOpenConnSize:    auroraMaxOpenConnSize,
			ConnMaxLifetimeMin: auroraConnMaxLifetimeMin,
			Reader: Reader{
				Host:     "", // Aurora Endpoint
				Port:     "3306",
				URLQuery: auroraURLQuery,
				User:     "user_r",
				Pass:     "pass_r",
			},
			Writer: Writer{
				Host:     "", // Aurora Endpoint
				Port:     "3306",
				URLQuery: auroraURLQuery,
				User:     "user_w",
				Pass:     "pass_w",
			},
		},
	},
}
