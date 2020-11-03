# Comment Service

评论服务

用下面的命令生成

```
micro new --namespace=go.micro --type=service comment
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.service.comment
- Type: service
- Alias: comment

## Dependencies

Micro  服务依赖服务发现。模式是 MDNS，一个零配置的系统。

如果你需要一个灵活的多主机设置，我们推荐 etcd。

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

为了方便，这里有个 Makefile 文件

编译二进制

```
make build
```

运行服务
```
./comment-service
```

构建一个 docker image
```
make docker
```
