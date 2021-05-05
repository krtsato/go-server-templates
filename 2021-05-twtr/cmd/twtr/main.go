package main

import (
	"context"
	"os"
	"os/signal"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	// 設定値を取得
	port := "9999"

	// ロガー生成

	// サーバ初期化
	chiSrv := InjectDependencies()

	//サーバ起動
	if err := chiSrv.ListenAndServe(ctx, port); err != nil {
		panic(err)
	}
}
