# demo-go
## 介绍
    'demo-go'是一个自定义go-脚手架
## 目录结构
```
demo-go
├── app
│   └── api //对外api层
│   └── controller //控制器层 
│   └── service //业务代码层   
│   └── dao //数据库访问层 
│   └── model //模型层
├── common //公共代码层
│   └── config // 配置
│      └── config.go // 配置文件
│      └── config.yaml // 配置文件定义
│   └── global // 全局设置
│      └── global.go // 全局变量
│   └── router // 路由层
│      └── router.go // 路由设置
│   └── log // 日志设置
│      └── logger.go // 日志类
│   └── utils // 工具层
│      └── utils.go // 工具类
│   └── middleware // 中间件层
│      └── logger.go // 日志中间件
│      └── cors.go // 跨域中间件
│   └── initall // 初始化函数
│      └── initall.go // 初始化各类服务
├── logs // 日志文件记录层
│      └── access_日期.log // 完整请求栈日志
│      └── error_日期.log // 错误日志
│      └── slow_query_日期.log // 慢查询日志
├── database // 数据库连接层
│   └── db.go // 数据库连接类
│   └── drivers // 数据库驱动层(兼容自定义驱动)
│      └── mysql.go // mysql驱动
│      └── dm.go  // 达梦数据库驱动
│      └── kingbase.go  // 人大金仓数据库驱动
│      └── postgres.go  // postgres数据库驱动
│      └── ... // 其他数据库驱动
├── main.go // 程序入口
├── go.mod
├── go.sum
├── .gitignore
└── README.md
``` 
