.PHONY: docker port-forward usage mock
build-go-docker-dev:
	rm webook-server || true
	GOOS=linux GOARCH=amd64 go build -o webook-server .
	docker build -t noahouo/webook-server-dev:${VERSION} .
build-go-docker-k8s:
	rm webook-server || true
	GOOS=linux GOARCH=amd64 go build -tags=k8s -o webook-server .
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
mock:
	@mockgen -source=./internal/service/user.go -package=svcmocks -destination=./internal/service/mocks/user.mock.go
	@mockgen -source=./internal/service/code.go -package=svcmocks -destination=./internal/service/mocks/code.mock.go
	@mockgen -source=./internal/repository/code.go -package=repomocks -destination=./internal/repository/mocks/code.mock.go
	@mockgen -source=./internal/repository/user.go -package=repomocks -destination=./internal/repository/mocks/user.mock.go
	@mockgen -source=./internal/repository/dao/user.go -package=daomocks -destination=./internal/repository/dao/mocks/user.mock.go
	@mockgen -source=./internal/repository/cache/user.go -package=cachemocks -destination=./internal/repository/cache/mocks/user.mock.go
	@mockgen -package=redismocks -destination=./internal/repository/cache/redismocks/cmd.mock.go github.com/redis/go-redis/v9 Cmdable
	@go mod tidy