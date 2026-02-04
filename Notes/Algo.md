## 基础算法



1. 冒泡排序 
    - 序列中相邻元素两两对比，把最大或者最小的移动到最后      
    - 时间复杂度    o(n^2)   
    - 空间复杂度    o(1)
2. 选择排序 
    - 循环遍历序列，选出最大或者最小的，跟最前面的元素交换  
    - 时间    o(n^2)  
    - 空间    o(1)
3. 插入排序 
    - 把序列划分已排序和未排序两部分，每次从未排序队列取一个元素插入到已排序部分
    - 时间    o(n^2)  
    - 空间    o(1)
4. 希尔排序
    - 是插入排序的一种改进版本，也称为缩小增量排序。它通过将待排序的列表分成若干子列表，对每个子列表进行插入排序，逐步缩小子列表的间隔，最终完成
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

	// 基数排序 根据不同位依次进行排序 最终遍历了每一位得到有序的队列

