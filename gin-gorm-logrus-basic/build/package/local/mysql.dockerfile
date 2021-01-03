FROM mysql:8.0.22

# local の場合のみ環境変数として pwd を定義
ENV TZ=Asia/Tokyo \
    MYSQL_ROOT_PASSWORD=pass_root

# tzdata は削除しない
# fatal errors during processing of zoneinfo directory 回避のため
RUN apt-get update \
    && apt-get install -y --no-install-recommends tzdata \
    && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && mysql_tzinfo_to_sql /usr/share/zoneinfo/Asia/Tokyo 'Asia/Tokyo' \
    && rm -rf /var/cache/apt/*