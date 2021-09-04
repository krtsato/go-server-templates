package webapi

import (
	"github.com/gin-gonic/gin"
	"github.com/krtsato/go-server-templates/202103-skeleton/internal/config"
	"github.com/krtsato/go-server-templates/202103-skeleton/internal/logger/access"
	"github.com/krtsato/go-server-templates/202103-skeleton/internal/logger/app"
	"github.com/krtsato/go-server-templates/202103-skeleton/internal/webapi/controller"
	"github.com/krtsato/go-server-templates/202103-skeleton/internal/webapi/filter"
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

	access.ApplyLogger(cfg.Log)
	if err := app.ApplyLogger(cfg.Log); err != nil {
		return &config.AppConfig{}, err
	}

	return cfg, nil
}

func newGinEngine(webCfg config.Web, controllers ...controller.GinRouterGroup) *gin.Engine {
	// middleware を適用
	engine := gin.Default()
	filters := []filter.GinFilter{
		filter.NewRoleFilterReader(webCfg),
		filter.NewAccessLogFilter(),
		filter.NewErrorFilter(),
	}
	for _, f := range filters {
		engine.Use(f.Execute)
	}

	// 各 controller を GinRouterGroup としてエンドポイント化
	for _, c := range controllers {
		c.ApplyEndpoints(engine.Group(""))
	}
	return engine
}
