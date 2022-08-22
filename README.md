<!--
 * @Author: JF-011101 2838264218@qq.com
 * @Date: 2022-08-20 22:34:15
 * @LastEditors: JF-011101 2838264218@qq.com
 * @LastEditTime: 2022-08-21 11:58:46
 * @FilePath: \dytt\README.md
 * @Description: functions and args
-->
# dytt
dytt(douyin-tiktok)是基于 gRPC微服务 + Gin HTTP服务完成的抖音后端项目

## 一、项目特点

1. 采用RPC框架（Grpc）脚手架生成代码进行开发，基于 **RPC 微服务** + **Gin 提供 HTTP 服务**

2. 基于《[接口文档在线分享](https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/)[- Apifox](https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/)》提供的接口进行开发，使用Insomnia进行API调试并生成[测试文档](https://github.com/JF-011101/DYTT/blob/master/Insomnia_2022-08-22)。使用《[极简抖音App使用说明 - 青训营版](https://bytedance.feishu.cn/docs/doccnM9KkBAdyDhg8qaeGlIz7S7) 》提供的APK进行Demo测试， **功能完整实现** ，前端接口匹配良好。

3. 代码结构采用 (HTTP API 层 + RPC Service 层+Dal 层) 项目 **结构清晰** ，代码 **符合规范**

4. 使用 **JWT** 进行用户token的校验

5. 使用 **ETCD** 进行服务发现和服务注册；

6. 使用 **Minio** 实现视频文件和图片的对象存储

7. 使用 **Gorm** 对 MySQL 进行 ORM 操作；

8. 使用 **OpenTelemetry** 实现链路跟踪；

9. 数据库表建立了索引和外键约束，对于具有关联性的操作一旦出错立刻回滚，保证数据一致性和安全性



## 二、项目说明

## 1. 服务调用关系

![](https://raw.githubusercontent.com/JF-011101/Image_hosting_rep/main/arc1.png)
## 2. 数据库 ER 图

![](https://raw.githubusercontent.com/JF-011101/Image_hosting_rep/main/dytter1.jpg)
## 3. 代码介绍

### 3.1 代码目录结构介绍

| 目录 | 子目录 | 说明 | 备注 |
| --- | --- | --- | --- |
| [cmd](https://github.com/jf-011101/dytt/tree/master/cmd) | [api](https://github.com/jf-011101/dytt/tree/master/cmd/api) | api 服务的 **业务代码** | 包含 [Gin](https://github.com/jf-011101/dytt/blob/master/cmd/api/main.go)和 [RPC_client](https://github.com/jf-011101/dytt/tree/master/cmd/api/rpc) |
|| [comment](https://github.com/jf-011101/dytt/tree/master/cmd/comment) | command 服务的业务代码 |
|| [favorite](https://github.com/jf-011101/dytt/tree/master/cmd/favorite) | favorite 服务的业务代码 |
|| [feed](https://github.com/jf-011101/dytt/tree/master/cmd/feed) | feed 服务的业务代码 |
|| [publish](https://github.com/jf-011101/dytt/tree/master/cmd/publish) | publish 服务的业务代码 |
|| [relation](https://github.com/jf-011101/dytt/tree/master/cmd/publish) | relation 服务的业务代码 |
|| [user](https://github.com/jf-011101/dytt/tree/master/cmd/user) | user 服务的业务代码 |
| [config](https://github.com/jf-011101/dytt/tree/master/config) | 微服务及 pkg 的 **配置文件** |
| [dal](https://github.com/jf-011101/dytt/tree/master/dal) | [db](https://github.com/jf-011101/dytt/tree/master/dal/db) | 包含 [Gorm 初始化](https://github.com/jf-011101/dytt/blob/master/dal/db/init.go) 、[Gorm 结构体及 数据库操作逻辑](https://github.com/jf-011101/dytt/blob/master/dal/db/user.go) |
|| [pack](https://github.com/jf-011101/dytt/tree/master/dal/pack) | 将 [Gorm 结构体](https://github.com/jf-011101/dytt/blob/master/dal/pack/user.go#L25) 封装为 [protobuf 结构体](https://github.com/jf-011101/dytt/blob/master/Grpc_gen/user/user.pb.go#L268)的 **业务逻辑** | Protobuf 结构体由 Grpc自动生成 |
| [idl](https://github.com/jf-011101/dytt/tree/master/idl) | proto **接口定义文件** |
| [Grpc_gen](https://github.com/jf-011101/dytt/tree/master/Grpc_gen) | Grpc **自动生成的代码** |
| [pkg](https://github.com/jf-011101/dytt/tree/master/pkg) | [ilog](https://github.com/jf-011101/dytt/tree/master/pkg/ilog) | 简单的自定义 **Logger** 及其接口 |
|| [errno](https://github.com/jf-011101/dytt/tree/master/pkg/errno) | **错误码**| 错误码设计逻辑:[jf-011101/ErrnoCod](https://github.com/jf-011101/ErrnoCode) |
|| [jwt](https://github.com/jf-011101/dytt/tree/master/pkg/jwt) | 基于 [golang-jwt](http://github.com/golang-jwt/jwt)的代码封装 |
|| [minio](https://github.com/jf-011101/dytt/tree/master/pkg/minio) | **Minio** 对象存储初始化及代码封装 |
|| [ttviper](https://github.com/jf-011101/dytt/tree/master/pkg/ttviper) | **Viper** 配置存取初始化及代码封装 |

