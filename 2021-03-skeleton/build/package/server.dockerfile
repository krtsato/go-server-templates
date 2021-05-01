FROM golang:1.15.7-alpine3.13

ARG APP_ENV=dev
ARG APP_NAME=app-name

ENV APP_ENV=${APP_ENV} \
    APP_NAME=${APP_NAME} \
    APP_DIR=/var/app/${APP_NAME} \
    LOG_DIR=/var/log/${APP_NAME} \
    TZ=Asia/Tokyo

COPY ./bin/server ${APP_DIR}/bin/skeleton-api
COPY ./configs/application.yml ${APP_DIR}/configs/application.yml

# コンテナ内で環境変数 TZ を使う場合 tzdata は削除しない
RUN apk update \
  && apk add --no-cache tzdata \
  && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
  && rm -rf /var/cache/apk/* \
  && mkdir -p ${LOG_DIR}

WORKDIR ${APP_DIR}
EXPOSE 9999

CMD ["sh", "-c", "${APP_DIR}/bin/skeleton-api"]