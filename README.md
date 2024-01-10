# issue 3141

[kratos issue 3141](https://github.com/go-kratos/kratos/issues/3141) 复现工程。

版本：

* kratos: v2.7.2
* protoc-gen-go: v1.32.0
* protoc: v4.25.1

改动部分：
* [proto 文件](./api/helloworld/v1/greeter.proto)
* [service 逻辑](./internal/service/myservice.go)
* [client 测试代码](./internal/data/client_test.go)

客户端有 bug，复现方法：

1. `kratos run`
2. `go test -run TestHttpClient ./internal/data` （或者用 IDE 打开 client 测试代码手动点击测试）。
3. 观察控制台输出。

本来应该能正常请求的，结果 404：

```sh
--- FAIL: TestHttpClient (0.00s)
    client_test.go:27: error: code = 404 reason =  message =  metadata = map[] cause = proto: syntax error (line 1:1): unexpected token 404
FAIL
FAIL    github.com/kvii/issue3141/internal/data 2.620s
FAIL
```

看服务端控制台日志。实际上请求的 url 是 "/my-api/" 而非 "/my-api/id"，所以 404 了。

```sh
INFO ts=2024-01-10T15:41:11+08:00 caller=server/http.go:38 service.id=192.168.1.66 service.name= service.version= trace.id= span.id= module=http msg=request method=POST url=/my-api/
```

服务端是正常的，复现方法：

1. `kratos run`
2. `curl http://localhost:8000/my-api/1 -X POST -H 'Content-Type: application/json' -d '{"msg":[{}]}'`
3. 观察控制台输出

结果正常打印了所有字段。

```sh
id:"1"  msg:{}
```