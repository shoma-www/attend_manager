# attend_manager_grpc

grpcを用いたビジネスロジック用のサーバー  
APIの受け口はapiに作成して、処理の実態（演算とDBの更新）はこちらでおこなう  

## コマンド
### grpc
- pbファイルの出力  
`protoc --go_out=plugins=grpc:./ grpc/proto/check.proto`  

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

```bash
grpcurl -d '{"group_name": "nakamura family", "login_id": "root", "password": "root", "user_name": "中村家"}' \
-plaintext localhost:50051 proto.AttendanceGroup/Create
```

### entの生成
- generate schema
`entc init User`
- generate file
`go generate ./grpc/ent`

### mockの生成
- interfaceからmockを生成
`mockgen -source grpc/service/repository.go -destination grpc/mock_service/repository.go`
