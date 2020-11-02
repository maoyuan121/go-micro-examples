# Options

Go-micro 使用 [functional options](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)。它是一种设计模式，允许在不更改方法签名的情况下添加新选项。

每个 package 都有一个 [Option](https://godoc.org/github.com/micro/go-micro#Option) 类型

```
type Option func(*Options)
```

 [Name](https://godoc.org/github.com/micro/go-micro#Name) Option 函数用来设置服务名。

下面是他的实现。

```
func Name(n string) Option {
	return func(o *Options) {
		o.Server.Init(server.Name(n))
	}
}
```

## Usage


```
import "github.com/micro/go-micro/v2"


service := micro.NewService(
	micro.Name("my.service"),
)
```
