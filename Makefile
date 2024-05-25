build:
	go build cmd/frontier/frontier.go

install:
	go install cmd/frontier/frontier.go

test:
	go test -v ./...
