# Form

这个例子演示在 API service 中如何访问 form 和 multipart form。

## Contents

- web - 是一个 web 前端，有一个 form
- api - 是一个 api 服务

## Usage

运行 micro api

```
micro api --handler=api
```

运行 micro web

```
micro web
```

运行 api 服务

``` 
go run api/main.go
```

运行 web 服务

```
go run web/main.go
```

浏览 localhost:8082/form
