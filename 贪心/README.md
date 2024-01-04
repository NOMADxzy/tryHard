## 贪心大法

### 1. 概念
顾名思义，贪心算法总是作出在当前看来最好的选择。也就是说贪心算法并不从整体最优考虑，它所作出的选择只是在某种意义上的局部最优选择。
当然，希望贪心算法得到的最终结果也是整体最优的。虽然贪心算法不能对所有问题都得到整体最优解，但对许多问题它能产生整体最优解。
如单源最短路经问题，最小生成树问题等。在一些情况下，即使贪心算法不能得到整体最优解，其最终结果却是最优解的很好近似。


#### 基本思路：
1. 建立数学模型来描述问题。
2. 把求解的问题分成若干个子问题。
3. 对每一子问题求解，得到子问题的局部最优解。
4. 把子问题的解局部最优解合成原来解问题的一个解。
      


#### 实现该算法的过程：
1. 从问题的某一初始解出发;
2. while 能朝给定总目标前进一步 do
3. 求出可行解的一个解元素;
4. 由所有解元素组合成问题的一个可行解。从问题的某一初始解出发

### 2. 解题技巧（我的总结）

> 1> 最需要原则，元素选择问题，贪婪的选择某个最需要被选择的元素
> 
| 题目                                                                                 | 说明                                                 | 实现                                                                           |
|------------------------------------------------------------------------------------|----------------------------------------------------|------------------------------------------------------------------------------|
| [621. 任务调度器](https://leetcode.cn/problems/task-scheduler/description/)             | 当前剩下最多的任务最需要被调度                                    | [我的提交](https://leetcode.cn/problems/task-scheduler/submissions/478746885/) |
| [670. 最大交换](https://leetcode.cn/problems/maximum-swap/description/)                | 最靠左的元素最需要被增大，故从左到右遍历到i位置，i右边最大的位置大于i则结果为交换i和该最大的位置 | [我的提交](https://leetcode.cn/problems/maximum-swap/submissions/471160121/) |
| [678. 有效的括号字符串](https://leetcode.cn/problems/valid-parenthesis-string/description/) | 左括号最需要被消除，两个栈存放左括号和星号，每次遍历到右括号，优先消除左括号，没有左括号时再消除星号 | [我的提交](https://leetcode.cn/problems/valid-parenthesis-string/submissions/471177162/) |
| [738. 单调递增的数字](https://leetcode.cn/problems/monotone-increasing-digits/)           | 左边元素最需要选大的，从左往右遍历到i位置，i后全变为9，i前变为单增的最大形式           | [我的提交](https://leetcode.cn/problems/valid-parenthesis-string/submissions/471177162/) |
| [767. 重构字符串](https://leetcode.cn/problems/reorganize-string/description/)          | 记录当前所有字母的数量，最多的元素（不和上一个相同）是最需要被选择的元素               | [我的提交](https://leetcode.cn/problems/reorganize-string/submissions/471209764/) |
| [870. 优势洗牌](https://leetcode.cn/problems/advantage-shuffle/description/)           | nums1中最小的元素最需要超越对方的数，双指针                           | [我的提交](https://leetcode.cn/problems/advantage-shuffle/submissions/471365063/) |
| [881. 救生艇](https://leetcode.cn/problems/boats-to-save-people/description/)         | 体重最重的人最需要充分利用船                                     | [我的提交](https://leetcode.cn/problems/boats-to-save-people/submissions/471367754/) |
| [630. 课程表 III](https://leetcode.cn/problems/course-schedule-iii/description/)      | 时间越短的课程越优先，按照结束时间从小到大排序并遍历                         | [我的提交](https://leetcode.cn/problems/course-schedule-iii/submissions/485409704/) |
| [407. 接雨水 II](https://leetcode.cn/problems/trapping-rain-water-ii/description/)    | 木桶效应，先接最短边周围的水，再更新木桶边                              | [我的提交](https://leetcode.cn/problems/trapping-rain-water-ii/submissions/488936221/) |
| [948. 令牌放置](https://leetcode.cn/problems/bag-of-tokens/description/)              | 贪心的从小到大翻开token，从大到小闭合token                         | [我的提交](https://leetcode.cn/problems/bag-of-tokens/submissions/490914671/) |
| [1111. 有效括号的嵌套深度](https://leetcode.cn/problems/bag-of-tokens/description/     | 贪心的将左括号分配给累积左括号更小的一边                               | [我的提交](https://leetcode.cn/problems/maximum-nesting-depth-of-two-valid-parentheses-strings/submissions/492209285/) |

> 2> 递归回溯，每次做局部最优选择，注意使用状态压缩防止重复搜索
>
| 题目                                                              | 说明             | 实现                                                                            |
|-----------------------------------------------------------------|----------------|-------------------------------------------------------------------------------|
| [55. 跳跃游戏](https://leetcode.cn/problems/jump-game/description/) | 当前跨步越大越优，按顺序尝试 | [我的提交](https://leetcode.cn/problems/jump-game/submissions/460023694/) |

> 3> 问题分解子问题，贪心的求解每个子问题
>
| 题目                                                             | 说明                            | 实现                                                                            |
|----------------------------------------------------------------|-------------------------------|-------------------------------------------------------------------------------|
| [135. 分发糖果](https://leetcode.cn/problems/candy/description/) | 从左到右/从右到左 分别给每个需要增加糖果的增加到恰好可以 | [我的提交](https://leetcode.cn/problems/candy/submissions/487449654/) |


> 4> 弄清什么是局部，贪心优化，排序等方法确定具体怎么取局部
>
| 题目                                                              | 说明                    | 实现                                                                            |
|-----------------------------------------------------------------|-----------------------|-------------------------------------------------------------------------------|
| [781. 森林中的兔子](https://leetcode.cn/problems/rabbits-in-forest/description/) | 排序，只有答案相同的兔子才有可能是相同颜色 | [我的提交](https://leetcode.cn/problems/rabbits-in-forest/submissions/471254724/) |

> 5> 利用数据结构（小顶堆，哈希表）维护贪心所需的信息
>
| 题目                                                              | 说明                        | 实现                                                                            |
|-----------------------------------------------------------------|---------------------------|-------------------------------------------------------------------------------|
| [659. 分割数组为连续子序列](https://leetcode.cn/problems/split-array-into-consecutive-subsequences/description/) | 哈希表、小顶堆，每个元素只分配给长度最小的累积序列 | [我的提交](https://leetcode.cn/problems/split-array-into-consecutive-subsequences/submissions/490113579/) |

> 6> 贪心构造
>
| 题目                                                              | 说明                                 | 实现                                                                            |
|-----------------------------------------------------------------|------------------------------------|-------------------------------------------------------------------------------|
| [1405. 最长快乐字符串](https://leetcode.cn/problems/longest-happy-string/description/) | 贪心规则：最多的字符最需要被取，字符数量超过其余的取2个，否则取1个 | [我的提交](https://leetcode.cn/problems/longest-happy-string/submissions/493194366/) |


### 3. 更多练习


### 4. 参考
1. [贪心算法详解 ](https://mp.weixin.qq.com/s?__biz=MzU1NjEwMTY0Mw==&mid=2247551986&idx=1&sn=cf291ece55b4f6ef650e13f191cc189b&chksm=fbc87296ccbffb8068158832461e9dab75ee8f7b3a88aa015e709d593ca110ef135e456135d4&scene=27) 