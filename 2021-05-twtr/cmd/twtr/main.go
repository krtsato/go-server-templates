package main

func main() {
	// 設定値を取得
	port := "9999"

	// ロガー生成

	// サーバ初期化
	chiSrv := InjectDependencies()

	//サーバ起動
	if err := chiSrv.ListenAndServe(port); err != nil {
		panic(err)
	}
}
