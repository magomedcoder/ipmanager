.PHONY: run
run:
	go run ./cmd/ipmanager run -config ./configs/ipmanager.yaml

.PHONY: build
build:
	go run ./tools/vfsgen
	go build -o ./build/ipmanager ./cmd/ipmanager

.PHONY: gen
gen:
	wire ./internal/app/di

	protoc --proto_path=./api/proto \
	   --go_out=paths=source_relative:./api/pb \
	   --go-grpc_out=paths=source_relative:./api/pb \
	   ./api/proto/*.proto