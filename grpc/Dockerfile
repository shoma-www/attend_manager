
#build stage
FROM golang:alpine AS dev

LABEL Name=attend_manager_grpc Version=0.0.1
ENV GO111MODULE=on
WORKDIR /go/app
COPY . .
WORKDIR /go/app/grpc
RUN apk add --no-cache git \
    && go build -o grpc \
    && go get github.com/githubnemo/CompileDaemon

# 時間をJSTに変更する
RUN apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*

ENTRYPOINT CompileDaemon --build="go build -o grpc" --command="./grpc"
EXPOSE 50051
