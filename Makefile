install:
	go mod download
	CGO_ENABLED=0 sudo go build -o /usr/local/bin/go-dot
	go-dot completion fish > ~/.config/fish/completions/go-dot.fish



