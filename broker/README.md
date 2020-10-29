# Broker

[broker](https://godoc.org/github.com/micro/go-micro/broker#Broker) 是 PubSub 的接口

貌似默认用的是内存

## Contents

- main.go - uns pub-sub as two go routines for 10 seconds.
- producer - publishes messages to the broker every second
- consumer - consumes any messages sent by the producer
