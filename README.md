# ffly-plus

一款适合于快速开发业务的 Go Gin Demo，主要是提供 API, RPC 服务。

## 总体设计文档

### 需求

`一般以产品的语言描述，这一块可以拷贝产品需求文档中的story list部分`

### 名词解释

`非相关领域内的同学需要看到文档需要提前了解的一些概念性质的东西`

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

#### ✨ 技术栈

* 框架路由使用 [Gin][3] 路由
* 中间件使用 [Gin][4] 框架的中间件
* 数据库组件 [GORM][5]
* 文档使用 [Swagger][6] 生成
* 配置文件解析库 [Viper][7]
* 使用 [JWT][8] 进行身份鉴权认证
* 校验器使用 [validator][9]  也是 Gin 框架默认的校验器
* 包管理工具 [Go Modules][10]
* 使用 make 来管理 Go 工程
* 使用 JSON 文件进行多环境配置

#### 开发规范

遵循: [Uber Go 语言编码规范][1]

#### 📖 开发规约

* [错误码设计][2]

#### 架构图

#### 交互流程

`简要的交互可用文字说明，复杂的交互建议使用流程图，交互图或其他图形进行说明`

### 模块简介

`架构图中如果有很多模块，需要对各个模块的功能进行简要介绍`

#### 用户模块

* 注册
* 查询信息
* 更新
* 删除

### 数据库设计

### 接口细节

`输入什么参数，输出什么参数，根据接口前端、后端、APP、QA就能够并行做编码实现了`

#### 📝 接口文档

`http://127.0.0.1:8000/swagger/index.html`

## 性能测试

`ab -n 1000 -c 100 'http://127.0.0.1:8000/version'`

### 设计与折衷

`设计与折衷是总体设计中最重要的部分`

### 潜在风险

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

#### Supervisord 部署

##### 编译并生成二进制文件

```bash
make build
```

##### 环境准备

* `mkdir ~/data/{project,logs} -p`
* `pip install supervisor`

这里日志目录设定为 `/data/log`
如果安装了 Supervisord，可以在配置文件`supervisord.conf`中添加下面内容

```ini
[program:ffly-plus]
# environment=
directory=/data/project/ffly-plus
command=/data/project/ffly-plus/ffly-plus -c /data/project/ffly-plus/config/config.prod.json

autostart=true
autorestart=true
user=root
stdout_logfile=/data/log/ffly.log
startsecs = 2
startretries = 2
stdout_logfile_maxbytes=10MB
stdout_logfile_backups=10
stderr_logfile=/data/log/ffly.log
stderr_logfile_maxbytes=10MB
stderr_logfile_backups=10
```

* 重启服务

```bash
supervisorctl -c supervisord.conf
> restart all
```

### 日志

* `/data/logs/ffly-plus/`

### 监控

* `http://127.0.0.1:8001/debug/statsviz/`
* `http://127.0.0.1:8000/`
* `http://127.0.0.1:8000/debug/pprof/`

#### 版本查看

* `http://127.0.0.1:8000/version`
* `./ffly-plus -v`

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