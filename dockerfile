FROM golang:1.15

# Author
MAINTAINER Project:k8s-feature-test MXuDong <1586793553@qq.com>

# Copy file
WORKDIR /go/src/app

COPY . .

ENV GOPROXY https://goproxy.cn
# the application envs

RUN go get -d -v ./...
RUN go build -o ./app cmd/main.go

CMD ["./app -c ./config/conf.yaml"]