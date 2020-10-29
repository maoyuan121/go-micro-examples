# gRPC Config Server

这个例子实现一个 grpc 配置服务

## Get Started

### Run Server

```bash
go run srv/main.go
```

### Run Client

```bash
go run client/main.go
```

### Edit Config

Change values in srv/conf/micro.yml

```bash
micro:
  name: Micro
  version: 1.0.0
  message: hello
```

to 

```bash
micro:
  name: Micro
  version: 1.0.0
  message: hello john
```

The output from watching config after an edit

```bash
2019/04/28 10:57:15 Watch changes: {"message":"hello john","name":"Micro","version":"1.0.0"}
``` 
