## 站点简介
* 基于Gin的后台管理站点，运营管理平台，管理活动/活动类型等
* 后端采用GO语言 框架 Gin/Xorm。

## 内置功能

1.  用户管理：用户登陆/查询等。
2.  活动管理：聚会活动的创建/编辑/删除/查询等。
3.  活动类型管理：活动类型的创建/编辑/删除/查询等。

## github地址
[https://git.garena.com/wuhui/entry_task](https://git.garena.com/wuhui/entry_task)

## 配置
项目数据库文件 /data/db.sql 
创建数据库导入后修改配置/config/config-prod.ini


##运行
go run main.go 网站地址http://localhost:8081

账号：admin  密码：123456

## 依赖
> Gin框架 [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
>
> gotool[https://github.com/druidcaesa/gotool](https://github.com/druidcaesa/gotool)
>
>jwt [https://github.com/dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)
>
>xorm [https://github.com/go-xorm/xorm](https://github.com/go-xorm/xorm)
