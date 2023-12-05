## 枚举

### 1. 概念
数位DP其实就是一种优化之后的暴力枚举方法，有点类似于记忆化搜索。字面意思是就是在数位上进行DP枚举，对于数字的每一位进行枚举，通过相应的约束条件看这次枚举是否符合题意。

对于有一类题型：给定闭区间 \[l, r\]，需要求出区间中满足某种条件的数的总数，这个区间如果比较大，是无法通过简单的暴力枚举解决的，这时候就需要使用数位DP加以优化，具体通过 [2376. 统计特殊整数](https://leetcode.cn/problems/count-special-integers/) 来理解一下，主要参考灵神的 [视频讲解](https://www.bilibili.com/video/BV1rS4y1s721/?vd_source=286032bc2c5715c8b50b608028ce57df)




### 2. 解题技巧（我的总结）

> 1> 数位DP，求小于N的符合条件的数的个数，注意以下几点：
> 用一个 target数组 存储N的每一位
> 用一个 preLimit 变量表示前面的是否处于临界状态，当且仅当preLimit为true且当前位到达target对应位时，nextLimit为true
> 用一个 pos 表示当前处理到数字的第多少位，pos到len(target)时结束递归
> 
> 
| 题目                                                                   | 说明                          | 实现                                                                            |
|----------------------------------------------------------------------|-----------------------------|-------------------------------------------------------------------------------|
| [788. 旋转数字](https://leetcode.cn/problems/rotated-digits/description/) | dp时用preValid表示前面是否已经含有2569  | [我的提交](https://leetcode.cn/problems/rotated-digits/description/) |
| [2376. 统计特殊整数](https://leetcode.cn/problems/count-special-integers/description/) | 使用一个mark，第i位标记i是否被使用，从而压缩状态 | [我的提交](https://leetcode.cn/problems/count-special-integers/submissions/473515507/) |
| [60. 排列序列](https://leetcode.cn/problems/permutation-sequence/description/) | 从左到右递归累计排名                  | [我的提交](https://leetcode.cn/problems/permutation-sequence/submissions/485525438/) |

> 2> 前序、中序、后序遍历，二叉搜索树
>
| 题目                                                                   | 说明               | 实现                                                                            |
|----------------------------------------------------------------------|------------------|-------------------------------------------------------------------------------|
| [2476. 二叉搜索树最近节点查询](https://leetcode.cn/problems/closest-nodes-queries-in-a-binary-search-tree/description/) | 中序遍历 + 排序 + 二分查找 | [我的提交](https://leetcode.cn/problems/closest-nodes-queries-in-a-binary-search-tree/submissions/486976544/) |

### 3. 更多练习
- [ ] [233. 数字 1 的个数](https://leetcode.cn/problems/number-of-digit-one/)：[题解](https://leetcode.cn/problems/number-of-digit-one/solution/by-endlesscheng-h9ua/)
- [ ] [面试题 17.06. 2出现的次数](https://leetcode.cn/problems/number-of-2s-in-range-lcci/)：[题解](https://leetcode.cn/problems/number-of-2s-in-range-lcci/solution/by-endlesscheng-x4mf/)
- [ ] [600. 不含连续1的非负整数](https://leetcode.cn/problems/non-negative-integers-without-consecutive-ones/)：[题解](https://leetcode.cn/problems/non-negative-integers-without-consecutive-ones/solution/by-endlesscheng-1egu/)
- [x] [902. 最大为 N 的数字组合](https://leetcode.cn/problems/numbers-at-most-n-given-digit-set/)
- [ ] 1012.至少有 1 位重复的数字
- [ ] 1067.范围内的数字计数
- [ ] 1397.找到所有好字符串（有难度，需要结合一个知名字符串算法）

### 4. 参考
1. [数位DP学习整理——数位DP看完这篇你就会了 CSDN](https://blog.csdn.net/hzf0701/article/details/116717851)
2. [数位 DP 通用模板，附题单（Python/Java/C++/Go）-- 0x3F](https://leetcode.cn/problems/count-special-integers/solution/shu-wei-dp-mo-ban-by-endlesscheng-xtgx/)