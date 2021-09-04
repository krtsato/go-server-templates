// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/krtsato/go-server-templates/202105-twtr/pkg/interface/rest"
	"github.com/krtsato/go-server-templates/202105-twtr/pkg/interface/rest/controller"
	"github.com/krtsato/go-server-templates/202105-twtr/pkg/interface/rest/router"
)

// Injectors from wire.go:

func InjectDependencies() *rest.Server {
	systemController := controller.InjectSystem()
	facadeRouter := router.InjectFacade(systemController)
	server := rest.InjectServer(facadeRouter)
	return server
}