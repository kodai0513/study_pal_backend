FROM golang:1.23-alpine3.21

WORKDIR /src

COPY . .

RUN apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*

RUN apk update
RUN apk add curl
RUN apk add --no-cache postgresql-client
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN curl -sSf https://atlasgo.sh | sh

# ログに出力する時間をJSTにするため、タイムゾーンを設定
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

RUN go mod tidy
RUN go install github.com/air-verse/air@latest