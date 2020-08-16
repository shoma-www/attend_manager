
#build stage
FROM golang:alpine AS dev

LABEL Name=attend_manager_api Version=0.0.1
ENV GO111MODULE=on
WORKDIR /go/app
COPY . .
WORKDIR /go/app/api
RUN apk add --no-cache git \
    && go build -o api \
    && go get github.com/githubnemo/CompileDaemon

# 時間をJSTに変更する
RUN apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*

ENTRYPOINT CompileDaemon --build="go build -o api" --command="./api"
EXPOSE 8080
