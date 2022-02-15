bench:
	go test ./... -bench=.

doc:
	godoc -http=localhost:6060

format:
	go fmt ./...

lint:
	golangci-lint run ./...

test:
	go test -race ./...

# utils
precommit: format lint test
