## 深度优先算法

### 1. 概念


```markdown
1、对每一个可能的分支路径深入到不能再深入为止，而且每个结点只能访问一次。要特别注意的是，二叉树的深度优先遍历比较特殊，可以细分为先序遍历、中序遍历、后序遍历（我们前面使用的是先序遍历）。
2、不全部保留结点，占用空间少；有回溯操作(即有入栈、出栈操作)，运行速度慢。
3、一般无回溯操作，即入栈和出栈的操作，所以运行速度比深度优先搜索要快些
```


### 2. 解题技巧（我的总结）

> 1> 数位DP，求小于N的符合条件的数的个数，注意以下几点：
> 用一个 target数组 存储N的每一位
> 用一个 preLimit 变量表示前面的是否处于临界状态，当且仅当preLimit为true且当前位到达target对应位时，nextLimit为true
> 用一个 pos 表示当前处理到数字的第多少位，pos到len(target)时结束递归
> 
> 
| 题目                                                                   | 说明                                  | 实现                                                                            |
|----------------------------------------------------------------------|-------------------------------------|-------------------------------------------------------------------------------|
| [788. 旋转数字](https://leetcode.cn/problems/rotated-digits/description/) | dp时用preValid表示前面是否已经含有2569          | [我的提交](https://leetcode.cn/problems/rotated-digits/description/) |
| [2376. 统计特殊整数](https://leetcode.cn/problems/count-special-integers/description/) | 使用一个mark，第i位标记i是否被使用，从而压缩状态         | [我的提交](https://leetcode.cn/problems/count-special-integers/submissions/473515507/) |
| [60. 排列序列](https://leetcode.cn/problems/permutation-sequence/description/) | 从左到右递归累计排名                          | [我的提交](https://leetcode.cn/problems/permutation-sequence/submissions/485525438/) |
| [233. 数字 1 的个数](https://leetcode.cn/problems/number-of-digit-one/description/) | 使用 +(-)pos*10 + pre1sum 作为key存储中间结果 | [我的提交](https://leetcode.cn/problems/number-of-digit-one/submissions/487833449/) |
| [600. 不含连续1的非负整数](https://leetcode.cn/problems/non-negative-integers-without-consecutive-ones/description/) | 只要不限制数字的长度 就可以无视前导0 直接dfs           | [我的提交](https://leetcode.cn/problems/non-negative-integers-without-consecutive-ones/submissions/499424741/) |
| [3129. 找出所有稳定的二进制数组 I](https://leetcode.cn/problems/find-all-possible-stable-binary-arrays-i/description/) | dfs, 累计1，累计0的个数，剩余1，剩余0的个数          | [我的提交](https://leetcode.cn/problems/find-all-possible-stable-binary-arrays-i/submissions/554207032/) |

> 2> 前序、中序、后序遍历，二叉搜索树
>
| 题目                                                                   | 说明                            | 实现                                                                            |
|----------------------------------------------------------------------|-------------------------------|-------------------------------------------------------------------------------|
| [2476. 二叉搜索树最近节点查询](https://leetcode.cn/problems/closest-nodes-queries-in-a-binary-search-tree/description/) | 中序遍历 + 排序 + 二分查找              | [我的提交](https://leetcode.cn/problems/closest-nodes-queries-in-a-binary-search-tree/submissions/486976544/) |
| [538. 把二叉搜索树转换为累加树](https://leetcode.cn/problems/convert-bst-to-greater-tree/description/) | 先右后左遍历二叉搜索树                   | [我的提交](https://leetcode.cn/problems/convert-bst-to-greater-tree/submissions/489578707/) |
| [958. 二叉树的完全性检验](https://leetcode.cn/problems/check-completeness-of-a-binary-tree/description/) | dfs(root)返回三个值，深度、是否完美、是否完全   | [我的提交](https://leetcode.cn/problems/check-completeness-of-a-binary-tree/submissions/490780299/) |
| [987. 二叉树的垂序遍历](https://leetcode.cn/problems/vertical-order-traversal-of-a-binary-tree/description/) | 先dfs一边获取矩阵的宽高，再dfs向对应位置添加元素   | [我的提交](https://leetcode.cn/problems/vertical-order-traversal-of-a-binary-tree/submissions/491392426/) |
| [1373. 二叉搜索子树的最大键值和](https://leetcode.cn/problems/maximum-sum-bst-in-binary-tree/description/) | 返回4个值：树是否合法，树最大值，树最小值，元素和     | [我的提交](https://leetcode.cn/problems/maximum-sum-bst-in-binary-tree/submissions/493000220/) |
| [1519. 子树中标签相同的节点数](https://leetcode.cn/problems/number-of-nodes-in-the-sub-tree-with-the-same-label/description/) | 返回1个数组：记录26个字母的出现次数           | [我的提交](https://leetcode.cn/problems/number-of-nodes-in-the-sub-tree-with-the-same-label/submissions/493666630/) |
| [1609. 奇偶树](https://leetcode.cn/problems/even-odd-tree/description/) | 记录每层的最近值，dfs可保证每层都是从左到右顺序访问到的 | [我的提交](https://leetcode.cn/problems/even-odd-tree/submissions/493981516/) |

> 3> 递归分解问题
>
| 题目                                                                                                               | 说明                                                   | 实现                                                                                   |
|------------------------------------------------------------------------------------------------------------------|------------------------------------------------------|--------------------------------------------------------------------------------------|
| [282. 给表达式添加运算符](https://leetcode.cn/problems/expression-add-operators/description/)                             | 分解当前val +- 后续val = target，val再分解为含有多少个*              | [我的提交](https://leetcode.cn/problems/expression-add-operators/submissions/487984888/) |
| [224. 基本计算器](https://leetcode.cn/problems/basic-calculator/description/)                                         | 用栈预处理获取每个左括号对应右位置，然后 每轮读取 一符号+一数值                    | [我的提交](https://leetcode.cn/problems/basic-calculator/submissions/499661657/)         |
| [761. 特殊的二进制序列](https://leetcode.cn/problems/special-binary-string/description/)                                 | 特殊串一定是1***0这种形式，相邻交换即冒泡排序，且特殊串的所有特殊字串一定是相邻的          | [我的提交](https://leetcode.cn/problems/special-binary-string/submissions/501935866/)    |
| [LCR 087. 复原 IP 地址](https://leetcode.cn/problems/0on3uN/description/)                                            | 从i开始的后缀 分成 k个数 的所有可能字符串，递归分解                         | [我的提交](https://leetcode.cn/problems/0on3uN/submissions/517010566/)                   |
| [2192. 有向无环图中一个节点的所有祖先](https://leetcode.cn/problems/all-ancestors-of-a-node-in-a-directed-acyclic-graph/description/) | pres数组记录所有前驱节点，dfs返回i的祖先，节点i的祖先 = dfs(i前驱节点) + i前驱节点 | [我的提交](https://leetcode.cn/problems/all-ancestors-of-a-node-in-a-directed-acyclic-graph/submissions/531910708/)  |
| [LCR 119. 最长连续序列](https://leetcode.cn/problems/WhsWhI/description/?envType=study-plan-v2&envId=coding-interviews-special) | 哈希表记录出现与否，dfs(i)=dfs(i-1)+dfs(i+1)+1                 | [我的提交](https://leetcode.cn/problems/WhsWhI/submissions/536505147/?envType=study-plan-v2&envId=coding-interviews-special)  |
| [LCR 050. 路径总和 III](https://leetcode.cn/problems/6eUYwP/description/)        | dfs, 记录每个节点为起点的路径为各个值的数量，递归                          | [我的提交](https://leetcode.cn/problems/6eUYwP/submissions/538098526/)  |
| [LCR 113. 课程表 II](https://leetcode.cn/problems/QA2IGt/description/)        | dfs, 尝试访问每个节点的后继，入度为0代表可以访问                          | [我的提交](https://leetcode.cn/problems/QA2IGt/submissions/539867549/)  |
| [2673. 使二叉树所有路径值相等的最小代价](https://leetcode.cn/problems/make-costs-of-paths-equal-in-a-binary-tree/description/)        | dfs, 记录左右子树下的总值，差值就是左右节点要调整的值，递归向上传递                 | [我的提交](https://leetcode.cn/problems/make-costs-of-paths-equal-in-a-binary-tree/submissions/551691624/)  |
| [3040. 相同分数的最大操作数目 II](https://leetcode.cn/problems/maximum-number-of-operations-with-the-same-score-ii/description/)        | dfs, (起点,终点,目标值)，递归从外向里搜寻可进行的步骤数，取最大的                | [我的提交](https://leetcode.cn/problems/maximum-number-of-operations-with-the-same-score-ii/submissions/552367309/)  |

> 4> 去重
>
| 题目                                                                   | 说明                                        | 实现                                                                            |
|----------------------------------------------------------------------|-------------------------------------------|-------------------------------------------------------------------------------|
| [491. 递增子序列](https://leetcode.cn/problems/non-decreasing-subsequences/description/) | 某一位置累积序列acc，遍历后面每个比当前大的元素，使用map防止遍历相同的val | [我的提交](https://leetcode.cn/problems/non-decreasing-subsequences/submissions/489270807/) |

> 5> 反证法递归（假设条件）
>
| 题目                                                                   | 说明              | 实现                                                                            |
|----------------------------------------------------------------------|-----------------|-------------------------------------------------------------------------------|
| [785. 判断二分图](https://leetcode.cn/problems/is-graph-bipartite/description/) | 假设是二分图，按照定义探索矛盾 | [我的提交](https://leetcode.cn/problems/is-graph-bipartite/submissions/490514280/) |

> 6> 搜索所有答案（求最优值）剪枝
>
| 题目                                                                   | 说明                                                                  | 实现                                                                            |
|----------------------------------------------------------------------|---------------------------------------------------------------------|-------------------------------------------------------------------------------|
| [421. 数组中两个数的最大异或值](https://leetcode.cn/problems/maximum-xor-of-two-numbers-in-an-array/description/) | 根据二进制建立二叉树                                                          | [我的提交](https://leetcode.cn/problems/maximum-xor-of-two-numbers-in-an-array/submissions/469824077/) |
| [789. 逃脱阻碍者](https://leetcode.cn/problems/escape-the-ghosts/description/) | 想像成阻碍者能否在吃豆人到达某一位置前到达某一位置，能则剪枝                                      | [我的提交](https://leetcode.cn/problems/escape-the-ghosts/submissions/490546599/) |
| [1593. 拆分字符串使唯一子字符串的数目最大](https://leetcode.cn/problems/split-a-string-into-the-max-number-of-unique-substrings/description/) | 剩余能达到的最大数量小于已有最优数量，则剪枝                                              | [我的提交](https://leetcode.cn/problems/split-a-string-into-the-max-number-of-unique-substrings/submissions/493840857/) |
| [1901. 寻找峰值 II](https://leetcode.cn/problems/find-a-peak-element-ii/description/) | 每次都向周围4个中最大的移动                                                      | [我的提交](https://leetcode.cn/problems/find-a-peak-element-ii/submissions/495909101/) |
| [691. 贴纸拼词](https://leetcode.cn/problems/stickers-to-spell-word/description/) | state记录每个字母剩余个数，选择一个可以使state减小的word，dfs(state) = 1 + dfs(nextState) | [我的提交](https://leetcode.cn/problems/stickers-to-spell-word/submissions/499690857/) |
| [514. 自由之路](https://leetcode.cn/problems/freedom-trail/description/) | 缓存 （ring位置，key位置）进行状态压缩，预处理记录每个字符出现的位置列表                            | [我的提交](https://leetcode.cn/problems/freedom-trail/submissions/499875669/) |
| [LCR 085. 括号生成](https://leetcode.cn/problems/IDBivT/description/) | 记录左括号-右括号的数量，每个位置尝试使用'('和')'，递归到后一个位置      | [我的提交](https://leetcode.cn/problems/IDBivT/submissions/536327462/) |


> 7> 预处理条件，简化递归过程
>
| 题目                                                                   | 说明                               | 实现                                                                            |
|----------------------------------------------------------------------|----------------------------------|-------------------------------------------------------------------------------|
| [959. 由斜杠划分区域](https://leetcode.cn/problems/regions-cut-by-slashes/description/) | 每个方格分成4个部分，用map存储不同条件下的相邻区域，简化过程 | [我的提交](https://leetcode.cn/problems/regions-cut-by-slashes/submissions/490966770/) |

> 8> 深度优先按序搜索，查找第k个答案
>
| 题目                                                                   | 说明                                      | 实现                                                                            |
|----------------------------------------------------------------------|-----------------------------------------|-------------------------------------------------------------------------------|
| [1415. 长度为 n 的开心字符串中字典序第 k 小的字符串](https://leetcode.cn/problems/the-k-th-lexicographical-string-of-all-happy-strings-of-length-n/description/) | 深度优先，按字典序遍历，记录答案排名，到第k个时输出              | [我的提交](https://leetcode.cn/problems/the-k-th-lexicographical-string-of-all-happy-strings-of-length-n/submissions/493244190/) |
| [1718. 构建字典序最大的可行序列](https://leetcode.cn/problems/construct-the-lexicographically-largest-valid-sequence/description/) | 深度优先，c从大到小填值，i没使用且，pos和pos+i位置都是空的则可以填入 | [我的提交](https://leetcode.cn/problems/construct-the-lexicographically-largest-valid-sequence/submissions/494480407/) |
| [2182. 构造限制重复的字符串](https://leetcode.cn/problems/construct-string-with-repeat-limit/description/) | 深度优先，优先使用最大的字符，直至没有字符能够使用               | [我的提交](https://leetcode.cn/problems/construct-string-with-repeat-limit/submissions/499005609/) |


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
3. 总库：[tryHard](https://github.com/NOMADxzy/tryHard)