FROM golang:1.18
LABEL maintainer="i-curve" email="i-curve@qq.com"

COPY filesystem /usr/bin
VOLUME [ "/data" ]

CMD ["filesystem", "-c /etc/config.json"]
