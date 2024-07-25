# <Unix网络编程>

#### 写在前面
> 本书起始于读书任务，目的是强化网络编程基础，故本次使用选择阅读，有些章节工作相关性较低，选择跳过

## 第一章 简介
- 

## 第十五章 Unix域协议
> 1. 概述
> 2. Unix域套接字地址结构
> 3. socketpair函数
> 4. 套接字函数
> 5. Unix域字节流客户/服务器程序
> 6. Unix域数据报客户/服务器程序
> 7. 描述符传递
> 8. 接收发送者的凭证
> 9. 小结

- 定义
  - 不是一个实际的协议族，而是一种单个主机的通信方法。
  - 第二章介绍的进程间通信(IPC Inter-Progress Communication) ,实际上就是单个主机上的客户/服务器通信,Unix域协议也可以认为是IPC方法之一。本地IPC
- 提供两类套接字
  - 字节流套接字(类似TCP)
  - 数据报套接字(类似UDP)
- 为什么使用Unix域套接字
  1. 快 比TCP套接字快一倍
  2. 可用于同一主机不同进程间的传递描述符 (后面有例子)
  3. Unix域套接字把客户的凭证(用户id和组id)提供给服务器,从而可以提供额外的安全检查 (后面有例子)
- 协议地址
  - 普通文件系统中的路径名
  - ipv4协议地址由32位的地址和16位的端口组成
  - ipv6协议地址由128位的地址和16位端口号组成
  - 路径名不是普通的Unix文件,除非与Unix套接字关联,否则无法读写

- bind()
  - 

- socketpair //创建两个随后连接起来的套接字
  - socketpair(family int, type int, protocol int, sockfd [2]int)
  - 指定type为SOCK_STREAM 得到两个流管道,与调用pipe创建的普通unix管道类型,区别是全双工的

- 套接字函数
  1. 由bind创建的路径名默认访问权限是0777,并按照当前umask值修正
  2. Unix域套接字关联的路径名应该是一个绝对路径名。否则客户就必须跟服务器在同一个工作目录中。
  3. connect调用中指定的路径名必须是一个当前绑定在某个***打开***的Unix域套接字上,并且他们的套接字类型(***字节流或数据报***)也必须一致。所以有可能有以下出错的情况:
      - 路径名已存在却不是一个套接字
      - 路径名已存在也是一个套接字,不过没有与之关联的打开的描述符
      - 路径名已存在也是一个套接字,不过类型不相符合
  4. 调用connect连接一个Unix域套接字涉及的权限等同于调用open以只写方式访问相应的路径名
  5. Unix域字节流套接字类似TCP套接字: 都是为进程提供一个无记录边界的字节流接口
      - 无记录边界：上层传下来的是以bit流的形式传下来的,比如限定在一个固定数值的bits,到这个固定长度断一下,这就是无边界的。如果上层传下来的是一个完整的包的形式,比如有像固定的包头,CRC检验码,长度标志位等等这些信息的,打成一个包的形式发给下层的,这样的就是有边界的了。
  6. 对某个Unix域字节流套接字的connect调用发现监听这个套接字的队列已满,则立即返回一个ECONNREFUSED错误。这一点跟TCP套接字不同,TCP监听端会忽略新到达的SYN,TCP发起段会数次发送SYN重试
  7. Unix域数据报套接字类似于UDP套接字: 它们都提供一个保留记录边界的不可靠的数据报服务。
  8. 在未绑定的Unix域套接字发送数据报不会自动给这个套接字捆绑一个路径名,这一点不同于UDP套接字: 在未绑定的UDP套接字上发送数据报导致给这个套接字捆绑一个临时端口。这一点意味着除非数据报发送端已经捆绑了一个路径名到套接字,否则数据报接收端无法发回应答数据报。类似的,对于某个Unix域数据报套接字的connect调用不会给本套接字捆绑一个路径名,这一点不同于TCP和UDP


