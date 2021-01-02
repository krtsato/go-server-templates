FROM golang:1.15.6-alpine3.12

ARG APP_ENV

ENV APP_NAME=gin-gorm-logrus-basic \
    APP_DIR=/var/app/${APP_NAME} \
    LOG_DIR=/var/log/${APP_NAME} \
    APP_ENV=${APP_ENV} \
    TZ=Asia/Tokyo

WORKDIR ${APP_DIR}

COPY ./ ./

# コンテナ内で環境変数 TZ を使う場合 tzdata は削除しない
RUN apk update \
  && apk add --no-cache tzdata curl mysql-client \
  && rm -rf /var/cache/apk/* \
  && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
  && go get golang.org/x/tools/cmd/goimports \
  && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin \
  && curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin \
  && go mod download \
  && mkdir -p ${LOG_DIR}

CMD ["air", "-c", "./configs/.air.toml"]