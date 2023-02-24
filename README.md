## 微服务测试项目
本项目包含了客户端（client）和服务端（server）

测试目录：
- gk文件夹: 集成 gRPC 和 kafka，并完成 gRPC 调用，完成 kafka 消费
- protos文件夹：proto 文件所在目录
- protobuf文件夹：生成的 pb 文件

## build 镜像
`docker-compose -f ./docker-compose.yml up -d`

## 环境
- go 1.18.3
- protoc 3.20.3

