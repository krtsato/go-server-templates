package webapi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/krtsato/go-rest-templates/gin-gorm-logrus-basic/internal/webapi/controller"
)

// GinApp *gin.Engine による Web API
type GinApp struct {
	Engine *gin.Engine
}

// Start Server を起動
func (a *GinApp) Start() error {
	fmt.Println("Start GinApp")
	return a.Engine.Run(":9999") // TODO: cfg.Web.Port を参照させる
}

// Shutdown Server を終了
func (a *GinApp) Shutdown() {
	fmt.Println("Shutdown GinApp")
	/*
		if err := hlx.GetDbManager().Close(); err != nil {
			panic(err)
		}
	*/
}

// NewGinApp GinApp を生成
func NewGinApp( /*webCfg config.Web,*/ system *controller.System) *GinApp {
	engine := newGinEngine( /*webCfg, */ system)
	return &GinApp{Engine: engine}
}
