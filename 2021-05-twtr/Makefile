# ------------------------------------------------------------ #
# Env
# ------------------------------------------------------------ #

# App
APP_ENV=local # local | dev | prd
APP_NAME=twtr

# DB Command Args
TABLE_NAME=   # table name arg
MIGRATE_NAME= # migration name arg

# ------------------------------------------------------------ #
# Setup
# ------------------------------------------------------------ #

.PHONY: help
help: ## make コマンド一覧
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' ${MAKEFILE_LIST}

.PHONY: install-devtool
install-tools: ## local: tool コンテナで使うソフトウェアをインストール
	docker build -f build/package/local/tools.dockerfile -t twtr-tools --rm=true .

.PHONY: lint
lint: ## local: golangci-lint による静的解析
	docker run --rm -e APP_ENV=${APP_ENV} -v ${CURDIR}:/var/app/tools twtr-tools golangci-lint run -c ./configs/.golangci.yml

.PHONY: fmt
fmt: ## local: goimports のフォーマット
	docker run --rm -e APP_ENV=${APP_ENV} -v ${CURDIR}:/var/app/tools twtr-tools goimports -w ./

.PHONY: go-test
go-test: ## local: テストを実行
	go test -race -v ./cmd/... ./pkg/...

.PHONY: go-build
go-build: ## local: ソースコードをビルド
	go build -o ./${APP_NAME} ./cmd/${APP_NAME}/{main.go,wire_gen.go}

# ------------------------------------------------------------ #
# docker-compose
# ------------------------------------------------------------ #

.PHONY: local-up-all
local-up-all: ## local: docker-compose でイメージ作成・コンテナ起動
	docker-compose up --build

.PHONY: local-down-all
local-down-all: ## local: docker-compose ですべてのリソース削除
	docker-compose down -v --rmi all

.PHONY: local-logs
local-logs: ## local: docker-compose でコンテナログ出力
	docker-compose logs -f

.PHONY: local-tail-app
local-tail: ## local: tail でアプリケーションログ出力
	docker-compose exec ${APP_NAME} sh -c 'tail -f /var/log/${APP_NAME}/app.log'

.PHONY: local-exec-sh
local-exec: ## local: docker-compose で sh ログイン
	docker-compose exec ${APP_NAME} sh

.PHONY: local-cont-ls
local-cont-ls: ## local: docker container 一覧
	docker container ls -a

.PHONY: local-img-ls
local-img-ls: ## local: docker image 一覧
	docker image ls -a

.PHONY: local-cont-rm
local-cont-rm: ## local: exited した docker container 削除
	docker container ls -aqf status=exited | xargs docker rm -f

.PHONY: local-img-rm
local-img-rm: ## local: タグ無しの docker image 削除
	docker image ls -aqf dangling=true | xargs docker rmi -f

# ------------------------------------------------------------ #
# MySQL Client
# ------------------------------------------------------------ #

.PHONY: local-show-dbs
local-show-dbs: ## local: show databases
	mysql -u root -ppass_root -h localhost -P 9998 --protocol=tcp -e "show databases;"

.PHONY: local-show-tbls
local-show-tbls: ## local: show tables
	mysql -u root -ppass_root -h localhost -P 9998 --protocol=tcp -e "show tables from teamn_db;"

.PHONY: local-show-clms
local-show-clms: ## local: show columns
	mysql -u root -ppass_root -h localhost -P 9998 --protocol=tcp -e "show columns from teamn_db.${TABLE_NAME};"

.PHONY: local-show-tz
local-show-tz: ## local: show timezone
	mysql -u root -ppass_root -h localhost -P 9998 --protocol=tcp -e "show variables like '%time_zone%';"

.PHONY: local-select-mysql-user
local-select-mysql-user: ## local: show users
	mysql -u root -ppass_root -h localhost -P 9998 --protocol=tcp -e "select * from mysql.user\G;"