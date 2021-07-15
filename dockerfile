FROM golang:1.16-stretch

# Author
MAINTAINER Project:k8s-feature-test MXuDong <1586793553@qq.com>

WORKDIR /app

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

# Copy file
COPY . .

ENV GOPROXY https://goproxy.cn
# the application envs

RUN go get -d -v ./...
RUN go build  -o app cmd/main.go

CMD ["./app", "-c", "./config/conf.yaml"]