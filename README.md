## 基于gin框架的项目框架
### 项目目录结构
```text
├── README.md  // 项目概述
├── bootstrap  // 快速启动（需要初始化的内容）
│   ├── cache  // 缓存
│   │   ├── init.go  // 缓存初始化
│   │   └── redisConn.go  // redis缓存连接池
│   ├── database  // 数据库
│   │   ├── init.go  // 数据库连接池初始化
│   │   ├── mongoConn.go  // mongo连接池
│   │   └── mysqlConn.go  // mysql连接池
│   └── logger  // 日志
│       └── init.go  // 日志初始化
├── config  // 所有配置项
│   ├── config.go  // config初始化
│   ├── global.yaml  // 全局配置
│   ├── logConfig.go  // log配置
│   ├── mongoConfig.go  // mongo数据库配置
│   ├── mysqlConfig.go  // mysql数据库配置
│   └── redisConfig.go  // redis数据库配置
├── controllers  // 业务controllers
│   └── baseController
│       └── baseController.go
├── go.mod  
├── go.sum
├── log  // 日志记录
│   └── app.log
├── middleware  // 中间件
│   └── loggerWare.go
├── models  // model
│   └── baseModel
│       └── baseModel.go
├── routers  // 路由
│   └── routers.go
├── server.go  // 启动文件
├── static  // 静态文件
├── tests  // 测试
└── utils  // 工具包
    └── logger  // 日志工具包
        └── zapLogger.go
```
