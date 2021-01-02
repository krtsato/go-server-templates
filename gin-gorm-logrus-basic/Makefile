# ------------------------------------------------------------ #
# Env: local | dev | prd
# ------------------------------------------------------------ #

# App
APP_ENV=local
APP_NAME=${APP_ENV}-gin-gorm-logrus-basic

# AWS
REGION=ap-northeast-1
ECR=aws-account-id.dkr.ecr.${REGION}.amazonaws.com

# Docker
DOCKER_REPO_URI=$(ECR)/$(APP_NAME)
DOCKER_IMAGE_TAG=latest
DOCKERFILE_PATH=build/package/server.dockerfile

# Makefiles
MAKE_CFN=make -f ./scripts/make/cfn-deploy.mk \
	APP_NAME=${APP_NAME} \
	DOCKER_REPO_URI=${DOCKER_REPO_URI} \
	DOCKER_IMAGE_TAG=${DOCKER_IMAGE_TAG}

MAKE_DOCKER=make -f ./scripts/make/docker-ecr.mk \
	APP_NAME=${APP_NAME} \
	REGION=${REGION} \
	ECR=${ECR} \
	DOCKER_REPO_URI=${DOCKER_REPO_URI} \
	DOCKER_IMAGE_TAG=${DOCKER_IMAGE_TAG} \
	DOCKERFILE_PATH=${DOCKERFILE_PATH}

# ------------------------------------------------------------ #
# Setup
# ------------------------------------------------------------ #

.PHONY: help
help: ## make コマンド一覧
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' ${MAKEFILE_LIST}

.PHONY: lint
lint: ## local: golangci-lint による静的解析
	docker-compose exec server golangci-lint run -c ./configs/.golangci.yml

.PHONY: fmt-import
fmt-import: ## local: goimports のフォーマット
	docker-compose exec server goimports -v -w ./

.PHONY: fmt
fmt: ## local: デフォルトのフォーマット
	docker-compose exec server gofmt -w ./

.PHONY: go-build
go-build: ## dev/prd: デプロイ前にソースコードをビルド
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/server ./cmd/helix-api/{main.go,wire_gen.go}

# ------------------------------------------------------------ #
# Docker for local
# ------------------------------------------------------------ #

.PHONEY: local-up-all
local-up-all: ## local: docker-compose でイメージ作成・コンテナ起動
	docker-compose up --build -d

.PHONEY: local-down-all
local-down-all: ## local: docker-compose ですべてのリソース削除
	docker-compose down -v --rmi all

.PHONEY: local-logs
local-logs: ## local: docker-compose でログ出力
	docker-compose logs -f

.PHONEY: local-exec-sh
local-exec: ## local: docker-compose で sh ログイン
	docker-compose exec server sh

.PHONEY: local-cont-ls
local-cont-ls: ## local: docker container 一覧
	docker container ls -a

.PHONEY: local-img-ls
local-img-ls: ## local: docker image 一覧
	docker image ls -a

.PHONY: local-cont-rm
local-cont-rm: ## local: exited した docker container 削除
	docker container ls -aqf status=exited | xargs docker rm -f

.PHONY: local-img-rm
local-img-rm: ## local: タグ無しの docker image 削除
	docker image ls -aqf dangling=true | xargs docker rmi -f

# ------------------------------------------------------------ #
# Docker for ECR
# ------------------------------------------------------------ #

.PHONY: ecr-login
ecr-login: ## dev/prd: ECR ログイン
	aws ecr get-login-password --region ${REGION} | docker login --username AWS --password-stdin ${ECR}

.PHONEY: dev-docker-build
dev-docker-build: ## dev: Docker イメージをビルド
	${MAKE_DOCKER} docker-build APP_ENV=dev

.PHONY: dev-docker-push
dev-docker-push: ## dev: Docker イメージをプッシュ
	${MAKE_DOCKER} docker-push APP_ENV=dev

.PHONY: dev-docker-image-rm
dev-docker-image-rm: ## dev: Docker イメージを削除
	${MAKE_DOCKER} docker-image-rm APP_ENV=dev

.PHONY: dev-ecr-delete-repo
dev-ecr-delete-repo: ## dev: Docker イメージを持つ ECR Repository を削除
	${MAKE_DOCKER} ecr-delete-repo APP_ENV=dev

# ------------------------------------------------------------ #
# Deploy by CloudFormation
# ------------------------------------------------------------ #

.PHONY: dev-cfn-deploy-stack
dev-cfn-deploy-stack: ## dev: CFn スタックの作成・更新
	${MAKE_CFN} cfn-deploy-stack APP_ENV=dev

.PHONY: dev-cfn-deploy-noexe-changeset
dev-cfn-deploy-noexe-changeset: ## dev: デプロイ前の changeset 確認
	${MAKE_CFN} cfn-deploy-stack APP_ENV=dev CHANGESET_OPTION=--no-execute-changeset

.PHONY: dev-cfn-delete-stack
dev-cfn-delete-stack: ## dev: CFn スタックの削除
	${MAKE_CFN} cfn-delete-stack APP_ENV=dev

.PHONY: dev-aws-logs-tail
dev-aws-logs-tail: ## dev: デプロイ後のログ確認
	${MAKE_CFN} aws-logs-tail APP_ENV=dev

# ------------------------------------------------------------ #
# MySQL client
# ------------------------------------------------------------ #

.PHONY: local-mysql-show-dbs
local-mysql-show-dbs: ## local: show databases
	docker-compose exec server sh -c 'mysql -u root -proot_pwd -h mysql -P 3306 --protocol=tcp -e "show databases;"'

.PHONY: local-mysql-show-tbls
local-mysql-show-tbls: ## local: show tables
	docker-compose exec server sh -c 'mysql -u root -proot_pwd -h mysql -P 3306 --protocol=tcp -e "show tables from gin_gorm_logrus_basic_db"'

.PHONY: local-mysql-select-users
local-mysql-select-users: ## local: show users
	docker-compose exec server sh -c 'mysql -u root -proot_pwd -h mysql -P 3306 --protocol=tcp -e "select * from mysql.user\G"'

.PHONY: local-mysql-dump
local-mysql-dump: ## local: dump DB
	docker-compose exec server sh -c './scripts/mysql/dump.sh gin_gorm_logrus_basic_db'