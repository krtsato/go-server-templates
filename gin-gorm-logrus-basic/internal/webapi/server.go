package webapi

import (
	"fmt"

	// "github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/internal/logger/access"
	"github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/internal/config"
	// "github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/internal/logger/app"
	// "github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/internal/logger"
	// "github.com/ad-agency/helix-api/internal/webapi/controller"
	"github.com/gin-gonic/gin"
)

// Server Web API
type Server interface {
	Start() error
	Shutdown()
}

// LoadConfig APP_ENV に応じた設定値群を選択する
func LoadConfig() (*config.AppConfig, error) {
	appEnv, applyErr := config.ApplyAppEnv(config.GetOSEnv())
	if applyErr != nil {
		return &config.AppConfig{}, applyErr
	}

	cfg, loadErr := config.LoadConfig(appEnv)
	if loadErr != nil {
		return &config.AppConfig{}, loadErr
	}
	/*
		if applyErr := app.ApplyLogger(cfg.Log); applyErr != nil {
			return config.AppConfig{}, applyErr
		}

		access.ApplyLogger(cfg.Log)
		return cfg, nil
	*/
}

/*
func newGinEngine(webCfg config.Web, handlers ...handler.GinRouterGroup) *gin.Engine {
	engine := gin.Default()
	filters := []filter.GinFilter{
		filter.NewBergRoleFilterRead(webCfg),
		filter.NewAccessLogFilter(),
		filter.NewErrorFilter(),
	}
	for _, f := range filters {
		engine.Use(f.Execute)
	}

	// 各 handler を GinRouterGroup としてエンドポイント化
	for _, h := range handlers {
		h.ApplyEndpoints(engine.Group(""))
	}
	return engine
}
*/
