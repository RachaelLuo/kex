FROM golang:alpine AS builder
WORKDIR /usr/src/app
ENV GOPROXY https://goproxy.io,direct
COPY ./go.mod ./
COPY ./go.sum ./
# RUN go mod download
COPY . .
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && apk add --no-cache upx ca-certificates tzdata
RUN CGO_ENABLED=0 go build -tags=jsoniter -ldflags "-s -w" -o kex . && upx kex

FROM alpine as runner
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/src/app/kex /opt/app/
