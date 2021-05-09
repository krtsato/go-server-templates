package appconf

import (
	"fmt"
	"sync"

	"github.com/krtsato/go-server-templates/2021-05-twtr/configs"
)

// Facade is facade config.
type Facade struct {
	AppConf *configs.AppConf
	DBConf  *configs.DBConf
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
		facade.AppConf = app

		db, e := loadDBConf(appEnv)
		if e != nil {
			err = e
			return
		}
		facade.DBConf = db

		fmt.Printf("succeed in %v config setup\n", appEnv)
	})

	return facade, err
}
