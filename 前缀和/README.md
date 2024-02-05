## 前缀和

### 1. 概念

    前缀和算法是一种用于高效计算数组前缀和的算法。前缀和是指从数组的起始位置到某一位置的所有元素的和。

下面是前缀和算法的基本步骤：

    创建一个与原始数组相同长度的前缀和数组。初始时，前缀和数组的第一个元素与原始数组的第一个元素相同。

    从第二个元素开始，遍历原始数组，计算每个位置处的前缀和，即将前一个位置的前缀和与当前位置的元素相加。

    将计算得到的前缀和存储到前缀和数组的相应位置。

    完成遍历后，前缀和数组中存储了原始数组每个位置的前缀和值。

### 2. 解题技巧（我的总结）

> 1> 复杂问题求多个子前缀和简化算法过程
> 
| 题目                                                                       | 说明             | 实现                                                                            |
|--------------------------------------------------------------------------|----------------|-------------------------------------------------------------------------------|
| [1664. 生成平衡数组的方案数](https://leetcode.cn/problems/ways-to-make-a-fair-array/description/) | 对奇数位和偶数位分别求前缀和 | [我的提交](https://leetcode.cn/problems/ways-to-make-a-fair-array/submissions/477957501/) |

> 2> 前缀和➕哈希表解决子数组目标和问题（容斥原理）
>
| 题目                                                                       | 说明                        | 实现                                                                            |
|--------------------------------------------------------------------------|---------------------------|-------------------------------------------------------------------------------|
| [1074. 元素和为目标值的子矩阵数量](https://leetcode.cn/problems/number-of-submatrices-that-sum-to-target/description/) | 容斥原理                      | [我的提交](https://leetcode.cn/problems/number-of-submatrices-that-sum-to-target/submissions/491814427/) |
| [1010. 总持续时间可被 60 整除的歌曲](https://leetcode.cn/problems/pairs-of-songs-with-total-durations-divisible-by-60/description/) | 容斥原理                      | [我的提交](https://leetcode.cn/problems/pairs-of-songs-with-total-durations-divisible-by-60/submissions/491964237/) |
| [1442. 形成两个异或相等数组的三元组数目](https://leetcode.cn/problems/count-triplets-that-can-form-two-arrays-of-equal-xor/description/) | 注意初始值{0:{-1}, arr[0]:{0}} | [我的提交](https://leetcode.cn/problems/count-triplets-that-can-form-two-arrays-of-equal-xor/submissions/493361982/) |
| [2001. 可互换矩形的组数](https://leetcode.cn/problems/number-of-pairs-of-interchangeable-rectangles/description/) | 可替换的矩形，对应着唯一的最简比例，即组合计数问题，采用用容斥定理解决 | [我的提交](https://leetcode.cn/problems/number-of-pairs-of-interchangeable-rectangles/submissions/500720933/) |

### 3. 更多练习


### 4. 参考
1. 大部分参考自：[前缀和算法](https://blog.csdn.net/m0_56069910/article/details/132743061) 