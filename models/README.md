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
