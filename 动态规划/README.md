## 记忆化搜索到递推

### 1. 概念

**记忆化搜索**：递归搜索 + 保存中间计算结果，**如何思考状态转移是重点**

**递归**：1:1 翻译 DFS 为 DP


**三步走** :fire:

- 思考回溯怎么写：① 入参是什么；② 递归到哪里；③ 注意递归边界和入口
- 改成记忆化搜索
- 1:1 翻译成递推：① dfs 改成 dp 数组；② 递归改成循环；③ 递归边界对应 dp 数组的初始化

时间复杂度：状态个数 * 单个状态计算时间



### 2. 解题技巧（我的总结）

> 1> 复杂问题无法一次性状态转移的可以分成多个子问题，分别进行状态转移，最后合并结果
> 
| 题目                                                                         | 说明         | 实现                                                                            |
|----------------------------------------------------------------------------|------------|-------------------------------------------------------------------------------|
| [764. 最大加号标志](https://leetcode.cn/problems/largest-plus-sign/description/) | 分成四个方向分别求解 | [我的提交](https://leetcode.cn/problems/largest-plus-sign/submissions/473771185/) |

> 2> 使用两个变量滚动记录 dp，优化空间
>
| 题目          | 说明       | 实现                                                                            |
|-------------|----------|-------------------------------------------------------------------------------|
| [70. 爬楼梯](https://leetcode.cn/problems/climbing-stairs/description/) | 只记录上一个状态 | [我的提交](https://leetcode.cn/problems/climbing-stairs/submissions/464582419/) |


> 3> 一个状态划分两个或以上的子状态，分别对每个子状态进行转移，合并得到结果
>
| 题目                                                                             | 说明                        | 实现                                                                            |
|--------------------------------------------------------------------------------|---------------------------|-------------------------------------------------------------------------------|
| [357. 统计各位数字都不同的数字个数](https://leetcode.cn/problems/count-numbers-with-unique-digits/description/) | 划分是/否含0两个状态               | [我的提交](https://leetcode.cn/problems/count-numbers-with-unique-digits/submissions/473822249/) |
| [714. 买卖股票的最佳时机含手续费](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/description/) | 划分是/否持有股票两个状态             | [我的提交](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/submissions/466212396/) |
| [576. 出界的路径数](https://leetcode.cn/problems/out-of-boundary-paths/description/) | 划分网格位置 m*n 个状态            | [我的提交](https://leetcode.cn/problems/out-of-boundary-paths/submissions/474130421/) |
| [494. 目标和](https://leetcode.cn/problems/target-sum/)                           | 划分当前和 1000 + 1 + 1000 个状态 | [我的提交](https://leetcode.cn/problems/target-sum/submissions/474145931/) |
| [740. 删除并获得点数](https://leetcode.cn/problems/delete-and-earn/description/)      | 排序后划分 是/否 选择当前数 2 个状态     | [我的提交](https://leetcode.cn/problems/delete-and-earn/submissions/474333479/) |

> 4> 可对状态进行分类（26个字母等）从而大大降低算法复杂度
>
| 题目                                                                    | 说明            | 实现                                                                          |
|-----------------------------------------------------------------------|---------------|-----------------------------------------------------------------------------|
| [467. 爬楼梯](https://leetcode.cn/problems/unique-substrings-in-wraparound-string/description/) | 针对26个字母分类巧妙去重 | [我的题解](https://leetcode.cn/problems/unique-substrings-in-wraparound-string/solutions/2481172/ji-lu-26ge-zi-mu-jie-wei-de-zi-chuan-de-4v1uf/) |

> 5> 区间类型的动态规划
>
| 题目                                                                                | 说明                     | 实现                                                                            |
|-----------------------------------------------------------------------------------|------------------------|-------------------------------------------------------------------------------|
| [553. 最优除法](https://leetcode.cn/problems/optimal-division/submissions/474107247/) | 除法 = 左区间 / 右区间         | [我的提交](https://leetcode.cn/problems/optimal-division/submissions/474107247/) |
| [516. 最优除法](https://leetcode.cn/problems/longest-palindromic-subsequence/description/) | 当前区间 = 左端点 + 子区间 + 右端点 | [我的提交](https://leetcode.cn/problems/longest-palindromic-subsequence/submissions/474136039/) |



### 3. 更多练习

| 题目                                                         | 说明                                                         | 题解                                                         |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| [1043. 分隔数组以得到最大和](https://leetcode.cn/problems/partition-array-for-maximum-sum/) | 枚举最后一段的子数组下标 j，dfs(i) 表示 arr[0...i] 做分隔变换得到的最大元素和 | [0x3F](https://leetcode.cn/problems/partition-array-for-maximum-sum/solution/jiao-ni-yi-bu-bu-si-kao-dong-tai-gui-hua-rq5i/) |
| [1105. 填充书架](https://leetcode.cn/problems/filling-bookcase-shelves/) | dfs(i) 表示摆放 books[0...i] 可以做到的最小高度，枚举最后一层可以放置的书 | [通过](https://leetcode.cn/submissions/detail/426015835/)    |
| [1335. 工作计划的最低难度](https://leetcode.cn/problems/minimum-difficulty-of-a-job-schedule/) | dfs(i, j) 表示 j 天完成 job[0...i] 的最小难度，有限制的枚举最后一个天的工作 job[k...i] 取最大值 mx | [通过](https://leetcode.cn/submissions/detail/433036520/)    |
| [1416. 恢复数组](https://leetcode.cn/problems/restore-the-array/) | dfs(i) 表示 s[0...i] 字符串分割成数组的方案数，枚举末尾段字符子串 t 可能得分割方式，不能含有前导 0 并且不能超过 k | [通过](https://leetcode.cn/submissions/detail/427057510/)    |




### 4. 参考

- [动态规划入门：从记忆化搜索到递推【基础算法精讲 17】](https://www.bilibili.com/video/BV1Xj411K7oF)
- [教你一步步思考动态规划！（Python/Java/C++/Go）](https://leetcode.cn/problems/partition-array-for-maximum-sum/solution/jiao-ni-yi-bu-bu-si-kao-dong-tai-gui-hua-rq5i/)

