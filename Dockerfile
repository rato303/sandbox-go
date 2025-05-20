FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY . .

# go.mod を初期化（モジュール名は任意。ここでは sandbox-go）
RUN go mod init sandbox-go

# アプリケーションをビルド
RUN go build -o app

# 実行用の軽量イメージを構築
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/app .

EXPOSE 8080
CMD ["./app"]
