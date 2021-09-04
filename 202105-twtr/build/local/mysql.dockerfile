FROM mysql:8.0.24

# local の場合のみ環境変数として MYSQL_ROOT_PASSWORD を定義
ENV TZ=Asia/Tokyo \
    MYSQL_ROOT_PASSWORD=pass_root

# tzdata は削除しない
RUN apt-get update \
    && apt-get install -y --no-install-recommends tzdata \
    && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && mysql_tzinfo_to_sql /usr/share/zoneinfo/Asia/Tokyo 'Asia/Tokyo' \
    && rm -rf /var/cache/apt/*

EXPOSE 3306