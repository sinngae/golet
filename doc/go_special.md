Specialist of Go
---

# 执行顺序
`import –> const –> var –> init() –>main()`

全局变量也有可能执行一段代码，相当于做了init，在编译期计算，还是运行期？

# 逃逸分析
Go编译器逃逸分析，如果一个变量将会逃逸，则在堆上分配。
提程序员解决了应该放在堆上还是栈上的选择。

逃逸场景：
+ 变量大小不确定
+ 变量类型不确定
+ 变量分配的内存超过用户栈最大值
+ 暴露给外部指针

# slice扩容
如果当前容量小于1024，则判断所需容量是否大于原来容量2倍，如果大于，当前容量加上所需容量；否则当前容量乘2。

如果当前容量大于1024，则每次按照1.25倍速度递增容量，也就是每次加上cap/4 (也即256起步)。


# 协程泄露(Goroutine Leak)
协程泄漏是指协程创建之后没有得到释放。
主要原因有：
+ 缺少接收器，导致发送阻塞
+ 缺少发送器，导致接收阻塞
+ 创建协程的没有回收

# 常见的goroutine操作函数
+ runtime.GOMAXPROCS(num int)可以设置线程数目。该值默认为CPU逻辑核数，如果设的太大，会引起频繁的线程切换，降低性能。
+ runtime.Gosched()，用于让出CPU时间片，让出当前goroutine的执行权限，调度器安排其它等待的任务运行，并在下次某个时候从该位置恢复执行。
+ runtime.Goexit()，调用此函数会立即使当前的goroutine的运行终止（终止协程），而其它的goroutine并不会受此影响。runtime.Goexit在终止当前goroutine前会先执行此goroutine的还未执行的defer语句。 
  请注意千万别在主函数调用runtime.Goexit，因为会引发panic。

# 限制运行时操作系统线程的数量
> The GOMAXPROCS variable limits the number of operating system threads that can execute user-level Go code simultaneously. There is no limit to the number of threads that can be blocked in system calls on behalf of Go code; those do not count against the GOMAXPROCS limit.

从官方文档的解释可以看到，GOMAXPROCS 限制的是同时执行用户态 Go 代码的操作系统线程的数量，但是对于被系统调用阻塞的线程数量是没有限制的(实际为体现为协程？)。GOMAXPROCS 的默认值等于 CPU 的逻辑核数，同一时间，一个核只能绑定一个线程，然后运行被调度的协程。

因此对于 CPU 密集型的任务，若该值过大，例如设置为 CPU 逻辑核数的 2 倍，会增加线程切换的开销，降低性能。
对于 I/O 密集型应用，适当地调大该值，可以提高 I/O 吞吐率。（不一定行，可以试下）

另外对于协程，可以用带缓冲区的channel来控制，下面的例子是协程数为1024的例子


# mutex有几种模式？
mutex有两种模式：normal 和 starvation

+ 正常模式  
所有goroutine按照FIFO的顺序进行锁获取，被唤醒的goroutine和新请求锁的goroutine同时进行锁获取，通常新请求锁的goroutine更容易获取锁(持续占有cpu)，被唤醒的goroutine则不容易获取到锁。公平性：否。
+ 饥饿模式
所有尝试获取锁的goroutine进行等待排队，新请求锁的goroutine不会进行锁获取(禁用自旋)，而是加入队列尾部等待获取锁。公平性：是。

参考链接：Go Mutex 饥饿模式，GO 互斥锁（Mutex）原理

# go defer，多个 defer 的顺序，defer 在什么时机会修改返回值？
作用：defer延迟函数，释放资源，收尾工作；如释放锁，关闭文件，关闭链接；捕获panic;

避坑指南：defer函数紧跟在资源打开后面，否则defer可能得不到执行，导致内存泄露。

多个 defer 调用顺序是 LIFO（后入先出），defer后的操作可以理解为压入栈中

defer，return，return value（函数返回值） 执行顺序：首先return，其次return value，最后defer。defer可以修改函数最终返回值，修改时机：有名返回值或者函数返回指针 