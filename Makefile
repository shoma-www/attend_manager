clean:
	rm -f ./api/api
	rm -f ./grpc/grpc

cp-proto:
	rm -Rf ./api/proto
	cp -r ./grpc/proto ./api

build: clean
	cd api && go build -o api
	cd grpc && go build -o grpc
