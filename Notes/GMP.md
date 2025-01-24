## GMP

#### 拓扑图

                    +------------------------+
                ->  |  global goroutine List |  ->
                    +------------------------+

                                   +---+    +---+    +---+
                                   | . |    |   |    | G |
                                   | . |    | . |    |---|
                    +--            | . |    | . |    | G |
                    |   goroutine  |---|    | . |    |---|           P的local队列
                    |              | G |    |   |    | G |
                    |              |-—-|    |-—-|    |-—-|
                    |              | G |    | G |    | G |
        gorountine  |              +-—-+    +-—-+    +-—-+
          调度器     |                |        |        |
                    |   逻辑的        |        |        |
                    |   processor    P        P        P            GOMAXPROCS个
                    |
                +-- |-- 物理上(即线程)        M      M      M
                |   |  User-level Thread
                |   +--
                |
                +------------------------------------------
                |
操作系统调度器    |       Kernel Thread
                |
                +------------------------------------------
                |                  +-----+  +-----+
                |    物理多核CPU    | CPU1|  | CPU2|   ...
                |                  +-----+  +-----+
                +------------------------------------------


#### 协程调度的优先级与顺序
1. 本地goroutine队列
    - 优先运行最近加入的goroutine
    - 然后运行较早加入的goroutine
2. 全局goroutine队列
    - 当P执行61次循环时或者本地goroutine队列为空
    - 包含从其他P窃取或者新创建的goroutine
3. 网络轮询器
    - 当goroutine因io阻塞时使用
    - 处理网络相关的goroutine
4. 系统调用
    - 当goroutine执行系统调用
    - 可能导致 M 释放 P
5. 工作窃取
    - 空闲的 M 从其他 P 的队列窃取Goroutine
    - 平衡负载和避免空闲

#### 寻找可执行 G 过程
findRunnable()
    1. runqget() 
        - 作用: 从本地队列中寻找可运行的 G
        - 从队头获取g, 并通过原子操作更新runqhead
        - inheritTime 为 true 则新 G 会继承时间片,减少上下文切换开销
    2. globrunqget()
        - 作用: 从全局队列中获取可运行的 G
        - 最多为 P 队列的一半, 即 256/2 = 128 个
        - runqput() 将一个个 G 放入本地队列中, 从队尾加入
    3. netpoll()
        - 作用: 获取网络轮询器中可运行的网络协程
        - 通过kevent监听事件, 并处理用于唤醒轮询的事件 netpollBreakRd
    4. stealWork()
        - 作用: 从其他 P 窃取可运行的 G 或者 timer
        - 尝试4次, 最后一次会尝试获取timer, 也会尝试获取其他 P 的runnext
        - 获取其他 P 队列中 G 的一半
        - 从队头窃取, 使用原子操作runqhead

#### 协程切换时机
1. 基于协作的抢占式调度
    - 主动挂起 runtime.gopack()
    - 系统调用结束 exitsyscall()
    - 函数跳转 morestack()
2. 基于信号的抢占式调度
    - 信号调度 doSigPreempt()


#### 
P<sub>0</sub>      
>>>
