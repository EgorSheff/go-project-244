BIN_PATH = bin/gendiff

build:
	go build -ldflags="-w -s" -gcflags=all="-l -B" -o $(BIN_PATH) ./cmd/gendiff

lint:
	golangci-lint run ./...

test:
	go test -v -race -coverprofile=coverage.out ./...

clean:
	rm bin/* || true # Ignore errors
