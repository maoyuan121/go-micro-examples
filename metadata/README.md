# Metadata

演示发送 metadata/headers。

发给 micro api 的 http 请求头会自动转换成 metadata。

## Contents

- **srv** - RPC 服务，它获取 metadata
- **cli** - RPC 客户端，它调用 服务

## Run Service

Start go.micro.srv.greeter
```shell
go run srv/main.go
```

## Client

Call go.micro.srv.greeter via client
```shell
go run cli/main.go
```

