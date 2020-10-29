# Modify Config

这个例子是修改一个文件中配置文件。

Note: We currently only support setting values in memory, not passing down to the source.

## Usage

- example.conf is a toml file format
- modify.go - loads, modifies and writes back the file

```
go run modify.go
```
