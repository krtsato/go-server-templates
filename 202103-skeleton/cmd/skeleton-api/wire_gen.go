// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/krtsato/go-server-templates/202103-skeleton/internal/config"
	"github.com/krtsato/go-server-templates/202103-skeleton/internal/webapi"
	"github.com/krtsato/go-server-templates/202103-skeleton/internal/webapi/controller"
)

// Injectors from wire.go:

func InitializeGinApp(webCfg config.Web) (*webapi.GinApp, error) {
	system := controller.InjectSystem()
	ginApp := webapi.NewGinApp(webCfg, system)
	return ginApp, nil
}