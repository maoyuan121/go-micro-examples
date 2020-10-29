# RPC API

这个例子将 micro api 作为一个 RPC handler。

这种模式的 api 能让你提供 http 服务，通过 RPC 路由到标准 go-micro 服务。
          
这个模式下的 micro api 只支持 POST 方法，而且 content-type 不是 application/json 就是 application/protobuf。

## Usage

使用 rpc handler 运行 micro API

```
micro api --handler=rpc
```

运行这个例子

```
go run rpc.go
```

POST 请求 /example/call 将调用 go.micro.api.example Example.Call

```
curl -H 'Content-Type: application/json' -d '{"name": "john"}' "http://localhost:8080/example/call"
```

POST 请求 /example/foo/bar 将调用 go.micro.api.example Foo.Bar 

```
curl -H 'Content-Type: application/json' -d '{}' http://localhost:8080/example/foo/bar
```
