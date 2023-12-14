BINARY_NAME=smartNotify

build:
	@echo "Build project..."
	go build -o ${BINARY_NAME} cmd/main.go

test:
	@echo "Run tests..."
	go test ./...

run: build
	@echo "Run application..."
	./${BINARY_NAME}

clean:
	@echo "Cleanup..."
	go clean
	rm ${BINARY_NAME}

help:
	@echo "Available commands"
	@echo "  make build  - build the project"
	@echo "  make test "
	@echo "  make run  "
	@echo "  make clean  "
	@echo "  make help  "
