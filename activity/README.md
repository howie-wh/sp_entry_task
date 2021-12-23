## 站点简介
* 基于Gin的后台管理系统，聚会活动平台，此站点为用户端后台管理
* 后端采用GO语言 框架 Gin Xorm。

## 内置功能

1.  用户管理：用户注册/查询等。
2.  活动管理：聚会活动的查询/报名等操作。
3.  评论管理：评论查询，登陆用户对活动评论等操作。

## github地址
[https://git.garena.com/wuhui/entry_task](https://git.garena.com/wuhui/entry_task)

## 配置
项目数据库文件 db.sql 
创建数据库导入后修改配置/config/config-prod.ini

##运行
go run main.go 
网站地址：http://localhost:8082

## 依赖
> Gin框架 [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
>
> gotool[https://github.com/druidcaesa/gotool](https://github.com/druidcaesa/gotool)
>
>jwt [https://github.com/dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)
>
>xorm [https://github.com/go-xorm/xorm](https://github.com/go-xorm/xorm)

