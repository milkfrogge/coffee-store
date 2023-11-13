start-product-server:
	mkdir -p ./build/product/
	go build -o ./build/product/server.exe ./cmd/productServer/
	./build/product/server.exe

generate:
	make generate-product-api

generate-product-api:
	mkdir -p pkg/product_v1
	protoc --proto_path api/product_v1 \
	--go_out=pkg/product_v1 --go_opt=paths=source_relative \
	--go-grpc_out=pkg/product_v1 --go-grpc_opt=paths=source_relative \
	api/product_v1/category.proto
	protoc --proto_path api/product_v1 \
	--go_out=pkg/product_v1 --go_opt=paths=source_relative \
	--go-grpc_out=pkg/product_v1 --go-grpc_opt=paths=source_relative \
	api/product_v1/product.proto

migrate-product:
	migrate -database postgres://product:product@localhost:5432/product?sslmode=disable -path db/migrations/product up
