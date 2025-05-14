.PHONY: run
run:
	go run ./cmd/ipmanager run -config ./configs/ipmanager.yaml

.PHONY: migrate
migrate:
	go run ./cmd/ipmanager migrate -config ./configs/ipmanager.yaml

.PHONY: test-data
test-data:
	go run ./cmd/ipmanager test-data -config ./configs/ipmanager.yaml

.PHONY: test-client
test-client:
	go run ./test/grpc_client

.PHONY: build
build:
	cd web && yarn build-only
	go run ./tools/vfsgen
	go build -o ./build/ipmanager ./cmd/ipmanager

.PHONY: gen
gen:
	protoc --proto_path=./api/proto \
	   --go_out=paths=source_relative:./api/pb \
	   --go-grpc_out=paths=source_relative:./api/pb \
	   ./api/proto/*.proto

	wire ./internal/app/di

	cd web && yarn gen