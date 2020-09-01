### go-api
A Restful API Services Based on Golang/Gin

### todo
- [x] Service 初步框架
- [x] 集成gorm并实现基础的对象查询
- [x] 日志处理
- [ ] 实现多表JOIN、GROUPBY、HAVING
- [ ] PB协议
- [ ] redis/MongoDB/kfaka/rabbitMQ 基础接入

### Target
- 提供支持http/grpc的服务
- 支持Docker
- 基于日志/Redis的实时统计
- 流量控制、分发

### Packges
- [Gin： github.com/gin-gonic/gin](#https://github.com/gin-gonic/gin)
- [MySQL Driver： github.com/go-sql-driver/mysql](#https://github.com/go-sql-driver/mysql)
- [gorm： github.com/jinzhu/gorm](#https://github.com/jinzhu/gorm)
- [validator： github.com/go-playground/validator](#https://github.com/go-playground/validator)
- [logrus： github.com/sirupsen/logrus](#https://github.com/sirupsen/logrus)

### Docs
- [gorm：主要复杂查询、事务和锁](#https://gorm.io/zh_CN/docs/query.html)

### Extra 不实现
- 前台： Python 或者 PHP渲染
- 后台： Vue
