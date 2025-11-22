.PHONY: proto
proto:
	@echo "Generating Go code from proto files..."
	@mkdir -p pb
	protoc --proto_path=../proto \
		--go_out=. --go_opt=module=github.com/cloudtv/backend \
		--go-grpc_out=. --go-grpc_opt=module=github.com/cloudtv/backend \
		$$(find ../proto -name '*.proto')
	@echo "✓ Proto generation complete"

.PHONY: proto-install
proto-install:
	@echo "Installing protoc dependencies..."
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: proto-clean
proto-clean:
	@echo "Cleaning generated proto files..."
	@rm -rf pb
	@echo "✓ Clean complete"

