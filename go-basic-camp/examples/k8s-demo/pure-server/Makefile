.PHONY: docker port-forward usage
build-go-docker-dev:
	rm pureserver || true
	GOOS=linux GOARCH=amd64 go build -o pureserver main.go
	docker build -t noahouo/pure-server-dev:${VERSION} .
build-go-docker-k8s:
	rm pureserver || true
	GOOS=linux GOARCH=amd64 go build -tags=k8s -o pureserver main.go
	docker build -t noahouo/pure-server-k8s:${VERSION} .
port-forward:
	kubectl port-forward service/pureserver 8999:8999 &
	kubectl port-forward service/pure-mysql 11306:11306 &
	kubectl port-forward service/pure-redis 11379:11379 &
usage:
	@echo "Usage:"
	@echo "  make build-go-docker-dev: build go binary and docker image for dev"
	@echo "  make build-go-docker-k8s: build go binary and docker image for k8s"
	@echo "  make port-forward: port forward for pureserver, pure-mysql and pure-redis"
	@echo "  make usage: show usage"
