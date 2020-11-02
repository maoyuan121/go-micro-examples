# noproto

演示在没有 protobuf 的情况下使用 micro。


使用 micro 和标准 go 类型，并使用 json codec 进行 marshalling。服务有多个 codec，使用 `Content-Type` 请求头来确定使用哪个。
client 使用 `Content-Type: application/json`。因为我们可以将标准 Go 类型转为 json，所以不需要生成代码或使用 protobuf。
                                           

## Contents

- main.go - 是一个 micro greeter  服务
- client - 是一个 micro json client

## Run the example

Run the service

```shell
go run main.go
```

Run the client

```shell
go run client/main.go
```
