## 排序

### 1. 概念
[![Build Status](https://travis-ci.org/hustcc/JS-Sorting-Algorithm.svg?branch=master)](https://travis-ci.org/hustcc/JS-Sorting-Algorithm)

排序算法是《数据结构与算法》中最基本的算法之一。

排序算法可以分为内部排序和外部排序，内部排序是数据记录在内存中进行排序，而外部排序是因排序的数据很大，一次不能容纳全部的排序记录，在排序过程中需要访问外存。常见的内部排序算法有：**插入排序、希尔排序、选择排序、冒泡排序、归并排序、快速排序、堆排序、基数排序**等。用一张图概括：

1. [冒泡排序](details/1.bubbleSort.md)
2. [选择排序](details/2.selectionSort.md)
3. [插入排序](details/3.insertionSort.md)
4. [希尔排序](details/4.shellSort.md)
5. [归并排序](details/5.mergeSort.md)
6. [快速排序](details/6.quickSort.md)
7. [堆排序](details/7.heapSort.md)
8. [计数排序](details/8.countingSort.md)
9. [桶排序](details/9.bucketSort.md)
10. [基数排序](details/10.radixSort.md)

![十大经典排序算法 概览截图](details/pics/sort.png)


#### 时间复杂度

1. 平方阶 (O(n2)) 排序
   各类简单排序：直接插入、直接选择和冒泡排序。
2. 线性对数阶 (O(nlog2n)) 排序
   快速排序、堆排序和归并排序；
3. O(n1+§)) 排序，§ 是介于 0 和 1 之间的常数。
   希尔排序
4. 线性阶 (O(n)) 排序
   基数排序，此外还有桶、箱排序。


#### 关于稳定性

稳定的排序算法：冒泡排序、插入排序、归并排序和基数排序。

不是稳定的排序算法：选择排序、快速排序、希尔排序、堆排序。


#### 名词解释：
**In-place**：占用常数内存，不占用额外内存

**Out-place**：占用额外内存

**稳定性**：排序后 2 个相等键值的顺序和排序之前它们的顺序相同

### 2. 解题技巧（我的总结）

> 1> 通过排序将数组分为左右不同性质的部分，简化运算
> 
| 题目                                                                            | 说明                         | 实现                                                                            |
|-------------------------------------------------------------------------------|----------------------------|-------------------------------------------------------------------------------|
| [462. 最小操作次数使数组元素相等 II](https://leetcode.cn/problems/minimum-moves-to-equal-array-elements-ii/) | 排序后，操作次数 = 左边增加的值 + 右边减小的值 | [我的提交](https://leetcode.cn/problems/minimum-moves-to-equal-array-elements-ii/submissions/478474363/) |

> 2> 二重排序，第一个条件相同时，再通过第二个条件排序
>
| 题目                                                                            | 说明                  | 实现                                                                            |
|-------------------------------------------------------------------------------|---------------------|-------------------------------------------------------------------------------|
| [524. 通过删除字母匹配到字典里最长单词](https://leetcode.cn/problems/longest-word-in-dictionary-through-deleting/description/) | 双指针法判断是否是子序列        | [我的提交](https://leetcode.cn/problems/longest-word-in-dictionary-through-deleting/submissions/478642292/) |
| [939. 最小面积矩形](https://leetcode.cn/problems/minimum-area-rectangle/description/) | 按层从下往上、从左往右记录所有同层y对 | [我的提交](https://leetcode.cn/problems/minimum-area-rectangle/submissions/479072887/) |

> 3> 前缀树，用于排序和判断前缀关系
```go
type Trie struct {
	Children [26]*Trie
	isEnd    bool
}
```
| 题目                                                                       | 说明                         | 实现                                                                            |
|--------------------------------------------------------------------------|----------------------------|-------------------------------------------------------------------------------|
| [208. 实现 Trie (前缀树)](https://leetcode.cn/problems/implement-trie-prefix-tree/description/) | 前缀树 | [我的提交](https://leetcode.cn/problems/implement-trie-prefix-tree/submissions/478776637/) |



### 3. 更多练习


### 4. 参考
1. 开源项目地址：[https://github.com/hustcc/JS-Sorting-Algorithm](https://github.com/hustcc/JS-Sorting-Algorithm)，整理人 [hustcc](https://github.com/hustcc)。

2. GitBook 在线阅读地址：[https://sort.hust.cc/](https://sort.hust.cc/)。

3. 本项目使用 [lint-md](https://github.com/hustcc/lint-md) 进行中文 Markdown 文件的格式检查，务必在提交 Pr 之前，保证 Markdown 格式正确。 