# 2021-05-twtr

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
- go-sqlmock
- jwt-go
- k6
- ozzo-validation
- sync
- tbls
- wire
- zap

### Directories

[golang-standards/project-layout](https://github.com/golang-standards/project-layout) に準拠した  
ただし外部に閉じる必要がないため internal ディレクトリを除外した

```shell
.
├── build/package # dockerfiles
├── cmd/twtr      # package main
├── configs       # config files
├── deployments   # deployment files
├── docs          # documents
├── pkg
│   ├── appauth   # jwt and hash helper   
│   ├── appconf   # config scanner
│   ├── appctx    # context helper
│   ├── apperr    # errors inside app
│   ├── conv      # converter between entity and dto
│   ├── domain    # domain models
│   ├── interface # infra and webapi outside app
│   ├── logger    # access and app logger
│   └── usecase   # usecase interaction
└── scripts
    ├── make      # sub makefiles
    └── twtrdb    # db operation files

```

## Setup

### Quick Start

Makefile でコマンドを実行する

`% make help`

docker-compose で local 環境を構築する

`% make local-up-all`

docker-compose で local 環境を削除する

`% make local-down-all`

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
