FROM golang:1.16.3-alpine3.13

ARG APP_ENV=local
ARG APP_NAME=tools

ENV TZ=Asia/Tokyo \
    GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0 \
    PORT=9999 \
    APP_ENV=${APP_ENV} \
    APP_NAME=${APP_NAME} \
    APP_DIR=/var/app/${APP_NAME}

WORKDIR ${APP_DIR}

COPY ./ ./

RUN apk update \
  && apk add --no-cache curl \
  && rm -rf /var/cache/apk/* \
  && go get golang.org/x/tools/cmd/goimports \
  && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin \
  && go mod download