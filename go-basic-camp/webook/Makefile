.PHONY: docker port-forward usage
build-go-docker-dev:
	rm webook-server || true
	GOOS=linux GOARCH=amd64 go build -o webook-server main.go
	docker build -t noahouo/webook-server-dev:${VERSION} .
build-go-docker-k8s:
	rm webook-server || true
	GOOS=linux GOARCH=amd64 go build -tags=k8s -o webook-server main.go
	docker build -t noahouo/webook-server-k8s:${VERSION} .
port-forward:
	kubectl port-forward service/webook-server 8099:8099 &
	kubectl port-forward service/webook-mysql 13306:13306 &
	kubectl port-forward service/webook-redis 16379:16379 &
usage:
	@echo "Usage:"
	@echo "  make build-go-docker-dev: build go binary and docker image for dev"
	@echo "  make build-go-docker-k8s: build go binary and docker image for k8s"
	@echo "  make port-forward: port forward for webook-server, webook-mysql and webook-redis"
	@echo "  make usage: show usage"
