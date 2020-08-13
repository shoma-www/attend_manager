clean:
	rm -f ./api/api
	rm -f ./grpc/grpc
	rm -Rf ./api/proto

cp-proto:
	rm -Rf ./api/proto
	cp -r ./grpc/proto ./api

build: clean
	go build -o api ./api/main.go
	go build -o grpc ./grpc/main.go
