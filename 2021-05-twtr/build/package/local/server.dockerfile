
FROM golang:1.16.3-alpine3.13

ARG APP_ENV=local
ARG APP_NAME=twtr

ENV TZ=Asia/Tokyo \
    GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0 \
    PORT=9999 \
    APP_ENV=${APP_ENV} \
    APP_NAME=${APP_NAME} \
    APP_DIR=/var/app/${APP_NAME} \
    LOG_DIR=/var/log/${APP_NAME}

WORKDIR ${APP_DIR}

COPY ./ ./

RUN apk update \
  && apk add --no-cache tzdata curl \
  && rm -rf /var/cache/apk/* \
  && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
  && curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin \
  && go mod download \
  && mkdir -p ${LOG_DIR}

EXPOSE ${PORT}

CMD sh -c "./scripts/twtrdb/wait-for.sh mysql:3306 && air -c ./configs/.air.toml"