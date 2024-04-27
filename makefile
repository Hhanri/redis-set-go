build:
	@go build -o bin/fs

run: build
	@./bin/fs --listenAddr :5001

test:
	@go test ./... -v