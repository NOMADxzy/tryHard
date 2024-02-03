## 规律

### 1. 概念

- 模拟题目过程
- 探索问题规律
- 找出隐含条件

### 2. 解题技巧（我的总结）

> 1> 将 问题中的抽象条件 找出实际意思 
> 
| 题目                                                                            | 说明                                     | 实现                                                                            |
|-------------------------------------------------------------------------------|----------------------------------------|-------------------------------------------------------------------------------|
| [649. Dota2 参议院](https://leetcode.cn/problems/dota2-senate/description/) | 题中每个议员最好表现 = 禁止后面最近的对方议员 或 禁止前面最远的对方议员 | [我的提交](https://leetcode.cn/problems/dota2-senate/submissions/489933502/) |
| [1702. 修改后的最大二进制字符串](https://leetcode.cn/problems/maximum-binary-string-after-change/description/) | 10->01代表题中所有的1都可以全部移动到末尾，所有0都可以移动到开头   | [我的提交](https://leetcode.cn/problems/maximum-binary-string-after-change/submissions/494372983/) |
| [1536. 排布二进制网格的最少交换次数](https://leetcode.cn/problems/minimum-swaps-to-arrange-a-binary-grid/description/) | 交换相邻两行 = 冒泡过程，贪心的先满足前面的                | [我的提交](https://leetcode.cn/problems/minimum-swaps-to-arrange-a-binary-grid/submissions/500288067/) |


> 2> 迭代次数过多，寻找周期、规律
> ①二进制周期
> ②字符串循环结
>
| 题目                                                                     | 说明                                                                                      | 实现                                                                            |
|------------------------------------------------------------------------|-----------------------------------------------------------------------------------------|-------------------------------------------------------------------------------|
| [957. N 天后的牢房](https://leetcode.cn/problems/prison-cells-after-n-days/description/) | 题中最多有2^8=256种状态，每个状态仅与上一状态有关，故必然存在周期                                                    | [我的提交](https://leetcode.cn/problems/prison-cells-after-n-days/submissions/490928555/) |
| [1041. 困于环中的机器人](https://leetcode.cn/problems/robot-bounded-in-circle/description/) | 一个周期后方向发生改变或位置不发生改变 则一定会有环                                                              | [我的提交](https://leetcode.cn/problems/robot-bounded-in-circle/submissions/491787188/) |
| 2178. 拆分成最多数目的正偶数之和](https://leetcode.cn/problems/maximum-split-of-positive-even-integers/description/) | 从小到大累加，加上最后一个数后超过了finalNum，一定可以从前面删去一个数 来得到finalNum                                     | [我的提交](https://leetcode.cn/problems/maximum-split-of-positive-even-integers/submissions/499241015/) |
| 466. 统计重复个数](https://leetcode.cn/problems/count-the-repetitions/description/) | 根据贪心思想，优先选择靠前匹配上的字符；记录s1第一个匹配上的idx，存在某一位置i，区间[idx, i]长度是len(s1)的倍数，s2在其中出现了m次，这个区间就是循环结 | [我的提交](https://leetcode.cn/problems/count-the-repetitions/submissions/499941407/) |

> 3> 适当更改题目条件，满足规律
>
| 题目                                                                           | 说明                  | 实现                                                                            |
|------------------------------------------------------------------------------|---------------------|-------------------------------------------------------------------------------|
| [1104. 二叉树寻路](https://leetcode.cn/problems/path-in-zigzag-labelled-binary-tree/description/) | 利用Z字型索引每个元素，再对称变换索引 | [我的提交](https://leetcode.cn/problems/path-in-zigzag-labelled-binary-tree/submissions/492205983/) |



### 3. 更多练习

[找规律](https://blog.csdn.net/qq_49723651/article/details/123485604)

### 4. 参考 