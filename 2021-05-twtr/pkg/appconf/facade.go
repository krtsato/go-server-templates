package appconf

import (
	"fmt"
	"sync"

	"github.com/krtsato/go-server-templates/2021-05-twtr/configs"
)

// Facade is facade config.
type Facade struct {
	appConf *configs.AppConf
	dbCong  *configs.DBConf
}

// once is singleton for Facade config.
var once sync.Once

// LoadFacade returns singleton Facade config.
func LoadFacade() (*Facade, error) {
	var (
		err    error
		facade = new(Facade)
	)

	once.Do(func() {
		appEnv, e := confirmAppEnv(loadOSAppEnv())
		if e != nil {
			err = e
			return
		}
		fmt.Printf("start to load %s config\n", appEnv.String())

		app, e := loadAppConf(appEnv)
		if e != nil {
			err = e
			return
		}
		facade.setAppConf(app)

		db, e := loadDBConf(appEnv)
		if e != nil {
			err = e
			return
		}
		facade.setDBConf(db)

		fmt.Printf("succeed in %v config setup\n", appEnv)
	})

	return facade, err
}

// AppConf returns the AppConf of Facade config.
func (f Facade) AppConf() *configs.AppConf {
	return f.appConf
}

// DBConf returns the DBConf of Facade config.
func (f Facade) DBConf() *configs.AppConf {
	return f.appConf
}

// setAppConf sets AppConf in Facade config.
func (f Facade) setAppConf(app *configs.AppConf) {
	f.appConf = app
}

// setDBConf sets DBConf in Facade config.
func (f Facade) setDBConf(db *configs.DBConf) {
	f.dbCong = db
}
