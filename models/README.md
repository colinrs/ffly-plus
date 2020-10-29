# models

models 层或者叫 DAO，是数据访问层，负责访问 DB、MC、外部 HTTP 等接口，对上层屏蔽数据访问细节。后续更换、升级ORM引擎，不影响业务逻辑。能提高测试效率，单元测试时，用Mock对象代替实际的数据库存取，可以成倍地提高测试用例运行速度

## 具体职责有

* SQL 拼接和 DB 访问逻辑
* DB 的拆库分表逻辑
* DB 的缓存读写逻辑
* HTTP 接口调用逻辑

**Tips**: 如果是返回的列表，尽量返回map，方便service使用。

## 建议

* 推荐使用编写原生SQL
* 禁止使用连表查询，好处是易扩展，比如分库分表
* 逻辑部分在程序中进行处理

一个业务一个目录，每一个repo go文件对应一个表操作，比如用户是在user目录下，涉及用户相关的都可以放到这里，根据不同的模块分离到不同的文件，同时又避免了单个文件func太多的问题。比如： 用户基础服务- user.go

## 单元测试

关于数据库的单元测试可以用到的几个库：

* go-sqlmock https://github.com/DATA-DOG/go-sqlmock 主要用来和数据库的交互操作:增删改
* GoMock https://github.com/golang/mock

## Reference

* https://github.com/realsangil/apimonitor/blob/fe1e9ef75dfbf021822d57ee242089167582934a/pkg/rsdb/repository.go
* https://youtu.be/twcDf_Y2gXY?t=636
* Unit testing GORM with go-sqlmock in Go
* 如何使用Sqlmock对GORM应用进行单元测试

## DB设计规范

* ENGINE=InnoDB
* DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
* 禁止使用存储过程，视图，触发器，Event
* 禁止使用mysql关键字和保留字
* 字段名称见名知意
* 布尔类型字段，is_做为前缀。比如：is_deleted
* 必须添加注释
* 必须包含三个字段
  * 主键：`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT
  * 两个日期字段：create_time，update_time，对表的记录进行更新的时候，必须更新update_time
* 尽量控制单表数据量的大小在500万以内
* 禁止使用外键，如果要保证完整性，应由应用层实现
* 禁止在数据库中存储图片，文件等大的二进制数据。通常存储于文件服务器，数据库只存储文件地址信息
* 新增列，禁止指定位置 （FIRST / AFTER）
* 所有字段均定义为NOT NULL
* 所有字段都要有注释，禁止使用中文注释
* 禁止使用ENUM，用TINYINT代替
* 使用INT UNSIGNED （UNIX timestamp）存储时间
* 使用INT UNSIGNED存储IPV4地址
* 选择合适的数据类型，使用可存下数据的最小的数据类型
* 根据业务区分使用 TINYINT/SMALLINT/INT/BIGINT，分别会占用1/2/4/8字节
* 根据业务区分使用 CHAR/VARCHAR
* 字段长度固定，或者长度近似的业务场景，使用CHAR
* 字段长度相差较大，或者更新较少的业务场景，使用VARCHAR，能够减少空间
* VARCHAR长度设置为满足需求的最小长度（字符数）
* 尽可能不使用TEXT、BLOB
* 数值类型尽量使用严格数值数据类型 （比如INT, DECIMAL），而不是近似数值数据类型 （比如FLOAT, DOUBLE）
* 非负数值类型设置为UNSIGNED
* 禁止存储明文密码
* 建议单表索引数量不超过5个
* 建议组合索引字段数不超过5个
* 理解组合索引最左前缀原则，避免重复索引，如果建立了(a,b,c)，相当于建立了(a), (a,b), (a,b,c)
* 对于频繁的查询优先考虑使用覆盖索引
* 非必要不要进行JOIN查询，如果要进行JOIN查询，被JOIN的字段必须类型相同，并建立索引
* 禁止使用SELECT *，必须使用SELECT <字段列表> 查询
* LIKE查询时，禁止用%通配符最左前导
* 能用UNION ALL不要用UNION
* 最好不要使用SQL运算
* 最好不要在WHERE使用函数
* 关键业务SQL上线前，必须使用 EXPLAIN 分析查询语句，避免慢查询
* 超1000行的批量写（UPDATE、DELETE、INSERT）操作，要分批多次进行操作，每批次LIMIT 1000，两个批次之间引入延时（比如1s）
* 使用命令行连接MySQL时，使用-p/--password，禁止显示输入密码
* 禁止直接执行SQL更新生产数据库的数据，提单让DBA操作
* 禁止在线上做数据库压力测试、
* 禁止从测试、开发环境直连生产数据库

参考[阿里数据库设计规范][1]

[1]: https://developer.aliyun.com/article/709387
