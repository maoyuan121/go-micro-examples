# Mocking

这个例子演示如何 mock helloworld 服务

根据 proto 生成的代码包含了一个 `Service` 接口。可以很容易的 mock。

```go
type GreeterService interface {
	Hello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}
```

这个 `GreeterService` 可以很容易的使用 mock 来代替。

## Mock Client

```go
type mockGreeterService struct {
}

func (m *mockGreeterService) Hello(ctx context.Context, req *proto.Request, opts ...client.CallOption) (*proto.Response, error) {
        return &proto.Response{
                Greeting: "Hello " + req.Name,
        }, nil
}

func NewGreeterService() proto.GreeterService {
        return new(mockGreeterService)
}
```

## Use Mock

在测试环境中，我们可以使用 mock client。

```go
func main() {
	var c proto.GreeterService

	service := micro.NewService(
		micro.Flags(&cli.StringFlag{
			Name: "environment",
			Value: "testing",
		}),
	)

	service.Init(
		micro.Action(func(ctx *cli.Context) error {
			env := ctx.String("environment")
			// use the mock when in testing environment
			if env == "testing" {
				c = mock.NewGreeterService()
			} else {
				c = proto.NewGreeterService("helloworld", service.Client())
			}
                        return nil
		}),
	)

	// call hello service
	rsp, err := c.Hello(context.TODO(), &proto.Request{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Greeting)
}
```
