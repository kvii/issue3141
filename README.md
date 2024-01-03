# issue 3141

[kratos issue 3141](https://github.com/go-kratos/kratos/issues/3141) 复现工程。

版本：

* kratos: v2.7.2
* protoc-gen-go: v1.32.0
* protoc: v4.25.1

改动部分：
* [proto 文件](./api/helloworld/v1/greeter.proto)
* [service 逻辑](./internal/service/myservice.go)

复现方法：

1. `kratos run`
2. `curl http://localhost:8000/my-api/1 -X POST -H 'Content-Type: application/json' -d '{"msg":[{}]}'`
3. 观察控制台输出

结果正常打印了所有字段。

```sh
id:"1"  msg:{}
```