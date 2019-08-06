FROM alpine:3.10.1

ENV GIN_MODE=release

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add curl bash busybox-extras tzdata

COPY httpbasicauth /httpbasicauth

RUN chmod +x /httpbasicauth