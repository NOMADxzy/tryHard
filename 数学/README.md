## 枚举

### 1. 概念
1⃣️位运算
2⃣️加法器
3⃣️模运算
```text
1.基本性质：amod b的结果一定小于b。
2.加法性质： (a + b) mod c = (a mod c + b mod c) mod c。
3.乘法性质： (a * b) mod c = (a mod c * b mod c) mod c。
4.减法性质： (a - b) mod c = (a mod c- b mod c + c) mod c。
5.幂运算性质： a^b mod c = (a mod c)^b mod c。
```
4⃣️素数 

### 2. 解题技巧（我的总结）

> 1> 位表示信息
> 
| 题目                                                                           | 说明                                                                 | 实现                                                                            |
|------------------------------------------------------------------------------|--------------------------------------------------------------------|-------------------------------------------------------------------------------|
| [458. 可怜的小猪](https://leetcode.cn/problems/poor-pigs/description/) | 共k个时段内检测出结果的信息向量长度 = buckets 的 k+1进制数长度                            | [我的提交](https://leetcode.cn/problems/poor-pigs/submissions/488982309/) |
| [1558. 得到目标数组的最少函数调用次数](https://leetcode.cn/problems/minimum-numbers-of-function-calls-to-make-target-array/description/) | 减次数 = 二进制1的个数之和，除次数 = 二进制最高位数                                      | [我的提交](https://leetcode.cn/problems/minimum-numbers-of-function-calls-to-make-target-array/solutions/403707/de-dao-mu-biao-shu-zu-de-zui-shao-han-shu-diao-y-2/) |
| [1734. 解码异或后的排列](https://leetcode.cn/problems/decode-xored-permutation/description/) | 偶数和其加1值的异或结果为1，故题中所有数异或结果为1或0，给第一个数初值0，求出所有后续的数，再将所有数都异或真实firstVal | [我的提交](https://leetcode.cn/problems/decode-xored-permutation/submissions/494590270/) |


> 2> 素数查找，最小堆
> ①埃筛法
> ②多路归并
>
| 题目                                                                 | 说明                                                   | 实现                                                                            |
|--------------------------------------------------------------------|------------------------------------------------------|-------------------------------------------------------------------------------|
| [313. 超级丑数](https://leetcode.cn/problems/super-ugly-number/description/) | 所有丑数由已有丑数 * 素数列表诞生，从小到大多路归并                          | [我的提交](https://leetcode.cn/problems/super-ugly-number/submissions/468697927/) |
| [786. 第 K 个最小的素数分数](https://leetcode.cn/problems/k-th-smallest-prime-fraction/) | 所有丑数由已有丑数 * 素数列表诞生，从小到大多路归并                          | [我的提交](https://leetcode.cn/problems/k-th-smallest-prime-fraction/submissions/468733547/) |
| [1390. 四因数](https://leetcode.cn/problems/four-divisors/description/) | num有4个因数 等价于 num = x*y 其中x是素数，且y是素数或y=x*x            | [我的提交](https://leetcode.cn/problems/four-divisors/submissions/500598218/) |
| [1447. 最简分数](https://leetcode.cn/problems/simplified-fractions/description/) | 思想（多个非最简分数都对应前面一个最简分数）x/y，从小到大每找到一个最简就将所有的i*x/i*y给标记 | [我的提交](https://leetcode.cn/problems/simplified-fractions/submissions/500359700/) |
| [1508. 子数组和排序后的区间和](https://leetcode.cn/problems/range-sum-of-sorted-subarray-sums/description/) | 多路归并递增产生区间和                                          | [我的提交](https://leetcode.cn/problems/range-sum-of-sorted-subarray-sums/submissions/468741186/) |
| [2523. 范围内最接近的两个质数](https://leetcode.cn/problems/closest-prime-numbers-in-range/description/) | 线性筛 更高效，数组、哈希表预分配空间降低耗时                              | [我的提交](https://leetcode.cn/problems/closest-prime-numbers-in-range/submissions/502400202/) |

> 3> 汉明距离，每个维度的汉明距离相互独立
>
| 题目                                                                          | 说明                                                   | 实现                                                                            |
|-----------------------------------------------------------------------------|------------------------------------------------------|-------------------------------------------------------------------------------|
| [477. 汉明距离总和](https://leetcode.cn/problems/total-hamming-distance/description/) | 记录数组中所有数每个汉明位上1出现的个数，则最终每个位导致的距离 = cnt1 *（m - cnt1 ） | [我的提交](https://leetcode.cn/problems/poor-pigs/submissions/488982309/) |

> 4> 最大公约数，最小公倍数，辗转相除法
>
| 题目                                                                          | 说明                                  | 实现                                                                            |
|-----------------------------------------------------------------------------|-------------------------------------|-------------------------------------------------------------------------------|
| [592. 分数加减运算](https://leetcode.cn/problems/fraction-addition-and-subtraction/description/) | a,b 的最小公倍数 = a * b / (a,b的最大公约数)    | [我的提交](https://leetcode.cn/problems/fraction-addition-and-subtraction/submissions/489696067/) |
| [1806. 还原排列的最少操作步数](https://leetcode.cn/problems/minimum-number-of-operations-to-reinitialize-a-permutation/description/) | a,b 的最小公倍数 = a * b / (a,b的最大公约数)    | [我的提交](https://leetcode.cn/problems/minimum-number-of-operations-to-reinitialize-a-permutation/submissions/495170590/) |
| [2001. 可互换矩形的组数](https://leetcode.cn/problems/number-of-pairs-of-interchangeable-rectangles/description/) | 可替换的矩形，对应着唯一的最简比例，即组合计数问题，采用用容斥定理解决 | [我的提交](https://leetcode.cn/problems/number-of-pairs-of-interchangeable-rectangles/submissions/500720933/) |

> 5> 几何图形问题
>
| 题目                                                                         | 说明                    | 实现                                                                            |
|----------------------------------------------------------------------------|-----------------------|-------------------------------------------------------------------------------|
| [593. 有效的正方形](https://leetcode.cn/problems/valid-square/description/) | 枚举可能性，不断尝试，注意边界条件     | [我的提交](https://leetcode.cn/problems/valid-square/submissions/489702623/) |
| [858. 镜面反射](https://leetcode.cn/problems/mirror-reflection/) | 补出镜像即可                | [我的提交](https://leetcode.cn/problems/mirror-reflection/submissions/490986394/) |
| [1401. 圆和矩形是否有重叠](https://leetcode.cn/problems/circle-and-rectangle-overlapping/description/) | 找矩形上距离圆心的最近点（根据圆心的位置） | [我的提交](https://leetcode.cn/problems/circle-and-rectangle-overlapping/submissions/493163161/) |


> 6> 数学分析，简化问题
>
| 题目                                                                         | 说明                                                  | 实现                                                                            |
|----------------------------------------------------------------------------|-----------------------------------------------------|-------------------------------------------------------------------------------|
| [754. 到达终点数字](https://leetcode.cn/problems/reach-a-number/description/) | 考虑sum-target能被2整除                                   | [我的提交](https://leetcode.cn/problems/reach-a-number/submissions/490423313/) |
| [991. 坏了的计算器](https://leetcode.cn/problems/broken-calculator/description/) | 考虑先*2到大于target 再由大到小减去2的k次方直至得到target               | [我的提交](https://leetcode.cn/problems/broken-calculator/submissions/491418210/) |
| [1432. 改变一个整数能得到的最大差值](https://leetcode.cn/problems/max-difference-you-can-get-from-changing-an-integer/description/) | 分别从高往低是探索，考虑最大和最小，一定注意111不能缩小为100                   | [我的提交](https://leetcode.cn/problems/max-difference-you-can-get-from-changing-an-integer/submissions/493357212/) |
| [1503. 所有蚂蚁掉下来前的最后一刻](https://leetcode.cn/problems/last-moment-before-all-ants-fall-out-of-a-plank/description/) | 等价于交换身份，不发生碰撞                                       | [我的提交](https://leetcode.cn/problems/last-moment-before-all-ants-fall-out-of-a-plank/submissions/493627685/) |
| [1814. 统计一个数组中好对子的数目](https://leetcode.cn/problems/count-nice-pairs-in-an-array/description/) | 等价于 x-rev(x) == y-rev(y)                            | [我的提交](https://leetcode.cn/problems/count-nice-pairs-in-an-array/submissions/495172572/) |
| [2165. 重排数字的最小值](https://leetcode.cn/problems/smallest-value-of-the-rearranged-number/description/) | 只需要统计每种数字的个数，然后贪心的填数字就行了                            | [我的提交](https://leetcode.cn/problems/smallest-value-of-the-rearranged-number/submissions/500806976/) |
| [面试题 01.07. 旋转矩阵](https://leetcode.cn/problems/rotate-matrix-lcci/description/) | 只需遍历左上角的元素， 依次计算旋转后的位置，四个位置循环替换                     | [我的提交](https://leetcode.cn/problems/rotate-matrix-lcci/submissions/534003578/) |
| [2871. 将数组分割成最多数目的子数组](https://leetcode.cn/problems/split-array-into-maximum-number-of-subarrays/description/) | A&B 的 结果一定小于 A或B，故可以推出一定是分成若干份 与结果 为0的子数组（最后一份可不为0） | [我的提交](https://leetcode.cn/problems/split-array-into-maximum-number-of-subarrays/submissions/540664858/) |


> 7> 模运算问题
>
| 题目                                                                         | 说明                                                       | 实现                                                                            |
|----------------------------------------------------------------------------|----------------------------------------------------------|-------------------------------------------------------------------------------|
| [1015. 可被 K 整除的最小整数](https://leetcode.cn/problems/smallest-integer-divisible-by-k/description/) | newVal = 10*val + 1，故只需记录上一个val的模值，即可推出当前模值，出现的值且一定是周期性的 | [我的提交](https://leetcode.cn/problems/smallest-integer-divisible-by-k/submissions/491568700/) |
| [1497. 检查数组对是否可以被 k 整除](https://leetcode.cn/problems/check-if-array-pairs-are-divisible-by-k/description/) | 记录所有mod结果出现的次数，判断每一对是否相等                                 | [我的提交](https://leetcode.cn/problems/check-if-array-pairs-are-divisible-by-k/submissions/493611446/) |
| [1625. 执行操作后字典序最小的字符串](https://leetcode.cn/problems/lexicographically-smallest-string-after-applying-operations/description/) | 循环移动问题，总长度为n，增量为d时，最大增量Dmax = （i*d % n) ==0 时的i取值        | [我的提交](https://leetcode.cn/problems/lexicographically-smallest-string-after-applying-operations/submissions/494147403/) |
| [2575. 找出字符串的可整除数组](https://leetcode.cn/problems/find-the-divisibility-array-of-a-string/description/) | 从左到右记录前缀的模结果，每次更新为 (pre * 10%m) % m                      | [我的提交](https://leetcode.cn/problems/find-the-divisibility-array-of-a-string/submissions/512283172/) |
| [LCR 004. 只出现一次的数字 II](https://leetcode.cn/problems/WGki4K/description/) | 记录每个位上的出现次数v，最后答案的每个位取值为 v%3，注意负数    | [我的提交](https://leetcode.cn/problems/WGki4K/submissions/534308897/) |


> 8> 二进制，模拟加法器
>
| 题目                                                                         | 说明                               | 实现                                                                            |
|----------------------------------------------------------------------------|----------------------------------|-------------------------------------------------------------------------------|
| [1404. 将二进制表示减到 1 的步骤数](https://leetcode.cn/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one/description/) | 从右到左模拟加法器，c表示进位，ans遇'0'加1，遇'1'加2 | [我的提交](https://leetcode.cn/problems/number-of-steps-to-reduce-a-number-in-binary-representation-to-one/submissions/493169292/) |

> 9> 数学归纳法：从初始简单情况出发，递推出后面所有情况
>
| 题目                                                              | 说明                                                | 实现                                                                            |
|-----------------------------------------------------------------|---------------------------------------------------|-------------------------------------------------------------------------------|
| [1798. 你能构造出连续值的最大数目](https://leetcode.cn/problems/maximum-number-of-consecutive-values-you-can-make/description/) | 贪心的从小到大排序，归纳法：初始{0}连续， 前i个的和连续 -> 前i+1和也连续，否则终止循环 | [我的提交](https://leetcode.cn/problems/maximum-number-of-consecutive-values-you-can-make/submissions/495128436/) |
| [1798. 你能构造出连续值的最大数目](https://leetcode.cn/problems/number-of-ways-to-buy-pens-and-pencils/description/) | 贪心的从小到大排序，归纳法：初始{0}连续， 前i个的和连续 -> 前i+1和也连续，否则终止循环 | [我的提交](https://leetcode.cn/problems/maximum-number-of-consecutive-values-you-can-make/submissions/495128436/) |

> 10> 异或运算：可以任意调换顺序、任何数异或0都是本身、任何数异或本身都是0
>
| 题目                                                         | 说明                                         | 实现                                                                            |
|------------------------------------------------------------|--------------------------------------------|-------------------------------------------------------------------------------|
| [2425. 所有数对的异或和](https://leetcode.cn/problems/bitwise-xor-of-all-pairings/description/) | 对方长度是偶数则最终表达式中有本数组的所有元素                    | [我的提交](https://leetcode.cn/problems/bitwise-xor-of-all-pairings/submissions/500083980/) |
| [2240. 买钢笔和铅笔的方案数](https://leetcode.cn/problems/minimize-xor/description/) | 依次选择各个item，dp[i]表示 有i数量钱时购买任意数量前k个item的方案数 | [我的提交](https://leetcode.cn/problems/number-of-ways-to-buy-pens-and-pencils/submissions/500810358/) |

> 11> 分类讨论
>
| 题目                                                         | 说明                                | 实现                                                                            |
|------------------------------------------------------------|-----------------------------------|-------------------------------------------------------------------------------|
| [2249. 统计圆内格点数目](https://leetcode.cn/problems/count-lattice-points-inside-a-circle/description/) | 按每行分类，将每个圆在所有行上的截距添加到对应列表中，最后统一处理 | [我的提交](https://leetcode.cn/problems/count-lattice-points-inside-a-circle/submissions/501024130/) |


### 3. 更多练习


### 4. 参考
1. 大部分参考自：[算法基础——枚举](https://blog.csdn.net/weixin_45652283/article/details/131244459?utm_medium=distribute.pc_relevant.none-task-blog-2~default~baidujs_baidulandingword~default-1-131244459-blog-129442726.235^v38^pc_relevant_sort_base3&spm=1001.2101.3001.4242.2&utm_relevant_index=4) 
2. 总库：[tryHard](https://github.com/NOMADxzy/tryHard)