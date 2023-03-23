## 微服务测试项目
本项目包含了客户端（client）和服务端（server）。集成了 gRPC 和 kafka，项目运行时两者可以同时使用。

测试目录：
- gk文件夹: 集成 gRPC 和 kafka，并完成 gRPC 调用，完成 kafka 消费
- protos文件夹：proto 文件所在目录
- protobuf文件夹：生成的 pb 文件

## 构建镜像
`docker-compose -f ./docker-compose.yml up -d`

## 环境
- go 1.18.3
- protoc 3.20.3

## 测试
### 进入容器
开启 2 个窗口，都执行以下命令：
```shell
# docker exec -it go-micro /bin/sh
```

### 测试 gRPC
服务端：
```shell
# go run server/cmd/gk/main.go
```

客户端：
```shell
# go run client/cmd/helloworld/main.go
```

### 测试 kafka
PHP 生成一条消息
```shell
GET http://local.book.com/produce2
```

Go 消费消息
```shell
# go run server/cmd/gk/main.go
```
