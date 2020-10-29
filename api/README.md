# API

这个仓储包含了使用 micro api 的例子。

## Overview
[micro api](https://github.com/micro/micro/tree/master/api) 是一个 API 网关，它提供 HTTP 服务，并基于服务发现动态路由。
                                
micro api 的默认命名空间是 go.micro.api。我们的服务名包含这个作为前缀然后加上一个唯一的名字，例如： go.micro.api.example。 可以通过
flag `--namespace=` 修改命名空间。

micro api 有许多不同 handler，它们允许定义你想要的 api 服务类型。例如下面的例子。handler 可以通过 flag `--handler` 设置。
默认的 handler 是 "rpc"

## Contents

- api - 提供整个 http 请求头和请求的 rpc handler
- proxy - 使用 api 作为 http 反向代理
- rpc - 向 go-micro 应用程序发出 rpc 请求
- meta - 通过代码中的配置指定要使用的 handler

## Request Mapping

### API/RPC

Micro 将 http path 映射到 rpc 服务。可以在下面看到映射表。

api 默认的命名空间为 **go.micro.api**，你可以通过使用  `--namespace` 来设置你自己的命名空间。

URLs 的映射样例：

Path	|	Service	|	Method
----	|	----	|	----
/foo/bar	|	go.micro.api.foo	|	Foo.Bar
/foo/bar/baz	|	go.micro.api.foo	|	Bar.Baz
/foo/bar/baz/cat	|	go.micro.api.foo.bar	|	Baz.Cat

后两个 path fragment 对应服务和方法，前面的对应服务名。

API 版本号 URL 可以很容易的映射到服务名：

Path	|	Service	|	Method
----	|	----	|	----
/foo/bar	|	go.micro.api.foo	|	Foo.Bar
/v1/foo/bar	|	go.micro.api.v1.foo	|	Foo.Bar
/v1/foo/bar/baz	|	go.micro.api.v1.foo	|	Bar.Baz
/v2/foo/bar	|	go.micro.api.v2.foo	|	Foo.Bar
/v2/foo/bar/baz	|	go.micro.api.v2.foo	|	Bar.Baz

### Proxy Mapping

使用 `--handler=http` 启动 API 将反向代理请求到后台服务： 

Example

Path	|	Service	|	Service Path
---	|	---	|	---
/greeter	|	go.micro.api.greeter	|	/greeter
/greeter/:name	|	go.micro.api.greeter	|	/greeter/:name
