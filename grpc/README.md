# attend_manager_grpc

grpcを用いたビジネスロジック用のサーバー  
APIの受け口はapiに作成して、処理の実態（演算とDBの更新）はこちらでおこなう  

## pbファイルの出力
`protoc --go_out=plugins=grpc:./ proto/check.prot`