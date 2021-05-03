# 2021-05-twtr

2020/05 に作成した SNS テンプレート

## Requirement

### Versions

- macOS 64 bit
- Docker Engine v20.10~
- Go v1.17~
- MySQL v8.0~

### Package Overview

- chi
- ozzo-validation
- firebase
- wire
- zap
- sync
- ent
- tbls
- gofakeit
- gomock
- go-sqlmock
- k6

### Directories

参考: [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

```text
.
├── build/package
│   ├── local                    # local Dockerfiles
├── cmd/twtr                     # package main
├── configs                      # config files
├── deployments                  # deployment files
├── docs                         # documents
├── internal                    
│   ├── infra/mysql              # infra layer
│   │   ├── initdb.d             # ddl and dml files 
│   │   ├── dump                 # dump outputs
│   │   └── log                  # log outputs
│   ├── service                  # service layer
│   └── webapi                   # webapi layer
└── scripts
    ├── make                     # sub Makefiles
    └── mysql                    # mysql operation shells
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
% curl "http://localhost:9999/system/health-check"
```
