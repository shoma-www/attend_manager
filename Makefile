clean:
	rm -f ./api/api
	rm -f ./grpc/grpc

cp-proto:
	rm -Rf ./api/proto
	cp -r ./grpc/proto ./api

build: clean
	go build -o api ./api/...
	go build -o grpc ./grpc/...
