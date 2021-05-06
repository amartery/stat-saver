.PHONY: build
build:
	go build -v ./cmd/statserver/main.go

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: coverage
coverage:
	go test -covermode=atomic -coverpkg=./... -coverprofile=cover ./...
	cat cover | fgrep -v "mock" | fgrep -v "pb.go" | fgrep -v "easyjson" | fgrep -v "start.go" > cover2
	go tool cover -func=cover2

.DEFAULT_GOAL := build