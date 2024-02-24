## 枚举

### 1. 概念
枚举算法是我们在日常中使用到的最多的一个算法，它的核心思想就是:枚举所有的可能。
枚举法的本质就是从所有候选答案中去搜索正确的解,使用该算法需要满足两个条件：
(1)可预先确定候选答案的数量；
(2)候选答案的范围在求解之前必须有一个确定的集合

在实际问题中， 有些变量的取值被限定在一个有限的范围内。例如，一个星期内只有七天，一年只有十二个月， 一个班每周有六门课程等等。
枚举算法的优点

    枚举算法一般是现实生活问题的“直译”，所以比较直观，易于理解
    枚举算法建立在考察大量状态、甚至是穷举所有状态的基础上，所有算法的真确性容易证明

枚举算法的缺点 

    枚举算法的效率取决与枚举状态的数量和单个状态枚举的代价，所以枚举效率比较低

### 2. 解题技巧（我的总结）

> 1> 数组中每个元素做若干种可能的操作，直接枚举, 2的m次方种可能
> 
| 题目                                                                            | 说明                                                    | 实现                                                                            |
|-------------------------------------------------------------------------------|-------------------------------------------------------|-------------------------------------------------------------------------------|
| [1049. 最后一块石头的重量 II](https://leetcode.cn/problems/last-stone-weight-ii/description/) | 每块石头的重量在前面所有可能重量的基础上要么被加上要么被减去，枚举所有可能2^len(stones)种可能 | [我的提交](https://leetcode.cn/problems/last-stone-weight-ii/submissions/477223198/) |
| [1986. 完成任务的最少工作时间段](https://leetcode.cn/problems/maximum-product-of-the-length-of-two-palindromic-subsequences/description) | 通过二进制数记录元素选择状况，state1 & state2 == 0 即不相交，枚举出所有的回文状态即可 | [我的提交](https://leetcode.cn/problems/maximum-product-of-the-length-of-two-palindromic-subsequences/submissions/504255231/) |

> 2> 数组中选择任意个元素，2的m次方中可能，通过 排序，二分查找 优化算法
>
| 题目                                                                        | 说明                           | 实现                                                                            |
|---------------------------------------------------------------------------|------------------------------|-------------------------------------------------------------------------------|
| [1774. 最接近目标价格的甜点成本](https://leetcode.cn/problems/closest-dessert-cost/description/) | 每个配料最多2份 等价于 重复两次：在配料表中任取若干份 | [我的提交](https://leetcode.cn/problems/closest-dessert-cost/submissions/478095922/) |

> 3> 分类讨论
>
| 题目                                                                        | 说明                             | 实现                                                                            |
|---------------------------------------------------------------------------|--------------------------------|-------------------------------------------------------------------------------|
| [794. 有效的井字游戏](https://leetcode.cn/problems/valid-tic-tac-toe-state/description/) | XO不能同时win，Xwin，Owin,都没Win，四种情况 | [我的提交](https://leetcode.cn/problems/valid-tic-tac-toe-state/submissions/490566208/) |

> 4> 构造题，找到合适的方法、排列、枚举顺序一步一步构造
>
| 题目                                                                        | 说明                                                    | 实现                                                                            |
|---------------------------------------------------------------------------|-------------------------------------------------------|-------------------------------------------------------------------------------|
| [667. 优美的排列 II](https://leetcode.cn/problems/beautiful-arrangement-ii/description/) | 1，2，3... n,n-1,n-2,... 这两类依次交错排列能得到最多的k，根据k的大小选择交错多少步 | [我的提交](https://leetcode.cn/problems/beautiful-arrangement-ii/submissions/491176947/) |

> 5> 利用（对称、奇偶、字母等）性质优化枚举次数
>
| 题目                                                                        | 说明       | 实现                                                                            |
|---------------------------------------------------------------------------|----------|-------------------------------------------------------------------------------|
| [1492. n 的第 k 个因子](https://leetcode.cn/problems/the-kth-factor-of-n/description/) | 只需枚举前半部分 | [我的提交](https://leetcode.cn/problems/the-kth-factor-of-n/submissions/493607495/) |
| [1930. 长度为 3 的不同回文子序列](https://leetcode.cn/problems/unique-length-3-palindromic-subsequences/description/) | 枚举所有字母   | [我的提交](https://leetcode.cn/problems/unique-length-3-palindromic-subsequences/submissions/496212668/) |


### 3. 更多练习


### 4. 参考
1. 大部分参考自：[算法基础——枚举](https://blog.csdn.net/weixin_45652283/article/details/131244459?utm_medium=distribute.pc_relevant.none-task-blog-2~default~baidujs_baidulandingword~default-1-131244459-blog-129442726.235^v38^pc_relevant_sort_base3&spm=1001.2101.3001.4242.2&utm_relevant_index=4) 
2. 总库：[tryHard](https://github.com/NOMADxzy/tryHard)