# 開発ステージ
FROM golang:1.23-bookworm

WORKDIR /go/src/app

# 必要な開発ツールのインストール
RUN apt-get update -qq && \
    apt-get install -y git && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# 依存関係のキャッシュを最適化
COPY go.mod go.sum ./
RUN go mod download

# migration toolのインストール (最新版のgooseを使用)
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

CMD ["/bin/bash"]
