package appconf

import (
	"fmt"

	"github.com/krtsato/go-server-templates/2021-05-twtr/configs"
	"github.com/krtsato/go-server-templates/2021-05-twtr/pkg/apperr"
)

// loadDBConf returns DBConf loaded from yml.
func loadDBConf(e AppEnv) (*configs.DBConf, error) {
	dbConfs, err := configs.UnmarshalDBConfs()
	if err != nil {
		return nil, apperr.ErrorF(apperr.Config, "failed to unmarshal app config")
	}

	fmt.Println(dbConfs)
	return nil, nil
}
