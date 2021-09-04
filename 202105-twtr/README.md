# 202105-twtr

2020/05 に作成した SNS テンプレート

## Requirement

### Versions

- macOS 64 bit
- Docker Engine v20.10~
- Go v1.16~
- MySQL v8.0~

### Package Overview

- air
- chi
- ent
- gofakeit
- gomock
- gomaxprcs
- go-sqlmock
- jwt-go
- k6
- ozzo-validation
- pprof
- automaxprocs
- sync
- tbls
- wire
- zap

### Directories

[golang-standards/project-layout](https://github.com/golang-standards/project-layout) に準拠した  
ただし外部に閉じる必要がないため internal ディレクトリを除外した

```shell
.
├── build         # dockerfiles
├── cmd/twtr      # package main
├── configs       # config files
├── deployments   # deployment files
├── docs          # documents
├── pkg
│   ├── appauth   # hash, jwt and session helper
│   ├── appconf   # config scanner
│   ├── appctx    # context helper
│   ├── applog    # access and app logger
│   ├── conv      # converter between entity and dto
│   ├── domain    # domain models
│   ├── interface # db, rest and grpc outside app
│   ├── mock      # generated mock 
│   └── usecase   # usecase interaction
└── scripts  # sub makefiles and db operation files
```

## Setup

### Rules

[docs/git-rule.md](https://github.com/krtsato/go-server-templates/tree/main/202105-twtr/docs/git-rule.md)  
[docs/code-rule.md](https://github.com/krtsato/go-server-templates/tree/main/202105-twtr/docs/code-rule.md)

### Quick Start

Makefile でコマンドを実行する

`% make help`

docker compose で local 環境を構築する

`% make local-up-all`

docker compose で local 環境を削除する

`% make local-down-all`

wire でインジェクタを生成する
gomock でモックを生成する

`% make generate`

### Local Check

### lint

configs/.golangci.yml にルールを記載する

`% make lint`

### fmt

`% make fmt`

### Health Check

```shell
% make local-up-all
% curl "http://localhost:9999/system/health"
```
