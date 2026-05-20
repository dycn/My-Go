## 基础算法



1. 冒泡排序 
    - 序列中相邻元素两两对比，把最大或者最小的移动到最后      
    - 时间    o(n^2)   
    - 空间    o(1)
2. 选择排序 
    - 循环遍历序列，选出最大或者最小的，跟最前面的元素交换  
    - 时间    o(n^2)  
    - 空间    o(1)
3. 插入排序 
    - 把序列划分已排序和未排序两部分，每次从未排序队列取一个元素插入到已排序部分
    - 时间    o(n^2)  
    - 空间    o(1)
4. 希尔排序
    - 是插入排序的一种改进版本，也称为缩小增量排序。它通过将待排序的列表按照步长(间隔)先对间隔步长的元素进行插入排序，然后逐步缩小步长长度，最终完成有序
    - 时间    O(n log n) 到 O(n²) 之间
    - 空间    O(1)
    - ***适用于中等规模的数据集，不稳定排序***
    - ```
        func shellSort(arr []int) []int {
                n := len(arr)

                // 初始间隔（gap）取数组长度的一半，然后逐步缩小间隔
                for gap := n / 2; gap > 0; gap /= 2 {
                    // 对各个间隔分组进行插入排序
                    for i := gap; i < n; i++ {
                        temp := arr[i]
                        j := i
                        
                        // 对当前间隔的分组进行插入排序
                        for j >= gap && arr[j-gap] > temp {
                            arr[j] = arr[j-gap]
                            j -= gap
                        }
                        arr[j] = temp
                    }
                }
                return arr
        }
      ```
5. 快速排序 
    - 选一个基准数据把数组分成大于基准和小于基准两部分，然后再对这两个部分进行排序，冒泡排序基础上的递归分治法 
    - 时间 o(nlogn) ***每次将列表分成两半，需要 O(log n) 层递归,每层递归需要 O(n) 的时间来合并子列表***
    - 空间 o(n) 
    - ***适合内存中的大规模数据排序，不稳定排序，小规模数据性能较差，不如插入排序等简单算法***
    - ```
        func quickSort(arr []int) []int {
            return _quickSort(arr, 0, len(arr)-1)
        }

        func _quickSort(arr []int, left, right int) []int {
                if left < right {
                        partitionIndex := partition(arr, left, right)
                        _quickSort(arr, left, partitionIndex-1)
                        _quickSort(arr, partitionIndex+1, right)
                }
                return arr
        }

        func partition(arr []int, left, right int) int {
                pivot := left
                index := pivot + 1

                //input [3, 6, 8, 10, 1, 2, 1] i=1 index=1
                //[3, 1, 8, 10, 6, 2, 1] i=4 index=2
                //[3, 1, 2, 10, 6, 8, 1] i=5 index=3
                //[3, 1, 2, 1, 6, 8, 10] i=6 index=4
                for i := index; i <= right; i++ {
                        if arr[i] < arr[pivot] {
                                swap(arr, i, index)
                                index += 1
                        }
                }
                swap(arr, pivot, index-1)
                return index - 1
        }

        func swap(arr []int, i, j int) {
                arr[i], arr[j] = arr[j], arr[i]
        }
      ```
6. 归并排序 
    - 不断分隔成较小的子序列 直到只有1个元素，然后子序列再两两合并，得到一个有序的序列  
    - 时间   O(nlogn) ***每次将列表分成两半，需要 O(log n) 层递归,每层递归需要 O(n) 的时间来合并子列表***
    - 空间   O(n) 
    - ***适合大规模数据，适合外部排序（如对磁盘文件进行排序），稳定排序算法，对于小规模数据，性能可能不如插入排序等简单算法***
    - ```
        func mergeSort(arr []int) []int {
                length := len(arr)
                if length < 2 {
                        return arr
                }
                middle := length / 2
                left := arr[0:middle]
                right := arr[middle:]
                return merge(mergeSort(left), mergeSort(right))
        }

        func merge(left []int, right []int) []int {
                var result []int
                for len(left) != 0 && len(right) != 0 {
                        if left[0] <= right[0] {
                                result = append(result, left[0])
                                left = left[1:]
                        } else {
                                result = append(result, right[0])
                                right = right[1:]
                        }
                }

                for len(left) != 0 {
                        result = append(result, left[0])
                        left = left[1:]
                }

                for len(right) != 0 {
                        result = append(result, right[0])
                        right = right[1:]
                }

                return result
        }
      ```
