FROM alpine:latest

# Define the project name | 定义项目名称
ARG PROJECT=cmdb
# Define the config file name | 定义配置文件名
ARG CONFIG_FILE=cmdb.yaml
# Define the author | 定义作者
ARG AUTHOR="example@example.com"

LABEL org.opencontainers.image.authors=${AUTHOR}

WORKDIR /app
ENV PROJECT=${PROJECT}
ENV CONFIG_FILE=${CONFIG_FILE}

COPY ./${PROJECT}_rpc ./
COPY ./etc/${CONFIG_FILE} ./etc/

EXPOSE 13134

ENTRYPOINT ./${PROJECT}_rpc -f etc/${CONFIG_FILE}