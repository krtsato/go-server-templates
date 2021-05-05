package main

import (
	"context"
	"os"
	"os/signal"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	// TODO: load configs
	port := "9999"

	// TODO: generate logger

	chiSrv := InjectDependencies()

	//サーバ起動
	if err := chiSrv.ListenAndServe(ctx, port); err != nil {
		panic(err)
	}
}
