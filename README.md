# Juno


---
>  juno是一套轻量级 Go 微服务框架，包含大量微服务相关框架及工具。


## 特点
1. 灵活：仅仅搭建了项目框架，每个部分都可以用类似组件替代，组件升级不依赖于框架升级；
1. 简单：不过度设计，代码平实简单；
1. 通用：通用业务开发所需要的功能；
1. 高效：提高业务迭代的效率；
1. 健壮：通过良好的基础库设计，减少错用；
1. 高性能：性能高，但不特定为了性能做 hack 优化，引入 unsafe ；
1. 扩展性：良好的接口设计，来扩展实现，或者通过新增目录来扩展功能；
1. 容错性：为失败设计，大量引入对 SRE 的理解，鲁棒性高；
1. 工具链：包含大量工具链，比如代码生成，lint 工具等等；


## 创建项目

### 版本需求
需要使用 go v1.16 以上版本；


### 环境准备
需要安装好对应的依赖环境，以及工具：

- [go](https://golang.org/dl/) 
- [protoc](https://github.com/protocolbuffers/protobuf)
- [protoc-gen-go](https://github.com/protocolbuffers/protobuf-go)

开启GO111MODULE
go env -w GO111MODULE=on


### 安装


```

curl -L -o juno-cli https://
chmod +x juno-cli && mv juno-cli /usr/local/bin/

```

### 创建项目



```
# 创建项目模板
juno-cli new helloworld
cd helloworld
# 拉取项目依赖
go mod download

# 生成service模板
juno-cli service hello

# 生成endpoint模板
juno-cli endpoint hello

# 生成http模板
juno-cli http hello

# 生成grpc模板
juno-cli grpc hello

```

## 项目结构

![项目](http://ihs.joker.org.cn/img/20210705115702.png)



### 目录结构如下

```
├── Dockerfile                      
├── LICENSE
├── README.md
├── api  --- 为外部提供的服务（http,grpc）
│   ├── grpc ---grpc服务
│   │   ├── grpc.go  ---提供grpc实例  
│   │   ├── protos ---proto 文件以及生成的代码
│   │   │   ├── greeting.pb.go
│   │   │   ├── greeting.proto
│   │   │   └── greeting_grpc.pb.go
│   │   └── service ---根据proto文件实现的具体方法
│   │       ├── greeting.go
│   │       ├── greeting_test.go
│   │       └── service.go
│   └── http ---http服务
│       ├── handle  ---提供具体的http handle  
│       │   └── greeting.go
│       ├── http.go  ---提供http实例
│       ├── middleware  ---http中间件
│       │   └── middleware.go
│       ├── schema  ---http请求校验以及转化为endpoint的入参 
│       │   └── greeting.go
│       └── serialize ---endpoint的出参转化为http的response
│           ├── base.go
│           └── greeting.go
├── cmd  --- 整个项目启动的入口文件
│   ├── serve.go
│   ├── wire.go --- 使用wire来维护依赖注入
│   └── wire_gen.go
├── configs  --- 配置文件
│   └── config.yaml
├── generate.go 
├── go.mod
├── go.sum
├── init
│   ├── config
│   │   └── config.go
│   └── flag
│       └── flag.go
├── internal   ---所有不对外暴露的代码，通常的业务逻辑都在这下面，使用internal避免错误引用
│   ├── constant ---常量
│   │   └── constant.go
│   ├── database ---数据库
│   │   └── database.go
│   ├── endpoint --- 业务逻辑的组装层，类似 DDD 的 domain 层
│   │   ├── endpoint.go
│   │   └── greeting.go
│   ├── model --- 数据模型
│   │   └── greeting.go
│   └── service --- 业务数据访问，包含 cache、db 等封装
│       ├── greeting.go
│       └── service.go
├── main.go ---项目启动
└── pkg  ---公用包
    ├── model ---基础数据模型 
    │   └── model.go
    ├── res ---http返回类型
    │   └── response.go
    └── util ---工具包
        └── util.go
```


### 启动项目


```

# 先修改configs/config.yaml 中数据库连接的配置
go run main.go

```


### awesome 工具

#### air 代码热更新
代码变更，自动热更新重启。推荐使用 [github.com/cosmtrek/air]([https://github.com/cosmtrek/air])
`.air.conf`文件是一份默认配置，若已安装 air ，直接在路径下运行 `air` 即可



####  golangci-lint 代码静态检查
```
  // brew
  brew install golangci-lint
```

####  pre-commit  git commit hook

```
  // brew
    brew install pre-commit
  // python
  pip3 install pre-commit
  
  // install
   pre-commit install
  
```
