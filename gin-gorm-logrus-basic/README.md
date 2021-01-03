# gin-gorm-logrus-basic

## Requirement

### Versions

Go v1.15~  
aws-cli v2.0~  
macOS 64 bit  
Docker Engine v20.10~

### Directories

参考: [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

```gin-gorm-logrus-basic
.
├── bin                          # binary outputs
├── build/package
│   ├── local              # local Dockerfiles
│   └── server.dockerfile  # dev/prd Dockerfile
├── cmd/gin-gorm-logrus-basic    # package main
├── configs                      # config files
├── deployments                  # deployment files
├── docs                         # documents
├── internal                    
│   ├── infra/mysql
│   │   ├── initdb.d # ddl and dml files 
│   │   ├── dump     # dump outputs
│   │   └── log      # log outputs
│   ├── service
│   └── webapi
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

coming later...

## deployment

coming later...

## Tests

coming later...
