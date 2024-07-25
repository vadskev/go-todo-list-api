run-app:
	go run ./cmd/main.go -config-path=.env
	#go run ./cmd/main.go

lint:
	golangci-lint run ./...

coverage:
	go test -race -cover ./... -coverprofile coverage.out
	go tool cover -html coverage.out

vettest:
	go vet ./...