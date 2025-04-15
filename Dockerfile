# ========= 构建阶段 =========
FROM golang:1.24-alpine3.21 AS builder

# 替换为清华镜像源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
    && apk add --no-cache git

WORKDIR /app

COPY go.mod ./
# 设置七牛云代理
RUN go env -w GOPROXY=https://goproxy.cn,direct && go env -w GO111MODULE=on

RUN go mod download

COPY . .

# 静态编译，避免依赖 libc
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server main.go

# ========= 运行阶段 =========
FROM alpine:latest

# 同样替换为清华源，确保 apk 安装如 tzdata 时也快速
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
    && adduser -D -g '' appuser

WORKDIR /app

COPY --from=builder /app/server .

USER appuser

EXPOSE 8080

ENTRYPOINT ["./server"]
