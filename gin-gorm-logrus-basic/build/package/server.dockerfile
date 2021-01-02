FROM alpine:3.12.3

ARG APP_ENV

ENV APP_NAME='gin-gorm-logrus-basic' \
    APP_DIR=/var/app/${APP_NAME} \
    LOG_DIR=/var/log/${APP_NAME} \
    CNF_DIR=${APP_DIR}/configs \
    APP_ENV=${APP_ENV} \
    TZ='Asia/Tokyo'

COPY ./bin/server ${APP_DIR}/bin/server
COPY ./configs/application.yml ${CNF_DIR}/application.yml

# コンテナ内で環境変数 TZ を使う場合 tzdata は削除しない
RUN apk update \
  && apk add --no-cache tzdata \
  && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
  && rm -rf /var/cache/apk/* \
  && mkdir -p ${LOG_DIR}

WORKDIR ${APP_DIR}
EXPOSE 3000

CMD ["sh", "-c", "${APP_DIR}/bin/server"]