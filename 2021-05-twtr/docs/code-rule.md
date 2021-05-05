# Code Rule

## Naming

- 各レイヤのエンドポイント関数名は [Google Cloud API](https://cloud.google.com/apis/design/standard_methods?hl=ja) 標準メソッドに準拠する
    - ただし複数取得メソッドは `BatchGet` とする
    - e.g. `BatchGet(ctx, userIDs) (Users, error)`
    
## Documentation

- Global Export する変数/関数/型は定義の直上に `// {命名} {説明}` をコメントする
    - Lint ルールのため随時 `make fmt-imports lint` を実行すると良い

## DI

- DI が必要なパッケージルートに `injector.go` を配置し `wire.NewSet()` を定義する
    - cmd/app-name/wire.go で `wire.Build()` の引数として呼び出す

## Package Responsibility

### conv

- レイヤ間の型変換のためコンバータを使用する
    - 構造体フィールドを過不足なく渡すだけの場合, nil チェックは行わない
    - やむを得ず型変換にロジックが含まれる場合, nil チェックを行なう

- 型変換の方向
    - `toEntityFoo`: アプリ外部 → 内部
    - `toDTOBar`: アプリ内部 → 外部