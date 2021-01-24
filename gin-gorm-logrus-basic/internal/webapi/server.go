package webapi

import (
	"github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/internal/config"
	"github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/internal/logger/access"
	"github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/internal/logger/app"
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