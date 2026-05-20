
#### 数组
- 扩容
    - 小于256是翻倍扩容
    - 大于256是随着规模逐渐从翻倍扩容衰减到1.25倍扩容
    - 如果翻倍扩容也不满足需要的容量，则直接使用需要的容量作为扩容后的空间

#### Map
- 底层
    - 哈希表(hashmap)
    - ```
        // A header for a Go map.
        type hmap struct {
            count     int // 当前哈希表中键值对的数量
            flags     uint8
            B         uint8  // 当前哈希表持有的 buckets 数量, 因为哈希表中桶的数量都按2倍扩容,改字段存储对数，也就是 len(buckets) == 2^B
            noverflow uint16 // 溢出桶的大致数量
            hash0     uint32 // hash seed

            buckets    unsafe.Pointer // 存储 2^B 个桶的数组
            oldbuckets unsafe.Pointer // 扩容时用于保存之前 buckets 的字段 , 大小事buckets的一般
            nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

            extra *mapextra // optional fields
        }

        // mapextra 主要维护，当hmap中的buckets中有一些桶发生溢出，但有达不到扩容阈值时，存储溢出的桶
        type mapextra struct {
            overflow    *[]*bmap
            oldoverflow *[]*bmap

            // nextOverflow holds a pointer to a free overflow bucket.
            nextOverflow *bmap
        }
      ```
    - 关键字段
        - count：表示哈希表中键值对的数量；
        - B：这是以 2 为底的对数，用于确定桶（bucket）的数量。例如，当 B = 1 时，桶的数量为 2^1 = 2 个；当 B = 2 时，桶的数量为 2^2 = 4 个，以此类推；
        - hash0：是计算键的哈希值时用到的一个因子；
        - buckets：是一个指向桶数组的指针，每个桶用于存储键值对；
        - overflow：当桶中装不下更多元素，且 key 又被 hash 到这个桶时，就会放到溢出桶，原桶的 overflow 字段指向溢出桶（链地址法）。
        - bmap: 每个bmap存储8个key-value对
- 扩容
    - 增量扩容
        - 负载超过阈值会发生双倍扩容。阈值计算方式是元素个数(count)除以桶数量(2^B)如果大于等于6.5
        - 扩容过程
            - B的值加1，从而使桶数量翻倍
            - 数据迁移时，会将原来桶中数据重新分配到新的桶中。具体是根据键的哈希值重新计算在新桶中的位置。
    - 等量扩容
        - 当有大量的键被删除，溢出桶过多，链表过长，key-value数据饱和度很低，可能会触发等量扩容
        - 创建和旧桶数量一样多的新桶，然后把原来的键值对迁移到新桶中
    - 渐进扩容策略，桶被实际操作(写、删)到时，由使用者完成数据迁移，避免因为一次性迁移全量数据引发性能抖动。
- 遍历
    - 无序的
        - 取一个随机数，决定遍历起点
            - 起始桶节点不固定，桶的8个kv对的随机位置起始，每个bmap的8个kv对遍历完成后再遍历下一个bmap
            - 桶可能处于扩容迁移状态中，所以要在逻辑意义上把老桶还没迁移的数据模拟迁移到新桶后，再遍历，可能原先再老桶中的数据在新桶中位置不一样


#### channel
- 底层结构
    - 环形数组+双指针
    ```
    type hchan struct {
        qcount   uint           // 元素数量
        dataqsiz uint           // 容量
        buf      unsafe.Pointer // 环形数组指针
        elemsize uint16         // 元素类型大小
        closed   uint32         // 关闭标记符
        elemtype *_type         // 元素类型
        sendx    uint           // 写入指针
        recvx    uint           // 读指针
        recvq    waitq          // 因为读而陷入阻塞的协程队列
        sendq    waitq          // 因为写而陷入阻塞的协程队列
        lock     mutex
    }
    type waitq struct {
        first *sudog
        last  *sudog
    }
    type sudog struct {
        g *g

        next *sudog
        prev *sudog
        elem unsafe.Pointer // data element (may point to stack)

        success bool

        c        *hchan // channel
    }
    ```

#### 

#### 分布式锁
- 


#### defer
- 延迟调用
- 多个defer，后进先出
- 常见应用场景
    - 资源释放
    - 异常捕获和处理

#### panic

#### 面向对象
- 封装继承多态
    - 封装
        - 将数据和操作数据的方法绑定在一起，对外部隐藏实现细节
    - 继承
        - 继承属性和方法，可以重写父类方法以满足特定的需求
    - 多态
        - 不同对象对同一消息或方法调用产生不同相应或者行为
- 没有类class，结构体类似类
- 方法和函数
    - 方法调用者不严格要求指针类型或者值类型



#### 知识学习
- go.dev
- http server
    - hello,world
    - crud演化复杂逻辑
        - 接入认证系统
        - 受保护路由
        - 实现自己的中间件
        - context注入中间件 
- cli tools 例如cobra、bubble tea
    - 如何在cli里传不同flag
    - 如何把状态从本地代入云端aws
- grpc 
    - http后端有瓶颈 转为微服务架构
    - 微服务之间互相调用
- pipeline jobs/scripts
- 优秀资源
    - lets-go-further
    - learn go with test
    - writing an interpreter in go 测试驱动开发
    - 100 go mistamistakes and how to avoid them