FROM ubuntu:22.04

LABEL maintainer="i-curve" email="i-curve@qq.com"

COPY filesystem /usr/bin

CMD ["filesystem"]
