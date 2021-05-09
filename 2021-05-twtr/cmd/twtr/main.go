package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/krtsato/go-server-templates/2021-05-twtr/pkg/appconf"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	conf, err := appconf.LoadFacade()
	if err != nil {
		panic(err)
	}

	// TODO: generate logger

	chiSrv := InjectDependencies()

	if err := chiSrv.ListenAndServe(ctx, conf.AppConf().WebAPIPort()); err != nil {
		panic(err)
	}
}
