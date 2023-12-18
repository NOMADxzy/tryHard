## 枚举

### 1. 概念
位运算
加法器
模数 (A*B)%K = ((A%K)*(B%K))%K
素数 

### 2. 解题技巧（我的总结）

> 1> 通过位表示信息
> 
| 题目                                                                           | 说明                                      | 实现                                                                            |
|------------------------------------------------------------------------------|-----------------------------------------|-------------------------------------------------------------------------------|
| [458. 可怜的小猪](https://leetcode.cn/problems/poor-pigs/description/) | 共k个时段内检测出结果的信息向量长度 = buckets 的 k+1进制数长度 | [我的提交](https://leetcode.cn/problems/poor-pigs/submissions/488982309/) |

> 2> 素数查找，最小堆，多路归并
>
| 题目                                                                          | 说明                          | 实现                                                                            |
|-----------------------------------------------------------------------------|-----------------------------|-------------------------------------------------------------------------------|
| [313. 超级丑数](https://leetcode.cn/problems/super-ugly-number/description/) | 所有丑数由已有丑数 * 素数列表诞生，从小到大多路归并 | [我的提交](https://leetcode.cn/problems/super-ugly-number/submissions/468697927/) |
| [786. 第 K 个最小的素数分数](https://leetcode.cn/problems/k-th-smallest-prime-fraction/) | 所有丑数由已有丑数 * 素数列表诞生，从小到大多路归并 | [我的提交](https://leetcode.cn/problems/k-th-smallest-prime-fraction/submissions/468733547/) |
| [1508. 子数组和排序后的区间和](https://leetcode.cn/problems/range-sum-of-sorted-subarray-sums/description/) | 多路归并递增产生区间和                 | [我的提交](https://leetcode.cn/problems/range-sum-of-sorted-subarray-sums/submissions/468741186/) |

> 3> 汉明距离，每个维度的汉明距离相互独立
>
| 题目                                                                          | 说明                                                   | 实现                                                                            |
|-----------------------------------------------------------------------------|------------------------------------------------------|-------------------------------------------------------------------------------|
| [477. 汉明距离总和](https://leetcode.cn/problems/total-hamming-distance/description/) | 记录数组中所有数每个汉明位上1出现的个数，则最终每个位导致的距离 = cnt1 *（m - cnt1 ） | [我的提交](https://leetcode.cn/problems/poor-pigs/submissions/488982309/) |

> 4> 最大公约数，最小公倍数，辗转相除法
>
| 题目                                                                          | 说明                               | 实现                                                                            |
|-----------------------------------------------------------------------------|----------------------------------|-------------------------------------------------------------------------------|
| [592. 分数加减运算](https://leetcode.cn/problems/fraction-addition-and-subtraction/description/) | a,b 的最小公倍数 = a * b / (a,b的最大公约数) | [我的提交](https://leetcode.cn/problems/fraction-addition-and-subtraction/submissions/489696067/) |


### 3. 更多练习


### 4. 参考
1. 大部分参考自：[算法基础——枚举](https://blog.csdn.net/weixin_45652283/article/details/131244459?utm_medium=distribute.pc_relevant.none-task-blog-2~default~baidujs_baidulandingword~default-1-131244459-blog-129442726.235^v38^pc_relevant_sort_base3&spm=1001.2101.3001.4242.2&utm_relevant_index=4) 