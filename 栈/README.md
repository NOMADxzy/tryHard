## 栈

### 1. 概念
栈本质是一个线性表，所以它既可以是顺序表也可以是链表，通常为顺序表。与寻常的线性表的不同之处是：栈只可以在端点存入与输出（端点：一般为线性表的尾部，也就是栈顶），不可以随意的
的中间插入输出、头部插入输出。因为这个要求，造就了栈“先进后出，后进先出”的特点。    

`单调栈: ` 从栈底元素到栈顶元素呈单调递增或单调递减，栈内序列满足单调性的栈。

### 2. 解题技巧（我的总结）

> 1> 单调栈，记录数组/区间中的极值, 单调递增栈维护最小值、次小值... 单调递减栈维护最大值、次大值...
> 
| 题目                                                                            | 说明                                         | 实现                                                                            |
|-------------------------------------------------------------------------------|--------------------------------------------|-------------------------------------------------------------------------------|
| [1438. 绝对差不超过限制的最长连续子数组](https://leetcode.cn/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/description/) | 用单增栈和单减栈分别维护区间的最大最小值（队首），区间不满足要求时尝试从队首移除元素 | [我的提交](https://leetcode.cn/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/submissions/480918945/) |
| [239. 滑动窗口最大值](https://leetcode.cn/problems/sliding-window-maximum/description/) | 用单减栈维护窗口的最大值            | [我的提交](https://leetcode.cn/problems/sliding-window-maximum/submissions/480967154/) |

> 2> 单调栈，隐含历史信息
>
| 题目                                                                          | 说明                                                             | 实现                                                                            |
|-----------------------------------------------------------------------------|----------------------------------------------------------------|-------------------------------------------------------------------------------|
| [456. 132 模式](https://leetcode.cn/problems/132-pattern/) | 从右到左，单调递减栈，k记录上次被pop出去的值，一定有 栈顶 < k                            | [我的提交](https://leetcode.cn/problems/132-pattern/submissions/470350335/) |
| [84. 柱状图中最大的矩形](https://leetcode.cn/problems/largest-rectangle-in-histogram/description/) | 以每个元素为基准形成的最大矩形面积，记录前面连续大于等于它和后面连续大于等于它的个数，使用单调栈求nextSmaller即可 | [我的提交](https://leetcode.cn/problems/largest-rectangle-in-histogram/submissions/485671889/) |
| [1673. 找出最具竞争力的子序列](https://leetcode.cn/problems/find-the-most-competitive-subsequence/description/) | 维护单调递增栈，每个元素能挤掉栈尾大于它的，且能挤掉后剩余元素足够k个的                           | [我的提交](https://leetcode.cn/problems/find-the-most-competitive-subsequence/submissions/494296537/) |

> 3> 括号匹配，从头至尾每个位置累计的左括号数量 > 右括号数量
>
| 题目                                                                         | 说明                                     | 实现                                                                           |
|----------------------------------------------------------------------------|----------------------------------------|------------------------------------------------------------------------------|
| [301. 删除无效的括号](https://leetcode.cn/problems/remove-invalid-parentheses/description/) | 从左到右，判断每个字符删除/不删除，维护leftCnt < rightCnt | [我的提交](https://leetcode.cn/problems/remove-invalid-parentheses/submissions/487992077/) |
| [1541. 平衡括号字符串的最少插入次数](https://leetcode.cn/problems/minimum-insertions-to-balance-a-parentheses-string/description/) | 维护左括号数量始终>=0                           | [我的提交](https://leetcode.cn/problems/minimum-insertions-to-balance-a-parentheses-string/submissions/493731327/) |

> 4> 格式校验， 利用栈的特性不断消除元素,(树的先序遍历==栈)
>
| 题目                                                                        | 说明                                              | 实现                                                                           |
|---------------------------------------------------------------------------|-------------------------------------------------|------------------------------------------------------------------------------|
| [331. 验证二叉树的前序序列化](https://leetcode.cn/problems/verify-preorder-serialization-of-a-binary-tree/description/) | 数值节点当作2，#节点当作0，两个0会消除上一个元素，并再入一个0               | [我的提交](https://leetcode.cn/problems/verify-preorder-serialization-of-a-binary-tree/submissions/488584791/) |
| [388. 文件的最长绝对路径](https://leetcode.cn/problems/longest-absolute-file-path/description/) | \t个数表示层数，curLayer  < 栈顶，循环pop构建树直至curLayer > 栈顶 | [我的提交](https://leetcode.cn/problems/longest-absolute-file-path/submissions/488677676/) |

> 5> 计算表达式，将数值和括号一并入栈讨论
>
| 题目                                                                        | 说明                                                   | 实现                                                                           |
|---------------------------------------------------------------------------|------------------------------------------------------|------------------------------------------------------------------------------|
| [856. 括号的分数](https://leetcode.cn/problems/score-of-parentheses/description/) | 数值和左括号入栈，遇到右括号时按栈顶类型分别处理                             | [我的提交](https://leetcode.cn/problems/score-of-parentheses/submissions/490988393/) |
| [1190. 反转每对括号间的子串](https://leetcode.cn/problems/reverse-substrings-between-each-pair-of-parentheses/description/) | 字符串和左括号入栈，遇到右括号时将与前一个左括号之间的字符串pop出来，每个字符串都单独的reverse | [我的提交](https://leetcode.cn/problems/reverse-substrings-between-each-pair-of-parentheses/submissions/492379575/) |

### 3. 更多练习

**基础**

| 题目                                                         | 解析                                          | 答案                                                         |
| ------------------------------------------------------------ | --------------------------------------------- | ------------------------------------------------------------ |
| [496. 下一个更大元素 I](https://leetcode.cn/problems/next-greater-element-i/) | 对nums2使用单调递减栈求NGE                    | [通过](https://leetcode.cn/submissions/detail/373213692/)    |
| [503. 下一个更大元素 II](https://leetcode.cn/problems/next-greater-element-ii/) | 原数组扩一倍，然后使用%运算                   | [通过](https://leetcode.cn/submissions/detail/373216518/)    |
| [739. 每日温度](https://leetcode.cn/problems/daily-temperatures/) | 保存 NGE 数和对应的下标                       | [通过](https://leetcode.cn/submissions/detail/373218830/)    |
| [1019. 链表中的下一个更大节点](https://leetcode.cn/problems/next-greater-node-in-linked-list/) | 和 496 一样，只是需要先遍历链表得到数组       | [通过](https://leetcode.cn/submissions/detail/373495090/)    |
| [907. 子数组的最小值之和](https://leetcode.cn/problems/sum-of-subarray-minimums/) | 求 a\[i] 的左右边界，两次使用单调递增栈求边界 | [0x3F](https://leetcode.cn/problems/sum-of-subarray-minimums/solution/gong-xian-fa-dan-diao-zhan-san-chong-shi-gxa5/) |



**进阶**

| 题目                                                         | 解析                                                         | 答案                                                         |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| [42. 接雨水](https://leetcode.cn/problems/trapping-rain-water/) | 栈底到栈顶是**递减**的，注意**凹槽**的位置 :fire:            | [通过](https://leetcode.cn/submissions/detail/419904258/)    |
| [84. 柱状图中最大的矩形](https://leetcode.cn/problems/largest-rectangle-in-histogram/) | 栈底到栈顶是**递增**的，注意**凸起**的位置 :fire:            | [通过](https://leetcode.cn/submissions/detail/373561577/)    |
| [901. 股票价格跨度](https://leetcode.cn/problems/online-stock-span/) | 单调递减栈，先记下入栈之前数的下标                           | [通过](https://leetcode.cn/submissions/detail/375280534/)    |
| [402. 移掉 K 位数字](https://leetcode.cn/problems/remove-k-digits/) | **单调递增栈**，从头开始遍历，最终去掉前导0                  | [通过](https://leetcode.cn/submissions/detail/373237985/)    |
| [1673. 找出最具竞争力的子序列](https://leetcode.cn/problems/find-the-most-competitive-subsequence/) | 和 402 类似，不需要去掉前导0                                 | [通过](https://leetcode.cn/submissions/detail/373248009/)    |
| [1081. 不同字符的最小子序列](https://leetcode.cn/problems/smallest-subsequence-of-distinct-characters/) | 和 [316. 去除重复字母](https://leetcode.cn/problems/remove-duplicate-letters/) 一样，402 升级版，统计每个字符数目， | [通过](https://leetcode.cn/submissions/detail/373260316/)    |
| [321. 拼接最大数](https://leetcode.cn/problems/create-maximum-number/) | 单调递减栈，数组字典序最大（cpp造轮子），归并                | [通过](https://leetcode.cn/submissions/detail/373517666/)    |
| [1124. 表现良好的最长时间段](https://leetcode.cn/problems/longest-well-performing-interval/) | 确切说不是单调栈，思想类似，前缀和+单调栈：找递减序列并且从后往前遍历，前缀和+哈希表：记录 `s <= 0` 的最小下标 | [0x3F](https://leetcode.cn/problems/longest-well-performing-interval/solution/liang-chong-zuo-fa-liang-zhang-tu-miao-d-hysl/) |

> 正向遍历，移除元素或者保留元素使得剩下的数字最小（最大）或者剩下的序列字典序最小（最大）



**单调队列**

| 题目                                                         | 解析                                                         | 答案                                                         |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| [239. 滑动窗口最大值](https://leetcode.cn/problems/sliding-window-maximum/) | **双端队列**，很严格递减的，存储元素比较好理解               | [流程](https://leetcode.cn/problems/sliding-window-maximum/solution/shuang-xiang-dui-lie-jie-jue-hua-dong-chuang-kou-2/) |
| [862. 和至少为 K 的最短子数组](https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/) | 前缀和，**单调递增的队列**                                   | [0x3F](https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/solution/liang-zhang-tu-miao-dong-dan-diao-dui-li-9fvh/) |
| [2373. 矩阵中的局部最大值](https://leetcode.cn/problems/largest-local-values-in-a-matrix/) | 3*3数据量比较小可以直接遍历，如果比较大可以每行执行滑动窗口求最大，并且每列比较 | [图解](https://leetcode.cn/problems/largest-local-values-in-a-matrix/solution/javapythonmei-ju-mo-ni-dan-diao-dui-lie-fm0pn/) |



### 4. 参考
- [单调栈解决 Next Greater Number 一类问题](https://leetcode.cn/problems/next-greater-element-i/solution/dan-diao-zhan-jie-jue-next-greater-number-yi-lei-w/)
- [一招吃遍力扣四道题，妈妈再也不用担心我被套路啦～](https://leetcode.cn/problems/remove-k-digits/solution/yi-zhao-chi-bian-li-kou-si-dao-ti-ma-ma-zai-ye-b-5/)
- [暴力解法、栈（单调栈、哨兵技巧）](https://leetcode.cn/problems/largest-rectangle-in-histogram/solution/bao-li-jie-fa-zhan-by-liweiwei1419/)
- [单调栈—代码随想录 :fire:](https://programmercarl.com/0739.%E6%AF%8F%E6%97%A5%E6%B8%A9%E5%BA%A6.html#%E6%80%9D%E8%B7%AF)