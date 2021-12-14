all: get lint test

lint:
	go fmt ./...
get:
	go get ./...
test:
	go test ./...
upgrade:
	go list -m -u all
