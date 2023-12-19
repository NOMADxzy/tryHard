## 双指针

### 1. 概念

4.1 数组里的双指针

用暴力解法一定可解，双重循环得出结果。使用双指针的方法，可以借助一个额外变量，实现降维优化。
（1）相反方向运动

    两个指针在数组的头和尾，都往中间移动，直到相遇。
    两个指针在数组的中间，往数组的两端移动，直到到达边界。
    每次循环时，我们值比较当前的两个元素，在这两个元素当中，求出对应的结果。

（2）相同方向运动

    两个指针均在数组的一端，都往另一端移动，直到满足条件为止。
    两个指针的移动速度不同，类似于滑动窗口问题，快指针一直往前遍历，走到尾部则循环结束，慢指针视条件进行移动。
    每次循环时，看子区间是否满足某个条件，子区间是由双指针框起来， 输出的是子区间的变形。

4.2 双指针算法与前缀和

    二者均是对子区间进行操作，因此有一定的结合可能。
    前缀和是需要用到子区间和时，通过借助一个数组，去存储前缀和。

4.3 链表里的双指针

    以快慢指针为主，快指针一次 2 步，慢指针一次 1 步等等。
    典型例题：在链表当中找一个环，有环的话，快慢指针必定会相遇，若无环的话，则快指针就直接走到结尾。

### 2. 解题技巧（我的总结）

> 1> 有序数组采用双指针 合并/查找...
> 
| 题目                                                                            | 说明                    | 实现                                                                            |
|-------------------------------------------------------------------------------|-----------------------|-------------------------------------------------------------------------------|
| [4. 寻找两个正序数组的中位数](https://leetcode.cn/problems/median-of-two-sorted-arrays/description/) | 两个数组中有序右移较小的元素指针，直到中点 | [我的提交](https://leetcode.cn/problems/median-of-two-sorted-arrays/submissions/484064191/) |

> 2> 夹逼定理
>
| 题目                                                                            | 说明                                        | 实现                                                                            |
|-------------------------------------------------------------------------------|-------------------------------------------|-------------------------------------------------------------------------------|
| [633. 平方数之和](https://leetcode.cn/problems/sum-of-square-numbers/description/) | 想象成二维数组，每一步都只是 排除一行左/一列下 的所有元素，从而不会漏掉目标元素 | [我的提交](https://leetcode.cn/problems/sum-of-square-numbers/submissions/489916110/) |



### 3. 更多练习

| 题目                                                         | 说明                                                         | 题解                                                         |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| [76. 最小覆盖子串](https://leetcode.cn/problems/minimum-window-substring/) | 通过哈希表 needHash 判断滑动窗口包含了T的所有元素，使用 needCnt 免除每次遍历哈希表 :fire: | [滑窗](https://leetcode.cn/problems/minimum-window-substring/solution/tong-su-qie-xiang-xi-de-miao-shu-hua-dong-chuang-k/) |
| [395. 至少有 K 个重复字符的最长子串](https://leetcode.cn/problems/longest-substring-with-at-least-k-repeating-characters/) | 直接双指针无法解决，加入`枚举目标串字符种类数 [1, 26]`条件 就可以让区间具有二段性 :fire:，另外可以用[分治递归](https://leetcode.cn/problems/longest-substring-with-at-least-k-repeating-characters/solution/jie-ben-ti-bang-zhu-da-jia-li-jie-di-gui-obla/)求解 | [宫水三叶](https://leetcode.cn/problems/longest-substring-with-at-least-k-repeating-characters/solution/xiang-jie-mei-ju-shuang-zhi-zhen-jie-fa-50ri1/) |
| [1004. 最大连续1的个数 III](https://leetcode.cn/problems/max-consecutive-ones-iii/) | 维护区间 0 的个数，找出一个最长的子数组，该子数组内最多允许有 K 个 0（**正难则反**） | [负雪明烛](https://leetcode.cn/problems/max-consecutive-ones-iii/solution/fen-xiang-hua-dong-chuang-kou-mo-ban-mia-f76z/) |
| [1234. 替换子串得到平衡字符串](https://leetcode.cn/problems/replace-the-substring-for-balanced-string/) | 待替换子串**之外**的任意字符的出现次数都不能超过 m = n/4，维护**区间外**的字符出现次数 | [0x3F](https://leetcode.cn/problems/replace-the-substring-for-balanced-string/solution/tong-xiang-shuang-zhi-zhen-hua-dong-chua-z7tu/) |
| [1438. 绝对差不超过限制的最长连续子数组](https://leetcode.cn/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/) | 维护区间内**最大值和最小值的差值**，使用底层是红黑树的有序集合 multiset | [负雪明烛](https://leetcode.cn/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/solution/he-gua-de-shu-ju-jie-gou-hua-dong-chuang-v46j/) |
| [1658. 将 x 减到 0 的最小操作数](https://leetcode.cn/problems/minimum-operations-to-reduce-x-to-zero/) | 和 2516 类似，**逆向思考更直观**，正向思考须分前缀后缀考虑   | [0x3F](https://leetcode.cn/problems/minimum-operations-to-reduce-x-to-zero/solution/ni-xiang-si-wei-pythonjavacgo-by-endless-b4jt/) |
| [2379. 得到 K 个黑块的最少涂色次数](https://leetcode.cn/problems/minimum-recolors-to-get-k-consecutive-black-blocks/) | 最基本的滑动窗口，窗口大小固定                               | [通过](https://leetcode.cn/submissions/detail/411008458/)    |
| [2516. 每种字符至少取 K 个](https://leetcode.cn/problems/take-k-of-each-character-from-left-and-right/) | 正向思考需要先考虑前缀为空，然后枚举前缀，反向思考更容易理解，直接**滑动窗口取中间的字符**至多x个 | [双指针](https://leetcode.cn/submissions/detail/391726791/) [滑窗](https://leetcode.cn/submissions/detail/393522252/) |

### 4. 参考
1. [基础算法-双指针算法](https://blog.csdn.net/weixin_45891612/article/details/127993189) 