- 描述符传递
  - 之前的方法: 可以让父进程把描述符传递给子进程
    1. fork调用之后,子进程共享父进程所有打开的描述符
    2. exec调用执行之后,所有描述符通常保持打开状态不变
  - 第一个例子中,进程先打开一个描述符,再调用fork,然后父进程关闭这个描述符,子进程则处理这个描述符。这样一个打开的描述符就从父进程传递到子进程
  - 当前的UNIX系系统提供了用于从一个进程到任一其他进程传递任一打开的描述符的方法。也就是说，这两个进程之间无需存在亲缘关系。这种技术要求首先在这两个进程之间创建一个UNIX域套接口，然后使用sendmsg跨这个UNIX域套接口发送一个特殊消息。这个消息由内核处理，从而把打开的描述符从发送进程传递到接收进程。使用UNIX域套接口的描述符传递方法是最便于移植的编程技术。
  - SVR4内核使用另一种技术来传递打开的描述符 APUE(15章)讲解I_SENDFD和I_RECVFD两个ioctl命令。BSD技术允许单个sendmsg传递多个描述符,而SVR4技术的一次只能传递单个描述符

- 两个进程之间的传递描述符涉及的步骤
  1. 创建一个字节流或者数据报的Unix域套接字
    - 如果是fork方式,则让子进程打开待传递的描述子,再传递回父进程，那么父进程可以预先调用socketpair创建一个可用于父子进程之间交换描述符的流管道
    - 如果进程之间没有关系，那么服务器进程必须创建一个UNIX域字节流套接口，bind一个路径到该套接口，以允许客户进程connect到该套接口。客户然后可以向服务器发送一个打开某个描述符的请求，服务器再把该描述符通过UNIX域套接口传递回客户。客户和服务器之间也可以使用UNIX域数据报套接口，不过这么做缺乏优势，而且数据报存在被丢弃的可能性。
  2. 发送进程通过调用返回描述符的任一UNIX函数打开一个描述符，这些函数的例子有：open、pipe、mkfifo、socket和accept。可以在进程之间传递的描述符不限类型，这就是我们称这种技术为“描述符传递”而不是“文件描述符传递”的原因。比如(参数描述符、数据库中的行描述符)
  3. 发送进程创建一个msghdr结构(第十四章)，其中含有待传递的描述符。POSIX规定描述符作为辅助数据(msghdr结构的msg_control成员)发送。发送进程调用sendmsg跨来自步骤(1)的UNIX域套接口发送该描述符。至此我们说这个描述符“在飞行中（in flight）”。即使发送进程在调用sendmsg之后但在接收进程调用recvmsg之前就关闭了该描述符，对于接收进程它仍然保持打开状态。发送一个描述符导致该描述符的引用计数加1。
  4. 接收进程调用recvmsg在来自步骤（1）的UNIX域套接口上接收这个描述符。这个描述符在接收进程中的描述符号不同于它在发送进程中的描述子号是正常的。传递一个描述符并不是传递一个描述符号，而是涉及在接收进程中创建一个新的描述符，而这个描述符指引的内核中文件表项和发送进程中飞行前的那个描述符指引的相同。

  客户和服务器之间必须存在某种应用协议，以便描述符的接收进程预先知道何时期待接收。如果接收进程调用recvmsg时没有分配用于接手描述符的空间，而且之前已有一个描述符被传递并正等着被读取，这个早先传递的描述符就会被关闭。另外，在期待接收描述符的recvmsg调用中应该避免使用MSG_PEEK标志，否则后果不可预料。
    - MSG_PEEK标志会将套接字接收队列中的可读的数据拷贝到缓冲区，但不会使套接字接收队列中的数据减少，常见的是：例如调用recv或read后，导致套接字接收队列中的数据被读取后而减少，而指定了MSG_PEEK标志，可通过返回值获得可读数据长度，并且不会减少套接字接收缓冲区中的数据，所以可以供程序的其他部分继续读取。