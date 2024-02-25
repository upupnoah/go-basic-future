# 基础镜像
FROM ubuntu:22.04

# 编译应该交给 CI/CD 工具去做, 这里只是一个示例
# 把编译后的打包进来镜像, 放到工作目录 /app
COPY ./pureserver /app/pureserver
WORKDIR /app

ENTRYPOINT [ "/app/pureserver" ]