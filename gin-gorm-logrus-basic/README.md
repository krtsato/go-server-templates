# gin-gorm-logrus-basic

## Requirement

### Versions

Go v1.15~  
aws-cli v2.0~  
macOS 64 bit  
Docker Engine v20.10~

### Runtime

- local 環境：docker-compose
- CI/CD 環境：Ubuntu
- dev/prd デプロイ：ローカルマシンの aws-cli

### Directories

参考: [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

```text
.
├── bin                          # binary outputs
├── build/package
│   ├── local                    # local Dockerfiles
│   └── server.dockerfile        # dev/prd Dockerfile
├── cmd/gin-gorm-logrus-basic    # package main
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

configs/.golangci-lint.yml にルールを記載する

`% make lint`

### fmt

go-imports は gofmt の 上位互換

`% make fmt-import`

`% make fmt`

### Health Check

```shell
% make local-up-all
% curl "http://localhost:9999/system/health-check"
```

## Deployment

coming later...

## Tests

coming later...
