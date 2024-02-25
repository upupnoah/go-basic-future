.PHONY: docker
docker:
	rm main || true
	GOOS=linux GOARCH=amd64 go build -o pureserver main.go
	docker build -t noahouo/pure-server:${VERSION} .
# USAGE: make docker VERSION=v0.0.2