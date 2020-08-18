# build builder
FROM golang:1.14-alpine as builder

ARG CODEPATH

RUN apk add git && rm -rf /var/cache/apk/*

WORKDIR $GOPATH/src/$CODEPATH

COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOPROXY=https://goproxy.cn go build -o /opt/mystique app.go

# build server
FROM centos:7.6.1810

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

WORKDIR /opt

COPY --from=builder /opt/mystique .
COPY ./configs/ ./configs/
RUN chmod +x /opt/mystique

ENTRYPOINT ["/opt/mystique"]
CMD ["-conf", "/opt/configs/pro.yaml"]
