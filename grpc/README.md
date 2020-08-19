# attend_manager_grpc

grpcを用いたビジネスロジック用のサーバー  
APIの受け口はapiに作成して、処理の実態（演算とDBの更新）はこちらでおこなう  

## コマンド
- pbファイルの出力  
`protoc --go_out=plugins=grpc:./ proto/check.proto`  

grpc
```sh
# grpcurlをインストール！
go get github.com/fullstorydev/grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl
# protocbufferをインストール！
go get github.com/protocolbuffers/protobuf-go
go install github.com/protocolbuffers/protobuf-go
```

- Serviceやメソッドの一覧を取得  
`grpcurl -plaintext localhost:50051 list`  
`grpcurl -plaintext localhost:50051 list Check`  

詳細情報は下記で取得できる  
`grpcurl -plaintext localhost:50051 describe Check`  

- grpcをたたく
`grpcurl -plaintext localhost:50051  proto.Check.HealthCheck`(proto.Check/HealthCheckでも可)