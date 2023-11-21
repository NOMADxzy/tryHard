## 递归回溯

### 1. 概念

站在回溯树的一个节点上，你只需要思考 3 个问题：

1、路径：也就是已经做出的选择。

2、选择列表：也就是你当前可以做的选择。

3、结束条件：也就是到达决策树底层，无法再做选择的条件

```go
result = []
def backtrack(路径, 选择列表):
    if 满足结束条件:
        result.add(路径)
        return
    
    for 选择 in 选择列表:
        做选择
        backtrack(路径, 选择列表)
        撤销选择

```

> 典型的 DFS 但是没有回溯的题目：[17. 电话号码的字母组合](https://leetcode.cn/problems/letter-combinations-of-a-phone-number/)



### 2. 解题技巧（我的总结）

> 1> 使用哈希表存储答案，使用mark标记已访问的防止重复搜索，状态较少（<32）时用一个int即可
> 
| 题目                                                          | 说明                     | 实现                                                                            |
|-------------------------------------------------------------|------------------------|-------------------------------------------------------------------------------|
| [464. 我能赢吗](https://leetcode.cn/problems/can-i-win/description/) | 使用int记录已选数字            | [我的提交](https://leetcode.cn/problems/can-i-win/submissions/466755816/) |
| [1387. 将整数按权重排序](https://leetcode.cn/problems/sort-integers-by-the-power-value/description/) | 使用int记录已计算过的           | [我的提交](https://leetcode.cn/problems/sort-integers-by-the-power-value/submissions/477710824/) |
| [1361. 验证二叉树](https://leetcode.cn/problems/validate-binary-tree-nodes/description/) | 记录每个节点为根时其下的节点数目，直到等于n | [我的提交](https://leetcode.cn/problems/validate-binary-tree-nodes/submissions/483028473/) |

> 2> 岛屿问题：曼哈顿距离，广度优先搜索，某一性质（边界）扩散，通过更改标记防止重复搜索
>
| 题目                                                                   | 说明                                                     | 实现                                                                            |
|----------------------------------------------------------------------|--------------------------------------------------------|-------------------------------------------------------------------------------|
| [934. 最短的桥](https://leetcode.cn/problems/shortest-bridge/description/) | 两岛屿使用不同的标记，bfs扩散岛屿1的面积直至和岛屿2相遇                         | [我的提交](https://leetcode.cn/problems/shortest-bridge/submissions/482164939/) |
| [1020. 飞地的数量](https://leetcode.cn/problems/number-of-enclaves/description/) | 从边界扩散，剩余的是有效的1                                         | [我的提交](https://leetcode.cn/problems/number-of-enclaves/submissions/482424903/) |
| [1034. 边界着色](https://leetcode.cn/problems/coloring-a-border/description/) | dfs，是边界标记为-1，不是标记为0，（不是边界 ：4个方向的val=target/0/-1的数量==4） | [我的提交](https://leetcode.cn/problems/coloring-a-border/submissions/482530829/) |
| [1254. 统计封闭岛屿的数目](https://leetcode.cn/problems/number-of-closed-islands/description/) | 边缘的0开始扩散并标志为1，接着深度优先搜索所有独立岛屿                           | [我的提交](https://leetcode.cn/problems/number-of-closed-islands/submissions/483169743/) |

> 3> 二叉树技巧
> 只考虑根节点 (先考虑根节点为nil，再考虑和根节点的关系，递归到下一层) 
> 使用map记录父节点
>
| 题目                                                                               | 说明                          | 实现                                                                            |
|----------------------------------------------------------------------------------|-----------------------------|-------------------------------------------------------------------------------|
| [669. 修剪二叉搜索树](https://leetcode.cn/problems/trim-a-binary-search-tree/description/) | 根节点不在区间内/在区间内两种情况           | [我的提交](https://leetcode.cn/problems/trim-a-binary-search-tree/submissions/481581106/) |
| [623. 在二叉树中增加一行](https://leetcode.cn/problems/add-one-row-to-tree/description/) | 根节点在新层 前两层以上/前一层/新层 三种情况    | [我的提交](https://leetcode.cn/problems/add-one-row-to-tree/submissions/481586093/) |
| [863. 二叉树中所有距离为 K 的结点](https://leetcode.cn/problems/all-nodes-distance-k-in-binary-tree/description/) | 记录父节点，考虑下面的后再向上考虑每个父节点      | [我的提交](https://leetcode.cn/problems/all-nodes-distance-k-in-binary-tree/submissions/482123513/) |
| [971. 翻转二叉树以匹配先序遍历](https://leetcode.cn/problems/add-one-row-to-tree/description/) | 预处理统计每个节点下的节点数目，从而对数组划分     | [我的提交](https://leetcode.cn/problems/add-one-row-to-tree/submissions/481586093/) |
| [652. 寻找重复的子树](https://leetcode.cn/problems/find-duplicate-subtrees/description/) | 二叉树的序列化                     | [我的提交](https://leetcode.cn/problems/find-duplicate-subtrees/submissions/483213917/) |
| [1443. 收集树上所有苹果的最少时间](https://leetcode.cn/problems/minimum-time-to-collect-all-apples-in-a-tree/description/) | 当前节点time = sum(每个子节点time+2) | [我的提交](https://leetcode.cn/problems/minimum-time-to-collect-all-apples-in-a-tree/submissions/483523574/) |

> 4> 棋盘问题，三步:
> 递归终止，存储答案
> 遍历当前位置所有合法解，进入下一位置
> 当前位置无合法解，回退，撤销一切更改
>
| 题目                                                                    | 说明                   | 实现                                                                            |
|-----------------------------------------------------------------------|----------------------|-------------------------------------------------------------------------------|
| [51. N 皇后](https://leetcode.cn/problems/n-queens/description/) | 从上往下为每一行选择Q位置        | [我的提交](https://leetcode.cn/problems/n-queens/submissions/481597366/) |
| [1391. 检查网格中是否存在有效路径](https://leetcode.cn/problems/check-if-there-is-a-valid-path-in-a-grid/description/) | 使用4维特征记录每种街道的上下左右可行性 | [我的提交](https://leetcode.cn/problems/check-if-there-is-a-valid-path-in-a-grid/submissions/483493723/) |
| [1905. 统计子岛屿](https://leetcode.cn/problems/count-sub-islands/description/) | ps (每访问一个节点就更改其值以去重)       | [我的提交](https://leetcode.cn/problems/count-sub-islands/submissions/483826848/) |

> 5> 降低搜索空间，从排序、忽略顺序、只考虑奇偶性、对实体分类等方面简化
>
| 题目                                                                    | 说明                     | 实现                                                                            |
|-----------------------------------------------------------------------|------------------------|-------------------------------------------------------------------------------|
| [672. 灯泡开关 Ⅱ](https://leetcode.cn/problems/bulb-switcher-ii/description/) | 只考虑每组灯泡被按了奇偶次，灯泡可以分成4组 | [我的提交](https://leetcode.cn/problems/bulb-switcher-ii/submissions/481694656/) |
| [473. 火柴拼正方形](https://leetcode.cn/problems/matchsticks-to-square/description/) | 降序排列降低后面的枚举次数 | [我的提交](https://leetcode.cn/problems/matchsticks-to-square/submissions/474082583/) |

> 6> 深度优先套路
> dfs(env,pos,mark):
>   if mark[pos] return
>   else{
>       所有合法的下一位置nextPos:
>           处理环境
>           dfs(env,nextPos,mark)
>           恢复环境
>   }
>
| 题目                                                                    | 说明       | 实现                                                                            |
|-----------------------------------------------------------------------|----------|-------------------------------------------------------------------------------|
| [841. 钥匙和房间](https://leetcode.cn/problems/keys-and-rooms/description/) | 深度优先每个房间 | [我的提交](https://leetcode.cn/problems/keys-and-rooms/submissions/482103164/) |

> 7> 预处理条件，使用map形式方便查找`下一合法位置`
>
| 题目                                                                    | 说明                  | 实现                                                                            |
|-----------------------------------------------------------------------|---------------------|-------------------------------------------------------------------------------|
| [851. 喧闹和富有](https://leetcode.cn/problems/loud-and-rich/description/) | map记录比每个用户有钱的直接用户   | [我的提交](https://leetcode.cn/problems/loud-and-rich/submissions/482111789/) |
| [756. 金字塔转换矩阵](https://leetcode.cn/problems/pyramid-transition-matrix/description/) | map记录每个二元组上面能放的合法字符 | [我的提交](https://leetcode.cn/problems/pyramid-transition-matrix/submissions/481909980/) |
| [1042. 不邻接植花](https://leetcode.cn/problems/flower-planting-with-no-adjacent/description/) | map记录比每个花园的邻居       | [我的提交](https://leetcode.cn/problems/flower-planting-with-no-adjacent/submissions/482632241/) |

> 8> 利用二叉树性质，转换、简化题目问题为可以递归的问题
> 
| 题目                                                                    | 说明                                    | 实现                                                                            |
|-----------------------------------------------------------------------|---------------------------------------|-------------------------------------------------------------------------------|
| [1530. 好叶子节点对的数量](https://leetcode.cn/problems/number-of-good-leaf-nodes-pairs/description/) | 题意->求解所有树枝节点的合法路径->递归记录每个节点到旗下叶子的路径长度 | [我的提交](https://leetcode.cn/problems/number-of-good-leaf-nodes-pairs/submissions/483683617/) |
| [1559. 二维网格图中探测环](https://leetcode.cn/problems/detect-cycles-in-2d-grid/description/) | dfs中两个分支之间不会产生相连->dfs找环               | [我的提交](https://leetcode.cn/problems/detect-cycles-in-2d-grid/submissions/483803230/) |


### 3. 更多练习
#### 3.1 子集型回溯

| 题目                                                         | 说明                                                         | 题解                                                         |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| [78. 子集](https://leetcode.cn/problems/subsets/)            | 和 LC.39 类似，按照 begin 为起点遍历回溯就可以               | [通过](https://leetcode.cn/submissions/detail/395238958/)    |
| [90. 子集 II](https://leetcode.cn/problems/subsets-ii/)      | 在 LC.78 的基础上有重复元素的考虑，和 LC.40 类似的剪枝 :fire: | [通过](https://leetcode.cn/submissions/detail/395250094/)    |
| [131. 分割回文串](https://leetcode.cn/problems/palindrome-partitioning/) | 枚举字符之间的逗号，按照 idx 顺序回溯，判断回文              | [通过](https://leetcode.cn/submissions/detail/395280098/)    |
| [698. 划分为k个相等的子集](https://leetcode.cn/problems/partition-to-k-equal-sum-subsets/) | **抽象成 k 个桶**，每个桶的容量为子集和大小                  | [通过](https://leetcode.cn/submissions/detail/396107337/)    |
| [473. 火柴拼正方形](https://leetcode.cn/problems/matchsticks-to-square/) | 和 LC.698 一模一样，**抽象成 4 个桶**                        | [通过](https://leetcode.cn/submissions/detail/366042332/)    |
| [2305. 公平分发饼干](https://leetcode.cn/problems/fair-distribution-of-cookies/) | k 个桶，但是桶大小未知，从大到小DFS回溯                      | [通过](https://leetcode.cn/submissions/detail/396110148/)    |
| [854. 相似度为 K 的字符串](https://leetcode.cn/problems/k-similar-strings/) | DFS 回溯，剪枝，有点难...                                    | [通过](https://leetcode.cn/submissions/detail/396114920/)    |
| [1255. 得分最高的单词集合](https://leetcode.cn/problems/maximum-score-words-formed-by-letters/) | 参考灵神的子集型回溯，选或不选的思想，注意恢复现场 :fire:    | [0x3F](https://leetcode.cn/problems/maximum-score-words-formed-by-letters/solution/hui-su-san-wen-si-kao-hui-su-wen-ti-de-t-kw3y/) |


#### 3.2 组合型回溯

| 题目                                                         | 说明                                                         | 题解                                                         |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| [39. 组合总和](https://leetcode.cn/problems/combination-sum/) | 组合问题需要按照某种顺序搜索：每一次搜索的时候设置 **下一轮搜索的起点** `begin`，也可以排序之后加速剪枝 | [通过](https://leetcode.cn/submissions/detail/171894367/)    |
| [40. 组合总和 II](https://leetcode.cn/problems/combination-sum-ii/) | 和 LC.39 区别是这个有重复元素，需要去掉当前层（for循环中）**第二个数值重复的节点** :fire: | [剪枝](https://leetcode.cn/problems/combination-sum-ii/solution/hui-su-suan-fa-jian-zhi-python-dai-ma-java-dai-m-3/225211) |
| [77. 组合](https://leetcode.cn/problems/combinations/)       | 和 LC.39 类似，按照 begin 为起点遍历然后回溯就可以，剪枝：剩余元素个数需要多余 k | [通过](https://leetcode.cn/submissions/detail/395236585/)    |
| [216. 组合总和 III](https://leetcode.cn/problems/combination-sum-iii/) | 和 LC.40 类似，这题没有重复元素，两个剪枝：**小于最小的 \|\| 大于最大的** :fire: | [0x3F](https://leetcode.cn/problems/combination-sum-iii/solution/hui-su-bu-hui-xie-tao-lu-zai-ci-pythonja-feme/) |
| [22. 括号生成](https://leetcode.cn/problems/generate-parentheses/) | 直接 DFS 思路（选或不选）更简单，比较剩余左右括号数量剪枝    | [LWW](https://leetcode.cn/problems/generate-parentheses/solution/hui-su-suan-fa-by-liweiwei1419/) |
    


#### 3.3 排列型回溯

| 题目                                                         | 说明                                                         | 题解                                                         |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| [46. 全排列](https://leetcode.cn/problems/permutations/)     | 回溯入门问题，无重复元素的排列问题                           | [通过](https://leetcode.cn/submissions/detail/395204984/)    |
| [47. 全排列 II](https://leetcode.cn/problems/permutations-ii/) | 去重是关键，排序比较，**上一个相同的元素撤销之后再剪枝** :fire: | [剪枝图](https://leetcode.cn/problems/permutations-ii/solution/hui-su-suan-fa-python-dai-ma-java-dai-ma-by-liwe-2/) |
| [51. N 皇后](https://leetcode.cn/problems/n-queens/)         | 第 row 行填入 path[row] 这个数，求 path 满足条件的全排列     | [Carl](https://programmercarl.com/0051.N%E7%9A%87%E5%90%8E.html) |


### 4. 参考


- [回溯算法入门级详解 + 练习（持续更新）](https://leetcode.cn/problems/permutations/solution/hui-su-suan-fa-python-dai-ma-java-dai-ma-by-liweiw/)
- [回溯算法套路①子集型回溯【基础算法精讲 14】](https://www.bilibili.com/video/BV1mG4y1A7Gu/?vd_source=286032bc2c5715c8b50b608028ce57df)
- [经典回溯算法：集合划分问题「重要更新 🔥🔥🔥」](https://leetcode.cn/link/?target=https://lfool.github.io/LFool-Notes/algorithm/经典回溯算法：集合划分问题.html)
