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

> 1> 复杂问题无法一次性状态转移的可以分成多个子问题，分别进行状态转移，最后合并结果, （或者仅对其中一个子问题状态转移）
> 
| 题目                                                                         | 说明                                              | 实现                                                                            |
|----------------------------------------------------------------------------|-------------------------------------------------|-------------------------------------------------------------------------------|
| [764. 最大加号标志](https://leetcode.cn/problems/largest-plus-sign/description/) | 分成四个方向分别求解                                      | [我的提交](https://leetcode.cn/problems/largest-plus-sign/submissions/473771185/) |
| [838. 推多米诺](https://leetcode.cn/problems/push-dominoes/description/)   | 分成两个个方向分别转移                                     | [我的提交](https://leetcode.cn/problems/push-dominoes/submissions/476039969/) |
| [1031. 两个非重叠子数组的最大和](https://leetcode.cn/problems/maximum-sum-of-two-non-overlapping-subarrays/description/) | 仅对(i:)区间内长度为firstlen和secondLen的最大长度状态转移         | [我的提交](https://leetcode.cn/problems/maximum-sum-of-two-non-overlapping-subarrays/submissions/477055048/) |
| [1139. 最大的以 1 为边界的正方形](https://leetcode.cn/problems/largest-1-bordered-square/submissions/477414200/) | 分别求解四个方向连续1的个数                                  | [我的提交](https://leetcode.cn/problems/largest-1-bordered-square/submissions/477414200/) |
| [1525. 字符串的好分割数目](https://leetcode.cn/problems/number-of-good-ways-to-split-a-string/description/) | 分别求解i左右边不同字符的个数                                 | [我的提交](https://leetcode.cn/problems/number-of-good-ways-to-split-a-string/submissions/477773727/) |
| [1749. 任意子数组和的绝对值的最大值](https://leetcode.cn/problems/maximum-absolute-sum-of-any-subarray/description/)  | 分别求解子数组的最小和和最大和，简化状态转移                          | [我的提交](https://leetcode.cn/problems/maximum-absolute-sum-of-any-subarray/submissions/478023082/) |
| [1888. 使二进制字符串字符交替的最少反转次数](https://leetcode.cn/problems/minimum-number-of-flips-to-make-the-binary-string-alternating/description/)  | 看成左右两个部分分别为交替串，从左到右、从右到左分别求解，细分成为0/1结尾的交替串两个子状态 | [我的提交](https://leetcode.cn/problems/minimum-number-of-flips-to-make-the-binary-string-alternating/submissions/478169186/) |
| [123. 买卖股票的最佳时机 III](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/description/)   | 分成求解某一点左边卖出能获得的最大收益 + 某一点右边买入能获得的最大收益 之和        | [我的提交](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/submissions/485881364/) |

> 2> 使用两个变量滚动记录 dp，优化空间
>
| 题目          | 说明       | 实现                                                                            |
|-------------|----------|-------------------------------------------------------------------------------|
| [70. 爬楼梯](https://leetcode.cn/problems/climbing-stairs/description/) | 只记录上一个状态 | [我的提交](https://leetcode.cn/problems/climbing-stairs/submissions/464582419/) |


> 3> 一个状态划分两个或以上的子状态，分别对每个子状态进行转移，合并得到结果
>
| 题目                                                                                            | 说明                                                             | 实现                                                                                                               |
|-----------------------------------------------------------------------------------------------|----------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------|
| [357. 统计各位数字都不同的数字个数](https://leetcode.cn/problems/count-numbers-with-unique-digits/description/) | 划分是/否含0两个状态                                                    | [我的提交](https://leetcode.cn/problems/count-numbers-with-unique-digits/submissions/473822249/)                     |
| [714. 买卖股票的最佳时机含手续费](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/description/) | 划分是/否持有股票两个状态                                                  | [我的提交](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/submissions/466212396/) |
| [576. 出界的路径数](https://leetcode.cn/problems/out-of-boundary-paths/description/)                | 划分网格位置 m*n 个状态                                                 | [我的提交](https://leetcode.cn/problems/out-of-boundary-paths/submissions/474130421/)                                |
| [494. 目标和](https://leetcode.cn/problems/target-sum/)                                          | 划分当前和 1000 + 1 + 1000 个状态                                      | [我的提交](https://leetcode.cn/problems/target-sum/submissions/474145931/)                                           |
| [740. 删除并获得点数](https://leetcode.cn/problems/delete-and-earn/description/)                     | 排序后划分 是/否 选择当前数 2 个状态                                          | [我的提交](https://leetcode.cn/problems/delete-and-earn/submissions/474333479/)                                      |
| [926. 将字符串翻转到单调递增](https://leetcode.cn/problems/flip-string-to-monotone-increasing/description/) | 当前位分别为 '0','1' 时符合条件的最优值                                       | [我的提交](https://leetcode.cn/problems/flip-string-to-monotone-increasing/submissions/476950683/)                   |
| [935. 骑士拨号器](https://leetcode.cn/problems/knight-dialer/description/)                         | 长度为i的号码划分成从0~9开始的10个子状态，每个子状态之间通过马字形转移                         | [我的提交](https://leetcode.cn/problems/knight-dialer/submissions/476959108/)                                        |
| [1155. 掷骰子等于目标和的方法数](https://leetcode.cn/problems/number-of-dice-rolls-with-target-sum/description/) | dp(i)(j)表示第i次累计和为j的种类数，1=<j<=target                            | [我的提交](https://leetcode.cn/problems/number-of-dice-rolls-with-target-sum/submissions/477465423/)                 |
| [1186. 删除一次得到子数组最大和](https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion/description/) | dp(i)(0),dp(i)(1)分别表示前i个元素，删除和不删元素的最大子数组和                      | [我的提交](https://leetcode.cn/problems/number-of-dice-rolls-with-target-sum/submissions/477465423/)                 |
| [1262. 可被三整除的最大和](https://leetcode.cn/problems/greatest-sum-divisible-by-three/description/)    | dp(i)(0),dp(i)(1), dp(i)(2)分别表示前i个元素%3的最大和                     | [我的提交](https://leetcode.cn/problems/greatest-sum-divisible-by-three/submissions/477510062/)                 |
| [1504. 统计全 1 子矩形](https://leetcode.cn/problems/count-submatrices-with-all-ones/description/)          | dp(i,j,h) 表示右下端点为ij, 高度为h的矩形的最大宽度， dp(i,j,h) = dp(i,j-1,h) + 1 | [我的提交](https://leetcode.cn/problems/count-submatrices-with-all-ones/submissions/477754022/)                 |
| [1567. 乘积为正数的最长子数组长度](https://leetcode.cn/problems/maximum-length-of-subarray-with-positive-product/description/)   | dp(i)(0), dp(i)(1)分别表示i结尾乘积正/负数的最长子数组长度                        | [我的提交](https://leetcode.cn/problems/maximum-length-of-subarray-with-positive-product/submissions/477779171/)                 |
| [1594. 矩阵的最大非负积](https://leetcode.cn/problems/maximum-non-negative-product-in-a-matrix/description/)          | dp(i)(j)(0), dp(i)(j)(1)分别表示i,j位置累积非负/负的绝对值最大的乘积               | [我的提交](https://leetcode.cn/problems/maximum-non-negative-product-in-a-matrix/submissions/477845081/)                 |
| [1638. 统计只差一个字符的子串数目](https://leetcode.cn/problems/count-substrings-that-differ-by-one-character/description/)                | dp(i)(j)(0), dp(i)(j)(1)分别表示以i,j位置结尾累积相同/差1字符的最长子串长度           | [我的提交](https://leetcode.cn/problems/count-substrings-that-differ-by-one-character/submissions/477940428/)                 |
| [1824. 最少侧跳次数](https://leetcode.cn/problems/minimum-sideway-jumps/description/)                 | 划分成青蛙在三个跑道上的侧跳次数分别进行状态转移                                       | [我的提交](https://leetcode.cn/problems/minimum-sideway-jumps/submissions/478099769/)                 |

> 4> 可对状态进行分类（26个字母等）、对数据进行排序，从而能够动态规划
>
| 题目                                                                    | 说明            | 实现                                                                          |
|-----------------------------------------------------------------------|---------------|-----------------------------------------------------------------------------|
| [467. 环绕字符串中唯一的子字符串](https://leetcode.cn/problems/unique-substrings-in-wraparound-string/description/) | 针对26个字母分类巧妙去重 | [我的题解](https://leetcode.cn/problems/unique-substrings-in-wraparound-string/solutions/2481172/ji-lu-26ge-zi-mu-jie-wei-de-zi-chuan-de-4v1uf/) |
| [2008. 出租车的最大盈利](https://leetcode.cn/problems/maximum-earnings-from-taxi/description/) | 根据到达地点从小到大排序  | [我的题解](https://leetcode.cn/problems/maximum-earnings-from-taxi/submissions/480784919/) |

> 5> 区间类型的动态规划
>
| 题目                                                                                                    | 说明                     | 实现                                                                            |
|-------------------------------------------------------------------------------------------------------|------------------------|-------------------------------------------------------------------------------|
| [553. 最优除法](https://leetcode.cn/problems/optimal-division/submissions/474107247/)                     | 除法 = 左区间 / 右区间         | [我的提交](https://leetcode.cn/problems/optimal-division/submissions/474107247/) |
| [516. 最长回文子序列](https://leetcode.cn/problems/longest-palindromic-subsequence/description/)                | 当前区间 = 左端点 + 子区间 + 右端点 | [我的提交](https://leetcode.cn/problems/longest-palindromic-subsequence/submissions/474136039/) |
| [1039. 多边形三角剖分的最低得分](https://leetcode.cn/problems/minimum-score-triangulation-of-polygon/description/) | 当前区间 = 左区间 + 三角形 + 右区间 | [我的提交](https://leetcode.cn/problems/minimum-score-triangulation-of-polygon/submissions/477127316/) |
| [1130. 叶值的最小代价生成树](https://leetcode.cn/problems/minimum-cost-tree-from-leaf-values/description/) | 当前区间最优 = min(左区间最优 + 右区间最优 + 左区间最大值*右区间最大值) | [我的提交](https://leetcode.cn/problems/minimum-cost-tree-from-leaf-values/submissions/477271871/) |

> 6> 根据题意使用多维动态规划
>
| 题目          | 说明                        | 实现                                                                            |
|-------------|---------------------------|-------------------------------------------------------------------------------|
| [873. 最长的斐波那契子序列的长度](https://leetcode.cn/problems/length-of-longest-fibonacci-subsequence/description/) | 使用二维状态表示以i和j位置结尾的最长长度     | [我的提交](https://leetcode.cn/problems/length-of-longest-fibonacci-subsequence/submissions/475668428/) |
| [1027. 最长等差数列](https://leetcode.cn/problems/longest-arithmetic-subsequence/description/) | 使用二维状态表示以i和j位置结尾的最长长度     | [我的提交](https://leetcode.cn/problems/longest-arithmetic-subsequence/submissions/477005039/) |
| [87. 扰乱字符串](https://leetcode.cn/problems/scramble-string/description/) | 可以分解成子问题，s1、s2位置、长度三维动态规划 | [我的提交](https://leetcode.cn/problems/scramble-string/submissions/485794249/) |

> 7> 求概率问题从小规模下的概率递推到大规模条件的概率
>
| 题目                                                                                           | 说明                                 | 实现                                                                            |
|----------------------------------------------------------------------------------------------|------------------------------------|-------------------------------------------------------------------------------|
| [808. 分汤](https://leetcode.cn/problems/soup-servings/description/) | 知道A=0&B=0的概率，A=0&B>0的概率，递推至A=N,B=N | [我的提交](https://leetcode.cn/problems/soup-servings/submissions/475925300/) |

> 8> 博弈游戏类型（当前用户最优选择 = max(当前选择 + 当前选择导致的状态下当前用户（即另一用户）最优选择)）
>
| 题目                                                                   | 说明                                        | 实现                                                                            |
|----------------------------------------------------------------------|-------------------------------------------|-------------------------------------------------------------------------------|
| [486. 预测赢家](https://leetcode.cn/problems/predict-the-winner/description/) | 抽象成当前用户和下一用户之间的状态转移，利用区间类型动态规划            | [我的提交](https://leetcode.cn/problems/predict-the-winner/submissions/476010042/) |
| [464. 我能赢吗](https://leetcode.cn/problems/can-i-win/description/) | 类似的思路，使用记忆化搜索的方法                          | [我的提交](https://leetcode.cn/problems/can-i-win/submissions/466755816/) |
| [1140. 石子游戏 II](https://leetcode.cn/problems/stone-game-ii/description/) | dp(i)(j) 表示玩家在stones(i:)开始选择，M=j时能获得的最多分数 | [我的提交](https://leetcode.cn/problems/stone-game-ii/submissions/477461605/) |

> 9> 当前状态可能由前面一个或多个特定状态转移得到，根据题目条件分析
>
| 题目                                                                                          | 说明                                                                                 | 实现                                                                            |
|---------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------|-------------------------------------------------------------------------------|
| [823. 带因子的二叉树](https://leetcode.cn/problems/binary-trees-with-factors/description/)         | 排序数组，依次求每个节点为树根的情况                                                                 | [我的提交](https://leetcode.cn/problems/binary-trees-with-factors/submissions/476022723/) |
| [907. 子数组的最小值之和](https://leetcode.cn/problems/sum-of-subarray-minimums/description/)        | i为右端点的所有子数组最小值之和 = 上一个更小元素(位置j)为右...之和 + arr[i]*(i-j)                              | [我的提交](https://leetcode.cn/problems/sum-of-subarray-minimums/submissions/476941357/) |
| [983. 最低票价](https://leetcode.cn/problems/minimum-cost-for-tickets/description/)             | 第i天最低消费 = min( (1天票+前i-1天最低消费),(7天票+前i-7天最低消费),(30天票+前i-30天最低消费) )                 | [我的提交](https://leetcode.cn/problems/minimum-cost-for-tickets/submissions/476968274/) |
| [1024. 视频拼接](https://leetcode.cn/problems/video-stitching/description/)                     | 以每个片段结尾的消耗的片段 = 能和当前片段开头拼接上的所有片段结尾的片段数+1 (此题，易错点较多)                                | [我的提交](https://leetcode.cn/problems/video-stitching/submissions/476996686/) |
| [1048. 最长字符串链](https://leetcode.cn/problems/longest-string-chain/description/)              | 先排序， 当前最长链 = max(删除每个字母后的新word在前面的链长 + 1)                                          | [我的提交](https://leetcode.cn/problems/longest-string-chain/submissions/477134593/) |
| [1105. 填充书架](https://leetcode.cn/problems/filling-bookcase-shelves/description/)            | 当前本之前的最优 = min(从当前本往前k本放在一行 + i-k之前的最优 ) (k=1,2...)                                | [我的提交](https://leetcode.cn/problems/filling-bookcase-shelves/submissions/477244807/) |
| [1578. 使绳子变成彩色的最短时间](https://leetcode.cn/problems/minimum-time-to-make-rope-colorful/description/) | dp[i] = min(移除i + dp[i-1], 移除上一与i不同的j处的dp[j] + 移除前（i-j-1）的cost)                    | [我的提交](https://leetcode.cn/problems/minimum-time-to-make-rope-colorful/submissions/477785601/) |
| [1621. 大小为 K 的不重叠线段的数目](https://leetcode.cn/problems/number-of-sets-of-k-non-overlapping-line-segments/description/) | dp(i)(k) 表示i点前 分配 k段绳子的方法数                                                         | [我的提交](https://leetcode.cn/problems/number-of-sets-of-k-non-overlapping-line-segments/submissions/477865061/) |
| [1626. 无矛盾的最佳球队](https://leetcode.cn/problems/best-team-with-no-conflicts/description/) | 对队伍双重排序（相同年龄按分数排），大大简化算法                                                           | [我的提交](https://leetcode.cn/problems/best-team-with-no-conflicts/submissions/477936708/) |
| [1696. 跳跃游戏 VI](https://leetcode.cn/problems/jump-game-vi/description/)    | 使用大顶堆优化对前面k个状态最优的查找                                                                | [我的提交](https://leetcode.cn/problems/jump-game-vi/submissions/477991644/) |
| [1884. 鸡蛋掉落-两枚鸡蛋](https://leetcode.cn/problems/egg-drop-with-2-eggs-and-n-floors/description/)    | i楼时往第j层扔一个鸡蛋，碎了则N=1+(i-1), 没碎则N=[j+1:i]区间所需次数(即dp[0:i-j]) + 1，取二者较大值（从j=1开始不断增加即可） | [我的提交](https://leetcode.cn/problems/egg-drop-with-2-eggs-and-n-floors/submissions/478151431/) |
| [1959. K 次调整数组大小浪费的最小总空间](https://leetcode.cn/problems/minimum-total-space-wasted-with-k-resizing-operations/description/)    | 划分成2段，总段在k情况最少使用空间 = 后一段的最大值 + 前一段在k-1情况下的最少使用空间 （后一段长度从1逐渐增大）                     | [我的提交](https://leetcode.cn/problems/minimum-total-space-wasted-with-k-resizing-operations/submissions/478450968/) |

> 10> 我称之为状态扩散，从某个状态按规则扩散到其余新状态
>
| 题目          | 说明                                                          | 实现                                                                            |
|-------------|-------------------------------------------------------------|-------------------------------------------------------------------------------|
| [1162. 地图分析](https://leetcode.cn/problems/as-far-from-land-as-possible/description/) | 某位置距离陆地d，则其上下左右的海洋距离陆地d+1，扩散至没有海洋                           | [我的提交](https://leetcode.cn/problems/as-far-from-land-as-possible/submissions/477471003/) |
| [1947. 最大兼容性评分和](https://leetcode.cn/problems/maximum-compatibility-score-sum/description/) | mark第j位表示第j个教授是否被选择，mark中有k个1，初始时mark=0，从每个mark扩散到含k+1的mark | [我的提交](https://leetcode.cn/problems/maximum-compatibility-score-sum/submissions/478437655/) |

> 11> 节点距离问题，使用动态规划、贝尔曼福特算法    
> 每轮循环知道相邻 f -> t 之间的距离c，对其余节点i(i!=f && i!=t):    
> if c + d(t)(i) < d(f)(i) 则更新d(f)(i)    
> !!! 不要忘了还有: **c + d(f)(i) < d(t)(i) 则更新d(t)(i)**    
>
| 题目          | 说明          | 实现                                                                            |
|-------------|-------------|-------------------------------------------------------------------------------|
| [1334. 阈值距离内邻居最少的城市](https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/description/) | 求出每个节点间最小距离 | [我的提交](https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/submissions/477690027/) |
| [1311. 获取你好友已观看的视频](https://leetcode.cn/problems/get-watched-videos-by-your-friends/description/) | 求出每个好友间最小距离 | [我的提交](https://leetcode.cn/problems/get-watched-videos-by-your-friends/submissions/479635030/) |
|             |             |                                                                               |
> 12> 状态压缩dp，通常用于数组A和数组B任意匹配问题
>
| 题目          | 说明                                           | 实现                                                                            |
|-------------|----------------------------------------------|-------------------------------------------------------------------------------|
| [1947. 最大兼容性评分和](https://leetcode.cn/problems/maximum-compatibility-score-sum/description/) | mark第j位表示第j个教授是否被选择，mark中1的数量k表示现在为第k个学生匹配教授 | [我的提交](https://leetcode.cn/problems/maximum-compatibility-score-sum/submissions/478437655/) |


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

