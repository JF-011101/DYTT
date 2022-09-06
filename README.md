<!--
 * @Author: jf-011101 2838264218@qq.com
 * @Date: 2022-08-20 22:34:15
 * @LastEditors: jf-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-21 11:58:46
 * @FilePath: \dytt\README.md
 * @Description: functions and args
-->
# DYTT - 微服务实战
DYTT = **D**ou **Y**in **T**ik **T**ok

DYTT 是一个基于 gRPC 微服务框架 + Gin Web 框架开发的抖音后端项目，基于《[接口文档在线分享](https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/)[- Apifox](https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/)》提供的接口进行开发，使用 Insomnia 进行 API 调试并生成[测试文档](https://github.com/jf-011101/DYTT/blob/master/Insomnia_2022-08-22)。使用《[极简抖音App使用说明 - 青训营版](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7) 》提供的APK进行测试， 功能完整实现

## 一、项目概要
 

1. 采用 (HTTP API 层 + RPC Service 层+Dal 层) 项目结构；

2. 使用 x509 证书对服务间通信进行加密和认证；

3. 使用 [go-grpc-middleware](https://github.com/grpc-ecosystem/go-grpc-middleware)中的日志记录、认证、和恢复；

4. 使用 **JWT** 进行用户token的校验；

5. 使用 **ETCD** 进行服务发现和服务注册；

6. 使用 **Minio** 实现视频文件和图片的对象存储；

7. 使用 **Gorm** 对 MySQL 进行 ORM 操作；

8. 使用 **Zipkin** 实现链路跟踪；

9. 数据库表建立了索引和外键约束，对于具有关联性的操作一旦出错立刻回滚，保证数据一致性和安全性；

10. HTTP 服务和 RPC 服务的优雅停止。



## 二、项目说明

## 1. 服务调用关系

![](https://raw.githubusercontent.com/jf-011101/Image_hosting_rep/main/arc1.png)
## 2. 数据库 ER 图

![](https://raw.githubusercontent.com/jf-011101/Image_hosting_rep/main/dytter1.jpg)
## 3. 代码目录结构介绍

| 目录 | 子目录/文件 | 说明 | 备注 |
|:--|:--|:--|:--|
| cmd | api | api 服务入口 |  |
|| comment | command 服务入口 ||
|| favorite | favorite 服务入口 ||
|| feed | feed 服务入口 ||
|| publish | publish 服务入口 ||
|| relation | relation 服务入口 ||
|| user | user 服务入口 ||
| config | *.yml |微服务及 db 的 配置文件||
||cert|CA证书、私钥等||
||sql|sql文件||
| dal | db | 包含 Gorm 初始化、Gorm 结构体及 数据库操作逻辑 ||
|| pack | 将 Gorm 结构体封装为 protobuf 结构体的业务逻辑 | Protobuf 结构体由 gRPC自动生成 |
| idl | *.proto |接口定义文件||
|internal|pkg|项目外不共享的包||
||其他目录|API服务与各微服务的业务代码||
| grpc_gen | / |gRPC 自动生成的代码||
| pkg | errno                   | 错误码 ||
|docs| *.md | 部分配置文档 |                                |
|scripts| *.sh | 工具安装以及代码生成脚本 ||
|third_party| forked | forked pkg(code、errors)                                                   ||
|.| Makefile | 服务快速构建启动                                             ||
|| README.md | 项目说明 ||
|| docker-compose.yml | 项目运行环境docker构建 ||
## 三、项目启动

1. 配置config
    - 参照 config/* 以及 docs/cergen.md 自行配置

2. 运行环境构建
    - 启动 etcd、minio、mysql、zipkin服务
    ```
    docker-compose up -d
    ```

3. 启动服务
    - 项目根目录下执行 make [serverName].server以启动某个服务(serverName: api/user/comment/favorite/feed/publish/relation)
    - 访问 http://127.0.0.1:9411/zipkin/ 可以观测到追踪的服务链。


