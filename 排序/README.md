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
| 题目                                                                            | 说明                                      | 实现                                                                            |
|-------------------------------------------------------------------------------|-----------------------------------------|-------------------------------------------------------------------------------|
| [462. 最小操作次数使数组元素相等 II](https://leetcode.cn/problems/minimum-moves-to-equal-array-elements-ii/) | 排序后，操作次数 = 左边增加的值 + 右边减小的值              | [我的提交](https://leetcode.cn/problems/minimum-moves-to-equal-array-elements-ii/submissions/478474363/) |
| [1996. 游戏中弱角色的数量](https://leetcode.cn/problems/the-number-of-weak-characters-in-the-game/description/) | 排序后，从右往左，在每个角色严格右区间内判断，利用dpRightMax记录历史 | [我的提交](https://leetcode.cn/problems/the-number-of-weak-characters-in-the-game/submissions/480770831/) |

> 2> 二重(多重)排序，第一个条件相同时，再通过第二个条件...排序
>
| 题目                                                                            | 说明                  | 实现                                                                            |
|-------------------------------------------------------------------------------|---------------------|-------------------------------------------------------------------------------|
| [524. 通过删除字母匹配到字典里最长单词](https://leetcode.cn/problems/longest-word-in-dictionary-through-deleting/description/) | 双指针法判断是否是子序列        | [我的提交](https://leetcode.cn/problems/longest-word-in-dictionary-through-deleting/submissions/478642292/) |
| [939. 最小面积矩形](https://leetcode.cn/problems/minimum-area-rectangle/description/) | 按层从下往上、从左往右记录所有同层y对 | [我的提交](https://leetcode.cn/problems/minimum-area-rectangle/submissions/479072887/) |
| [1366. 通过投票对团队排名](https://leetcode.cn/problems/rank-teams-by-votes/description/) | 自定义多维排序             | [我的提交](https://leetcode.cn/problems/rank-teams-by-votes/submissions/479647745/) |
| [1424. 对角线遍历 II](https://leetcode.cn/problems/diagonal-traverse-ii/description/1424) | 第一维：i+j的值，第二维：j     | [我的提交](https://leetcode.cn/problems/diagonal-traverse-ii/submissions/) |
> 3> 前缀树，用于排序和判断前缀关系，（优化212）
```go
type Trie struct {
	Children [26]*Trie
	isEnd    bool
}
```
| 题目                                                                      | 说明                       | 实现                                                                                    |
|-------------------------------------------------------------------------|--------------------------|---------------------------------------------------------------------------------------|
| [208. 实现 Trie (前缀树)](https://leetcode.cn/problems/implement-trie-prefix-tree/description/) | 前缀树                      | [我的提交](https://leetcode.cn/problems/implement-trie-prefix-tree/submissions/478776637/) |
| [1268. 搜索推荐系统](https://leetcode.cn/problems/search-suggestions-system/description/) | 一眼前缀树                    | [我的提交](https://leetcode.cn/problems/search-suggestions-system/submissions/479543984/) |
| [1233. 删除子文件夹](https://leetcode.cn/problems/remove-sub-folders-from-the-filesystem/description/) | 使用哈希表存储子前缀树，ref存储在列表中的位置 | [我的提交](https://leetcode.cn/problems/remove-sub-folders-from-the-filesystem/submissions/482986495/) |
| [140. 单词拆分 II](https://leetcode.cn/problems/word-break-ii/description/) | 前缀树                      | [我的提交](https://leetcode.cn/problems/word-break-ii/submissions/487470460/) |
| [212. 单词搜索 II](https://leetcode.cn/problems/word-search-ii/description/) | 对已搜索到的单词剪去，大大降低搜索时间      | [我的提交](https://leetcode.cn/problems/word-search-ii/submissions/487614649/) |

> 4> 数对问题，排序结合哈希表，从小到大在哈希表中依次去除数对，直到不够或终止（注意要删除一对，不能只删除一个）
>
| 题目                                                                            | 说明    | 实现                                                                            |
|-------------------------------------------------------------------------------|-------|-------------------------------------------------------------------------------|
| [954. 二倍数对数组](https://leetcode.cn/problems/array-of-doubled-pairs/description/) | 排序，哈希 | [我的提交](https://leetcode.cn/problems/array-of-doubled-pairs/submissions/) |
| [1296. 划分数组为连续数字的集合](https://leetcode.cn/problems/divide-array-in-sets-of-k-consecutive-numbers/description/) | 排序，哈希 | [我的提交](https://leetcode.cn/problems/divide-array-in-sets-of-k-consecutive-numbers/submissions/479558102/) |
| [532. 数组中的 k-diff 数对](https://leetcode.cn/problems/k-diff-pairs-in-an-array/description/) | 排序，哈希 | [我的提交](https://leetcode.cn/problems/k-diff-pairs-in-an-array/submissions/) |

> 5> 按照某种性质进行排序
>
| 题目                                                                            | 说明                            | 实现                                                                            |
|-------------------------------------------------------------------------------|-------------------------------|-------------------------------------------------------------------------------|
| [2327. 知道秘密的人数](https://leetcode.cn/problems/number-of-people-aware-of-a-secret/description/) | 按知道秘密的顺序维护一个队列，每天新增的=队列前若干个之和 | [我的提交](https://leetcode.cn/problems/number-of-people-aware-of-a-secret/submissions/479881461/) |

> 6> 分类排序，复杂问题
>
| 题目                                                                            | 说明                                      | 实现                                                                            |
|-------------------------------------------------------------------------------|-----------------------------------------|-------------------------------------------------------------------------------|
| [1727. 重新排列后的最大子矩阵](https://leetcode.cn/problems/largest-submatrix-with-rearrangements/description/) | 记录每列每个位置开始的最大连续1个数，分类讨论从每个位置开始的最大矩形(排序) | [我的提交](https://leetcode.cn/problems/largest-submatrix-with-rearrangements/submissions/479963584/) |

> 7> 变种排序
>
| 题目                                                                            | 说明                                                               | 实现                                                                            |
|-------------------------------------------------------------------------------|------------------------------------------------------------------|-------------------------------------------------------------------------------|
| [899. 有序队列](https://leetcode.cn/problems/orderly-queue/description/) | k>2时一定能使s有序，[参考](https://leetcode.cn/problems/orderly-queue/solutions/1717847/you-xu-dui-lie-by-capital-worker-p1oz/) | [我的提交](https://leetcode.cn/problems/orderly-queue/submissions/491207673/) |
| [1535. 找出数组游戏的赢家](https://leetcode.cn/problems/find-the-winner-of-an-array-game/description/) | 等价于冒泡排序                                                          | [我的提交](https://leetcode.cn/problems/find-the-winner-of-an-array-game/submissions/493729970/) |
| [761. 特殊的二进制序列](https://leetcode.cn/problems/special-binary-string/description/) | 特殊串一定是1***0这种形式，相邻交换即冒泡排序，且特殊串的所有特殊字串一定是相邻的                      | [我的提交](https://leetcode.cn/problems/special-binary-string/submissions/501935866/) |
| [2948. 交换得到字典序最小的数组](https://leetcode.cn/problems/make-lexicographically-smallest-array-by-swapping-elements/description/) | 等价于部分序列冒泡排序，通过pairs数组排序后，得到limit内的序列，在该序列上排序 | [我的提交](https://leetcode.cn/problems/make-lexicographically-smallest-array-by-swapping-elements/submissions/506656084/) |

### 3. 更多练习


### 4. 参考
1. 开源项目地址：[https://github.com/hustcc/JS-Sorting-Algorithm](https://github.com/hustcc/JS-Sorting-Algorithm)，整理人 [hustcc](https://github.com/hustcc)。
2. GitBook 在线阅读地址：[https://sort.hust.cc/](https://sort.hust.cc/)。
3. 本项目使用 [lint-md](https://github.com/hustcc/lint-md) 进行中文 Markdown 文件的格式检查，务必在提交 Pr 之前，保证 Markdown 格式正确。 
4. 总库：[tryHard](https://github.com/NOMADxzy/tryHard)