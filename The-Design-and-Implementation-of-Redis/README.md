## Redis设计与实现 (based on redis v2.9)

#### 章节编排
1. 数据结构与对象
    - 数据库里面的每个键值对都是有对象组成
        - 键是字符串对象
        - 值可以是 字符串对象、列表对象、哈希对象、集合对象、有序集合对象
2. 单机数据库的实现
    - 数据库 对实现原理进行介绍，说明保存键值对的方法，保存键值对的过期时间的方法，自动删除过期键值对的方法
    - RDB持久化
    - AOF持久化
    - 事件
        - 文件事件 应答客户端的连接请求，接收客户端发送的命令请求，以及向客户端返回命令回复
        - 时间事件 维护和管理保持redis服务器正常运作
    - 客户端 对Redis服务器维护和管理客户端状态
    - 服务器 对单机Redis服务器的运作机制介绍
3. 多机数据库的实现
    - 复制 主从复制的实现原理
    - 哨兵 对`Redis Sentinel`对原理进行介绍
    - 集群 对Redis集群实现原理介绍
4. 独立功能的实现
    - 发布于订阅
    - 事务 `MULTI`、`EXEC`、`WATCH`等命令的实现进行介绍 解释事务如何实现
    - Lua脚本
    - 排序 对SORT命令及其可选项的实现原理进行介绍
    - 二进制位数组
    - 慢查询日志
    - 监视器 


#### 数据类型
1. 简单动态字符串





2. 压缩列表(ziplist)
    1. 类似数组，通过一片连续的内存空间来存储数据，但是不同的是存储的数据大小不一定相同
    2. 
    