7. 二分查找
    1. 二分搜索是一种在有序数组中查找某一特定元素的搜索算法
    2. 搜索过程从数组的中间元素开始，如果中间元素正好是要查找的元素，则搜索过程结束；如果某一特定元素大于或者小于中间元素，则在数组大于或小于中间元素的那一半中查找，而且跟开始一样从中间元素开始比较。如果在某一步骤数组为空，则代表找不到。这种搜索算法每一次比较都使搜索范围缩小一半。
    3. ```
        // 基础二分查找（迭代版本）
        func binarySearch(arr []int, target int) int {
            left, right := 0, len(arr)-1
            
            for left <= right {
                // 防止溢出：left + (right - left) / 2
                mid := left + (right-left)/2
                
                if arr[mid] == target {
                    return mid // 找到目标，返回索引
                } else if arr[mid] < target {
                    left = mid + 1 // 目标在右半部分
                } else {
                    right = mid - 1 // 目标在左半部分
                }
            }
            
            return -1 // 未找到
        }
       ```

8. 堆排序
    - 堆是一种特殊的完全二叉树，从头到尾从大到小排序
    - 对于任意数组(序号从1开始)中的一个节点，它的左子节点的序号是 2*x 右子节点序号是 2*x+1
    - ```
        func heapSort(arr []int) []int {
            arrLen := len(arr) // 数组长度
            
            // 构建最大堆 从最后一个非叶子节点开始
            for i := arrLen/2 - 1; i >= 0; i-- {
                heapify(arr, i, arrLen)
            }

            for i := arrLen - 1; i >= 0; i-- {
                swap(arr, 0, i)
                arrLen -= 1
                heapify(arr, 0, arrLen)
            }
            return arr
        }

        func heapify(arr []int, i, arrLen int) {
            left := 2*i + 1
            right := 2*i + 2
            largest := i
            if left < arrLen && arr[left] > arr[largest] {
                largest = left
            }
            if right < arrLen && arr[right] > arr[largest] {
                largest = right
            }
            if largest != i {
                swap(arr, i, largest)
                heapify(arr, largest, arrLen)
            }
        }

        func swap(arr []int, i, j int) {
            arr[i], arr[j] = arr[j], arr[i]
        }   
      ```


	// 基数排序 根据不同位依次进行排序 最终遍历了每一位得到有序的队列


	// 	/*
	// 		在Go语言中，strings.Builder的性能通常比bytes.Buffer好，主要有以下几个原因：
	// 		1. 零拷贝：strings.Builder在内部使用了可变长度的[]byte切片来存储字符串，
	// 		而bytes.Buffer使用了固定长度的[]byte切片。当进行字符串拼接时，strings.Builder
	// 		可以直接修改切片中的内容，而不需要进行额外的内存分配和拷贝操作，从而避免了不必要的性能开销。
	// 		2.预分配内存：strings.Builder在初始化时会预分配一定大小的内存空间，避免了频繁的内存分配和释放操作。
	// 		这样可以减少内存分配的次数，提高性能。
	// 		3.字符串连接优化：strings.Builder提供了WriteString方法，
	// 		可以直接将字符串追加到内部的[]byte切片中，而不需要进行类型转换和拷贝操作。
	// 		这样可以减少不必要的中间步骤，提高字符串连接的效率。
	// 	*/

	// 	//执行顺序：import –> const –> var –>init()–>main()

	// 	//2 个 nil 可能不相等吗？
	// 	//两个nil只有在类型相同时才相等。

	// 	/*
	// 		非接口的任意类型 T() 都能够调用 *T 的方法吗？反过来呢？
	// 		一个T类型的值可以调用*T类型声明的方法，当且仅当T是可寻址的。

	// 		反之：*T 可以调用T()的方法，因为指针可以解引用。
	// 	*/

	// 	// 为什么有协程泄露(Goroutine Leak)
	// 	// 阻塞、死锁(多个goroutine竞争资源)、创建goroutine未回收

	// 	// new和make的区别？
	// 	// new只用于分配内存，返回一个指向地址的指针。它为每个新类型分配一片内存，初始化为0且返回类型*T的内存地址，它相当于&T{}
	// 	// make只可用于slice,map,channel的初始化,返回的是引用。

	// 	// Go语言普通指针和unsafe.Pointer有什么区别
	// 	// unsafe.Pointer是Go的通用指针类型，可以理解为C语言中的void*，
	// 	// 它绕过了Go的类型类型安全。通指针受GC管理和类型约束，unsafe.Pointer不受类型约束但仍受GC跟踪

	// 	// Go语言Map的遍历为什么要设计成无序的
	// 	// map每次遍历,都会从一个随机值序号的桶,在每个桶中，再从按照之前选定随机槽位开始遍历,所以是无序的
	// 	// map 在扩容后，会发生 key 的搬迁，原来落在同一个 bucket 中的 key，搬迁后，有些 key 就要远走高飞了（bucket 序号加上了 2^B）

	// 	// Go语言Map的扩容时机是怎样的
	// 	/*
	// 		扩容是渐进式（gradual）的。它不会在触发扩容时"stop the world"来一次性把所有数据搬迁到新空间，
	// 		而是只分配新空间，然后在后续的每一次插入、修改或删除操作时，才会顺便搬迁一两个旧桶的数据。
	// 		这种设计将庞大的扩容成本分摊到了多次操作中，极大地减少了服务的瞬间延迟（STW），保证了性能的平滑性。

	// 		如果是触发双倍扩容，会新建一个buckets数组，新的buckets数量大小是原来的2倍，
	// 		然后旧buckets数据搬迁到新的buckets。如果是等量扩容，buckets数量维持不变，
	// 		重新做一遍类似双倍扩容的搬迁动作，把松散的键值对重新排列一次，使得同一个 bucket 中的
	// 		 key 排列地更紧密，这样节省空间，存取效率更高
	// 	*/

	// 	/* select是随机的
	// 	select {
	// 	case data := <-ch1:
	// 	    // 处理ch1的数据
	// 	case ch2 <- value:
	// 	    // 向ch2发送数据
	// 	case <-timeout:
	// 	    // 超时处理
	// 	default:
	// 	    // 所有channel都不可用时执行
	// 	}
	// 	*/

	// 	var SM sync.Map //适用于读多写少
	// 	SM.Store("abc", "abc")

	// 	// Mutex 有几种模式
	// 	/*
	// 		正常模式：这是默认模式，讲究的是性能。新请求锁的goroutine会和等待队列头部的goroutine竞争，新来的goroutine有几次"自旋"的机会，如果在此期间锁被释放，它就可以直接抢到锁。这种方式吞吐量高，但可能会导致队列头部的goroutine等待很久，即"不公平"。
	// 		饥饿模式：当一个goroutine在等待队列中等待超过1毫сан（1ms）后，Mutex就会切换到此模式，讲究的是公平。在此模式下，锁的所有权会直接从解锁的goroutine移交给等待队列的头部，新来的goroutine不会自旋，必须排到队尾。这样可以确保队列中的等待者不会被"饿死"。
	// 		当等待队列为空，或者一个goroutine拿到锁时发现它的等待时间小于1ms，饥饿模式就会结束，切换回正常模式。这两种模式的动态切换，是Go在性能和公平性之间做的精妙平衡。
	// 	*/

	// 	// 讲讲Go语言是如何分配内存的
	// 	/*
	// 		mcache（线程缓存）、mcentral（中央缓存）、mheap（页堆）
	// 		根据对象大小分为三类处理
	// 		微小对象（<16字节）：在mcache的tiny分配器中分配，多个微小对象可以共享一个内存块
	// 		小对象（16字节-32KB）：通过size class机制，预定义了67种大小规格，优先从P的mcache对应的mspan中分配，如果 mcache 没有内存，则从 mcentral 获取，如果 mcentral 也没有，则向 mheap 申请，如果 mheap 也没有，则从操作系统申请内存。
	// 		大对象（>32KB）：直接从mheap分配，跨越多个页面
	// 	*/

	// 	// 知道 golang 的内存逃逸吗？什么情况下会发生内存逃逸？
	// 	/*
	// 		主要逃逸场景：
	// 		返回局部变量指针：函数返回内部变量的地址，变量必须逃逸到堆上
	// 		interface{}类型：传递给interface{}参数的具体类型会逃逸，因为需要运行时类型信息
	// 		闭包引用外部变量：被闭包捕获的变量会逃逸到堆上
	// 		切片/map动态扩容：当容量超出编译期确定范围时会逃逸
	// 		大对象：超过栈大小限制的对象直接分配到堆上
	// 	*/

	// 	// go GC 三色标记法和混合写屏障
	// 	/*
	// 		GC开始时所有对象都是白色，从GC Root（全局变量、栈变量等）开始将直接可达对象标记为灰色。
	// 		然后不断从灰色队列中取出对象，扫描其引用的对象：如果引用对象是白色就标记为灰色，
	// 		当前对象所有引用扫描完成后标记为黑色。重复这个过程直到灰色队列为空。

	// 		GC的根对象到底是什么
	// 		根对象在垃圾回收的术语中又叫做根集合，它是垃圾回收器在标记过程时最先检查的对象，包括:
	// 		全局变量
	// 		执行栈
	// 		寄存器

	// 		STW 是什么意思
	// 		指的是用户代码被完全停止运行
	// 	*/

	/*

		编程
		给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
		在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。
		然而，你可以在 同一天 多次买卖该股票，但要确保你持有的股票不超过一股。 返回 你能获得的 最大 利润 。
		示例 1： 输入：prices = [7,1,5,3,6,4] 输出：7
		解释：在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4。 随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6 - 3 = 3。
	*/