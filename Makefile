all: get lint test

lint:
	go fmt ./...
get:
	go get ./...
test:
	go test ./...
deps:
	go list -m -u all
