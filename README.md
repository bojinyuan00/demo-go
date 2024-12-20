# demo-go
## 介绍
    'demo-go'是一个自定义go-脚手架
    适合go语言初学者学习使用，适合java、python、php等开发者，熟悉面向对象编程、容器依赖注入思想的开发者学习与使用。
    主要用于快速搭建go项目的脚手架，目录层级清晰分明，包含了常用的功能模块，如：
    - 日志记录 logrus （完整请求栈记录，错误日志记录，慢查询日志记录）
    - 统一错误处理（错误日志记录，统一错误处理）
    - 数据库连接 (支持redis,mysql,达梦,人大金仓,postgres等数据库。。支持自定义驱动，支持扩展)
    - 路由配置  gin框架 路由映射 路由模块分组，结构清晰
    - 自定义配置 (yaml格式) viper控制文件读取
    - 容器化管理，依赖注入 (wire)-自动生成依赖注入代码，只用实例化最上层类即可
    - 中间件 middleware 自定义配置   
    - - 跨域设置 cors   
    - - xss攻击防御   --待完善
    - - csrf攻击防御  --待完善
    - 自定义服务初始化   initall一切 初始化各类服务
    - 单元测试 test
    - 持续集成 cicd (github action)--待完善
    - 其他功能待补充...
## 目录结构
```
demo-go
├── app
│   └── api //对外api层
│   └── controller //控制器层 
│   └── service //业务代码层   
│   └── dao //数据库访问层 
│   └── model //模型层
│   └── provider //容器化依赖管理层
│      └── provider.go // 声明依赖关系，自动生成依赖注入代码，只用实例化最上层类即可
├── config // 配置
│     └── config.go // 配置文件
│     └── config.yaml // 配置文件定义
├── common //公共代码层
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
│   └── initall // 初始化层
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
├── test // 测试代码层
│   └── test.go // 测试代码
├── cicd // 持续集成工具层
├── main.go // 程序入口
├── go.mod
├── go.sum
├── .gitignore
└── README.md
``` 
