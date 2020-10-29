# Meta API

这个例子使用 metadata handler 使用 micro API。

这将允许我们写标准的 go-micro 服务，然后通过服务发现 metadata 设置 handler/endpoint。

## Usage

运行 micro API

```
micro api
```

运行这个例子。注意在注册 handler 时候的 endpoint metadata。

```
go run meta.go
```

POST 请求 /example 这将调用 go.micro.api.example Example.Call

```
curl -H 'Content-Type: application/json' -d '{"name": "john"}' "http://localhost:8080/example"
```

POST 请求 /foo/bar 这样调用 go.micro.api.example Foo.Bar

```
curl -H 'Content-Type: application/json' -d '{}' http://localhost:8080/foo/bar
```
