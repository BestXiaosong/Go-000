## 作业
按照自己的构想，写一个项目, 满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。
https://github.com/Mountaincnc/Go-000/tree/main/Week04/homework/internal

homework
├── api
│   └── article
│       └── v1
│           ├── article.pb.go
│           └── article.proto
├── cmd
│   └── week04
│       ├── main.go
│       ├── wire.go
│       └── wire_gen.go
├── go.mod
├── go.sum
├── internal
│   ├── biz
│   │   └── article.go
│   ├── data
│   │   └── article.go
│   ├── pkg
│   │   └── server
│   │       └── server.go
│   └── service
│       └── article.go
└── test
    └── main.go


