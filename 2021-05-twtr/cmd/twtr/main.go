package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/krtsato/go-server-templates/2021-05-twtr/pkg/appconf"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	conf, err := appconf.LoadFacade()
	if err != nil {
		panic(err)
	}

	// TODO: generate applog

	server := InjectDependencies()

	if err := server.ListenAndServe(ctx, conf.AppConf.Rest.Port); err != nil {
		panic(err)
	}
}
