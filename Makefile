install:
	go mod download
	CGO_ENABLED=0 sudo go build -o /usr/local/bin/go-dot




