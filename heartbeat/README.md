# Heartbeat

演示使用使用服务发现心跳

## Rationale

服务在启动时使用服务发现注册，在关闭时取消注册。有时这些服务可能会意外死亡或被强行杀死或面临瞬态网络问题。在这些情况下，陈旧的节点将留在服务发现中。
如果服务被自动删除就好了。

## Solution

出于这些原因，Micro 支持 register TTL 和 register 间隔的选项。TTL 指定服务能存在于服务发现多久，过期后将被删除。
Interval 是服务应该重新注册的时间。

这些选项在 go-micro 中可用，在 micro toolkit 中作为 flag

## Toolkit

```
micro --register_ttl=30 --register_interval=15 api
```

这个例子中我们设置了 ttl 为 30 秒，15 秒重新注册。

## Go Micro

声明 micro service 时，可以传入一个 time.Duration options。

```
service := micro.NewService(
	micro.Name("com.example.srv.foo"),
	micro.RegisterTTL(time.Second*30),  // 30 秒内存在
	micro.RegisterInterval(time.Second*15), // 15 秒重新注册
)
```
