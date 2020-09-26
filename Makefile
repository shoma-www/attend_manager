file = default
clean:
	rm -f ./api/api
	rm -f ./grpc/grpc

cp-proto:
	rm -Rf ./api/proto
	cp -r ./grpc/proto ./api

build: clean
	go build -o api ./api/...
	go build -o grpc ./grpc/...

migrate-%:
	$(eval METHOD := $*)
	migrate -path resource/mysql/migrations -database 'mysql://root:root@tcp(127.0.0.1:3306)/attend?query' $(METHOD)

data-dump:
	$(eval FILE_NAME := "./resource/mysql/dump/atend_data_dump_$(shell date "+%Y%m%d%H%M%S").sql")
	mysqldump -u root -p -h 127.0.0.1 -t attend > $(FILE_NAME)

migrate-file:
	migrate create -ext sql -dir ./resource/mysql/migrations -seq ${file}
