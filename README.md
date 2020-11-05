# ffly-plus

一款适合于快速开发业务的 Go Gin Demo，主要是提供 API, RPC 服务。

## 总体设计文档

### 需求

`一般以产品的语言描述，这一块可以拷贝产品需求文档中的story list部分`

本项目采用了一系列Golang中比较流行的组件，可以以本项目为基础快速搭建`Restful Web API` 和　`RPC` 服务，主要目的是为了方便大家学习怎么用Golang编写前后端分离的纯后端项目

### 名词解释

`非相关领域内的同学需要看到文档需要提前了解的一些概念性质的东西`

无

### 设计目标

`功能目标和性能目标，功能目标一般是对产品需求的技术描述，性能目标是根据产品给出的数据对性能进行的评估。一般来说，新服务必须要有性能目标一项，性能目标可能会影响设计方案`

### 性能目标

性能目标是新模块文档必不可少的一部分，很多项目对性能影响较大的话，也必须撰写性能目标，性能一般来说可能包含以下部分：

* 日平均请求：一般来自产品人员的评估；
* 平均QPS：日平均请求 除以 4w秒得出，为什么是4w秒呢，24小时化为86400秒，取用户活跃时间为白天算，除2得4w秒；
* 峰值QPS：一般可以以QPS的2~4倍计算；

## 详细设计文档

### 系统架构

`一般来说会有个简单的架构图，并配以文字对架构进行简要说明`

![示意图](tool/ffly-plus.png)

* ffly-plus 主要提供HTTP 和　RPC 服务
* 后端使用MySQL

#### 📗 目录结构

```sh
ffly-plus
├── config ## 配置
├── controller ## API实现,用来读取输入、调用业务处理、返回结果
│   └── api
│       └── v1
├── docs ## swag 文档
├── internal ##内部逻辑，业务目录
│   ├── cache　## 缓存
│   ├── code　## 错误码设计
│   ├── config　## 配置
│   ├── proto　## grpc proto
│   ├── sentinelm ## sentinel 限流
│   └── version　## 版本
├── models　## 数据库交互
├── pkg　## 一些封装好的 package
│   ├── token
│   └── utils
├── router ## 路由及中间件目录
│   ├── api
│   └── middleware
├── rpc ## 业务逻辑层
├── service ## 业务逻辑层
└── tool ## 小工具
    main.go # 项目入口文件
```

#### 交互流程

`简要的交互可用文字说明，复杂的交互建议使用流程图，交互图或其他图形进行说明`

##### 数据流向

![示意图](tool/data_flow.png)

用户的请求先到　controller，进行参数校验。然后到service层进行业务逻辑的处理，如果需要取数据，则由service向models获取数据。最终将结果返回给用户。

#### ✨ 技术栈

* 框架路由使用 [Gin][3] 路由
* 中间件使用 [Gin][4] 框架的中间件
* 数据库组件 [GORM][5]
* 文档使用 [Swagger][6] 生成
* 配置文件解析库 [Viper][7]
* 使用 [JWT][8] 进行身份鉴权认证
* 校验器使用 [validator][9]  也是 Gin 框架默认的校验器
* 包管理工具 [Go Modules][10]: 视频：https://www.bilibili.com/video/av63052644/
* 使用 make 来管理 Go 工程
* 使用 JSON 文件进行多环境配置

#### 📖 开发规范

* [Uber Go 语言编码规范][1]
* [错误码设计][2]
* [DB设计规范][11]

### 模块简介

`架构图中如果有很多模块，需要对各个模块的功能进行简要介绍`

#### 用户模块

* 注册
* 查询信息
* 更新
* 删除

### 数据库设计

```sql
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` bigint unsigned DEFAULT NULL,
  `updated_at` bigint unsigned DEFAULT NULL,
  `deleted_at` bigint unsigned DEFAULT NULL,
  `is_delete` tinyint(1) DEFAULT NULL,
  `user_name` varchar(30) NOT NULL,
  `password_digest` longtext,
  `nickname` varchar(30) NOT NULL,
  `status` longtext,
  `avatar` varchar(1000) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_name` (`user_name`),
  UNIQUE KEY `nickname` (`nickname`),
  KEY `user_name_idx` (`user_name`),
  KEY `nickname_idx` (`nickname`)
) ENGINE=InnoDB AUTO_INCREMENT=1000011 DEFAULT CHARSET=utf8mb4 
```

### 接口细节

`输入什么参数，输出什么参数，根据接口前端、后端、APP、QA就能够并行做编码实现了`

#### 📝 接口文档

`http://127.0.0.1:8000/swagger/index.html`

## 性能测试

`ab -n 1000 -c 100 'http://127.0.0.1:8000/version'`

### 设计与折衷

`设计与折衷是总体设计中最重要的部分`

暂无

### 潜在风险

暂无

## 运维

### 🚀 部署

#### 💻 常用命令

##### make

```sh
make - compile the source code
make clean - remove binary file and vim swp files
make ca - generate ca files
make docs - gen swag doc
make test - go test
make build - go build
```

#### 部署

##### 编译并生成二进制文件

```bash
make build
```

* ./ffly-plus 运行
* 你也可以选择用`supervisor`进行部署

### 监控

* `http://127.0.0.1:8001/debug/statsviz/`
* `http://127.0.0.1:8000/`
* `http://127.0.0.1:8000/debug/pprof/`

#### 版本查看

* `http://127.0.0.1:8000/version`
* `./ffly-plus -v`

## Features

* Graceful restart or stop (fvbock/endless)
* Cron
* Redis

## 其他 pre-commit install

* pip install [pre-commit][12]
* [pre-commit-golang][13]
* pre-commit install
* go 的一些代码检查工具
  * go get github.com/fzipp/gocyclo/cmd/gocyclo
  * go get -u github.com/tsenart/deadcode
  * go get -u github.com/alecthomas/gometalinter
  * go get -v -u github.com/go-critic/go-critic/cmd/gocritic
  * go get honnef.co/go/tools/cmd/staticcheck


## 项目参考

* https://github.com/Away0x/gin_weibo
* https://github.com/1024casts/snake
* https://github.com/eddycjy/go-gin-example
* https://github.com/Gourouting/singo
* https://github.com/Gourouting/giligili

[1]: https://github.com/xxjwxc/uber_go_guide_cn
[2]: https://github.com/colinrs/ffly-plus/tree/master/internal/code
[3]: https://github.com/gin-gonic/gin
[4]: https://github.com/gin-gonic/gin
[5]: https://github.com/jinzhu/gorm
[6]: https://swagger.io/
[7]: https://github.com/spf13/viper
[8]: https://jwt.io/
[9]: https://github.com/go-playground/validator
[10]: https://github.com/golang/go/wiki/Modules
[11]: https://github.com/colinrs/ffly-plus/tree/master/models
[12]: https://pre-commit.com/
[13]: https://github.com/dnephin/pre-commit-golang