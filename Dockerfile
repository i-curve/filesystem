FROM ubuntu:22.04

LABEL maintainer="i-curve" email="i-curve@qq.com"

EXPOSE 8000-8001
COPY filesystem /usr/bin

CMD ["filesystem"]
