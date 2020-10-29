# API

api 提供整个 http 请求头和请求的 rpc handler

这个例子使用 "api" handler。

这个 api 使用 [api.Request/Response](https://github.com/micro/go-api/blob/master/proto/api.proto) protos。


micro api 请求 handler 为你提供对 http 请求和响应的完全控制，同时仍然利用 RPC 和任何 tansport 插件，使用其他非 http 的协议，如 grpc, nats, kafka。
                                                      
## Usage

运行 micro api

```
micro api --handler=api
```

运行这个例子

```
go run api.go
```


## Calling the service

请求 /example/call，它将调用 go.micro.api.example Example.Call

```
curl "http://localhost:8080/example/call?name=john"
```

POST 请求 /example/foo/bar 它将调用 go.micro.api.example Foo.Bar

```
curl -H 'Content-Type: application/json' -d '{}' http://localhost:8080/example/foo/bar
```

## Set Namespace

使用自定义命名空间运行 micro api

```
micro api --handler=api --namespace=com.foobar.api
```

或者
```
MICRO_API_NAMESPACE=com.foobar.api micro api --handler=api
```

设置命名空间和服务名

```
service := micro.NewService(
        micro.Name("com.foobar.api.example"),
)
```   
