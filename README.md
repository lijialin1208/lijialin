
一、项目介绍
极简抖音互动方向与社交方向实现

//服务器有时候视频会比较慢
http://101.42.28.6:8080/(服务器已过期)

https://github.com/lijialin1208/lijialin
二、项目分工
| 团队成员 | 主要贡献 |
| --- | --- |
| 李嘉林 | 基本功能，互动方向拓展功能，社交方向拓展功能实现 |
| 杜昊杰 | 文件存储服务器搭建 |

三、项目实现
3.1 技术选型与相关开发文档
http框架：Hertz（开始使用的是gin框架，但是Hertz的学习成本低，且相比于gin更有优势，）
RPC框架：grpc
数据库：MYSQL，gorm
工具库：Viper、JWT、ffmpeg
由于本地测试ip地址经常发生变化，使用viper读取配置文件
使用JWT生成token，解析token
由于需要对视频进行抽帧获取封面，使用ffmpeg工具进行抽帧
服务之间的调用使用grpc框架（kitex框架不太熟）
3.2 架构设计
文件服务负责对外提供视频、图片等文件资源，数据库中存储文件的路径
![1280X1280](https://github.com/lijialin1208/lijialin/assets/87974640/bde7cf04-f766-4605-8b2d-b5e24acb4eef)

3.3 项目代码介绍
目录介绍
以douyin-user为例
config——配置文件
dal——数据库连接和初始化操作
pojo——实体类
pb——使用protoc生成的服务接口
server——具体服务类
tool——工具包，包括token的生成、解析，视频抽帧等等
main——初始化数据库，读取配置文件，注册服务
// 在douyin-api的  middleware——中间件
![eb6c242e-5cdc-49a6-8b62-902d272a3d83](https://github.com/lijialin1208/lijialin/assets/87974640/2f030002-b2db-4731-b869-ab9c25b9cc89)



四、测试结果
接口文档地址：
https://www.apifox.cn/apidoc/shared-53a457fa-2d41-4109-aa74-43af3e570b51

五、Demo 演示视频 （必填）


https://github.com/lijialin1208/lijialin/assets/87974640/d0fd2a08-fc5a-4b2f-ae36-70eb71f31a4f


六、项目总结与反思
1. 目前仍存在的问题
  视频响应比较慢有待优化
2. 已识别出的优化项
  构建联合索引
3. 架构演进的可能性
4. 项目过程中的反思与总结
部署方面：认识到了docker的重要性
编码方面：使用配置文件，不在for循环里面进行数据库操作
优化方面：虽然说基本功能可以实现，但我们也意识到了性能优化的重要的，如何找到性能瓶颈并解决，是我们接下来应该去刻意练习的地方

七、其他补充资料（选填）
注意：视频文件大小设置上传限制，不能大于16M
