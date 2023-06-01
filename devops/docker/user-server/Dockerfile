FROM golang:alpine AS builder

LABEL stage=gobuilder

# 设置 $USER 环境变量
ENV USER=docker

ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY src/common/config/client.yaml /app/config/client.yaml
COPY src/user/config/config.yaml /app/config/config.yaml
RUN go build -ldflags "-s -w" -o /app/main src/user/cmd/main.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai
# 设置 $USER 环境变量
ENV USER=docker

WORKDIR /app
COPY --from=builder /app /app

CMD ["./main"]
