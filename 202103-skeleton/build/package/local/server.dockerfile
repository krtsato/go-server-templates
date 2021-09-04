FROM golang:1.15.7-alpine3.13

ARG APP_ENV=local
ARG APP_NAME=app-name

ENV TZ=Asia/Tokyo \
    GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0 \
    APP_ENV=${APP_ENV} \
    APP_NAME=${APP_NAME} \
    APP_DIR=/var/app/${APP_NAME} \
    LOG_DIR=/var/log/${APP_NAME}

WORKDIR ${APP_DIR}

COPY ./ ./

# tzdata は削除しない
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