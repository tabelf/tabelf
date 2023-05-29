
# 1.目录结构

```
├── cmd #启动程序
│   └── actions  
├── deploy # 部署文件
│   ├── docker  
│   └── k8s  
├── gen #自动生成的代码
│   ├── entschema  
│   └── swagger  
├── service #
│   ├── api #http接口
│   │   ├── apis #接口的定义文件
│   │   ├── etc  
│   │   ├── internal #业务逻辑部分
│   │   │   ├── config  
│   │   │   ├── handler #框架生成的路由对象
│   │   │   │   ├── global  
│   │   │   │   └── swagger  
│   │   │   ├── logic #详细的接口逻辑
│   │   │   │   ├── global  
│   │   │   │   └── swagger  
│   │   │   ├── middleware #http拦截器
│   │   │   ├── svc  
│   │   │   └── types #自动生成的api参数
│   │   ├── models #数据模型操作层
│   │   └── pkg #一些工具/sdk包
│   │   ├── swagger  
│   │   └── utils  
│   ├── app #应用的配置client之类
│   ├── cron #定时任务
│   │   └── tasks  
│   ├── extensions #一些扩展lib
│   ├── rpc #rpc接口
│   │   ├── service #详细的接口逻辑
│   │   └── binghuang.go #rpc接口入口
│   └── testutils #
└── spec #定义的地方
│ └── schema

└── config.yml #配置信息
```

# 2.http 接口

* 在 `service/api/apis/`目录的api文件中定义接口路由和参数

* 在 `service/api/binghuang.api` 文件中 `import` 该文件

* 运行 `make api` 生成接口文件

* 在 `service/api/internal/logic/` 目录的 `logic` 文件中编写具体逻辑

# 3.grpc 接口

* 在 [BE / inf / justitia](https://be/inf/justitia) 项目中定义proto文件，并生成对应的go代码

* 在项目中引入最新的 `justitia` 包

* 在 `service/rpc/service/` 目录中实现proto中定义的接口具体逻辑

* 在 `service/rpc/binghuang.go` 文件中引入接口的 handler 和具体实现

# 4.增加数据库实体

* 生成 schema `make ARGS="YourModel" entinit`

* 在 `spec/schema/yourmodel.go` 定义数据库字段
- 运行 `make migrations` 生成 ORM 代码

# 4.定时任务

* 在 `service/cron/tasks/` 目录中增加业务逻辑

* 在 `service/cron/cron.go`  文件中的 `jobs` 数组中定义定时规则

# 5.异步任务


# 部署到linux流程
1. 进入到 main.go 文件所在到目录
2. 执行 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o binghuang main.go
   // -o 表示输出到指定文件
3. 在 linux 要赋予权限， chmod 777 binghuang
4. 执行运行 ./binghuang start_http --env prod