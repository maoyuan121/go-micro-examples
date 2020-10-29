# Proxy API

这个例子将 micro api 作为一个 http 反向代理。

使用 api 作为一个反向代理，你完全控制在 API 层使用任何语言和库。在这个这个例子中我们使用了 go-web 注册 http 服务。

## Usage

使用 http 反向代理 运行 micro api 

```
micro api --handler=http
```

运行代理 服务

```
go run proxy.go
```

请求 /example/call 它将调用 go.micro.api.example Example.Call

```
curl "http://localhost:8080/example/call?name=john"
```

post 请求 /example/foo/bar 它将调用 go.micro.api.example Foo.Bar

```
curl -H 'Content-Type: application/json' -d '{}' http://localhost:8080/example/foo/bar
```
