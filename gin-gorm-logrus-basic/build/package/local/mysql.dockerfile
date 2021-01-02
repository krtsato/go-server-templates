FROM mysql:8.0.22

ENV TZ=Asia/Tokyo

RUN apt-get update \
    && apt-get install -y --no-install-recommends tzdata \
    && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && mysql_tzinfo_to_sql /usr/share/zoneinfo \
    && rm -rf /var/cache/apt/*

#  && apt-get purge -y tzdata \