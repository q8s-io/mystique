# build app
FROM golang:1.14-alpine as builder

ARG CODEPATH

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
	&& apk update \
	&& apk add git \
	&& rm -rf /var/cache/apk/*

WORKDIR $GOPATH/src/$CODEPATH

COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOPROXY=https://mirrors.aliyun.com/goproxy/ go build -o /opt/app app.go

# build server
FROM centos:7.6.1810

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

WORKDIR /opt

COPY ./in.txt /opt/

COPY ./start.sh /opt/
RUN chmod +x /opt/start.sh

COPY --from=builder /opt/app .
RUN chmod +x /opt/app

CMD ["sh", "/opt/start.sh"]

