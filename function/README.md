# Function

This is an example of creating a micro function. A function is a one time executing service.

这个例子创建一个 micro function。 function 这种服务执行一次就回推出

## Contents

- main.go - is the main definition of the function
- proto - contains the protobuf definition of the API

## Run function

```shell
while true; do
	github.com/micro/examples/function
done
```

## Call function

```shell
micro call greeter Greeter.Hello '{"name": "john"}'
```
