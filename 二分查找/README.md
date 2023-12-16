## 二分查找

### 1. 概念
二分查找就是从一个排好序的序列中查找一个元素，和顺序查找不同的是，二分查找通过逐步缩小搜索区间的方式将搜索范围逐渐缩小，最终定位到我们的
目标值或者目标位置，**在STL中，`lower_bound()` 和 `upper_bound()` 就是利用二分的思想，前者在有序数组中找到第一个大于等于的目标值
的位置，后者找到第一个大于目标值的位置**


1. **模板一**

```go
!check(left) && check(right)

for left<right{
     mid := (left+right)/2
     if !check(left){
         left = mid + 1
     }else {
         right = mid
     }
 }
```

2**模板二**

```go
!check(left) && check(right)

for left<right{
     mid := (left+right+1)/2
     if !check(left){
         left = mid
     }else {
         right = mid + 1
     }
 }
```


### 2. 解题技巧（我的总结）

> 1> 有时不能一眼看出来，需要提炼有序部分
> 
| 题目                                                               | 说明                                          | 实现                                                                            |
|------------------------------------------------------------------|---------------------------------------------|-------------------------------------------------------------------------------|
| [792. 匹配子序列的单词数](https://leetcode.cn/problems/number-of-matching-subsequences/description/) | 记录26个字母分别在s中从前到后出现的位置                       | [我的提交](https://leetcode.cn/problems/number-of-matching-subsequences/submissions/474655365/) |
| [1477. 找两个和为目标值且不重叠的子数组](https://leetcode.cn/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/description/) | 因为元素都是>0, 所以前缀和是有序的，求每个位置结尾的子数组为target的长度   | [我的提交](https://leetcode.cn/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/submissions/477743362/) |
| [1838. 最高频元素的频数](https://leetcode.cn/problems/frequency-of-the-most-frequent-element/description/) | 排序，要变成i元素的区间的左边界越小，对应的k越大，从而探索最大左边界         | [我的提交](https://leetcode.cn/problems/frequency-of-the-most-frequent-element/submissions/480483834/) |
| [483. 最小好进制](https://leetcode.cn/problems/smallest-good-base/description/) | 1的位数固定时，表示的值随着进制的增大而增大，从64到2取位数，二分查找进制判断可行性 | [我的提交](https://leetcode.cn/problems/smallest-good-base/submissions/489297643/) |

> 2> 答案为固定数值时，可以直接对答案进行二分查找
>
| 题目                                                                                             | 说明                                   | 实现                                                                            |
|------------------------------------------------------------------------------------------------|--------------------------------------|-------------------------------------------------------------------------------|
| [287. 寻找重复数](https://leetcode.cn/problems/find-the-duplicate-number/description/)              | 答案是数组最小值和最大值之间的一个值                   | [我的提交](https://leetcode.cn/problems/find-the-duplicate-number/submissions/465110299/) |
| [719. 找出第 K 小的数对距离](https://leetcode.cn/problems/find-k-th-smallest-pair-distance/description/) | 排序数组 答案是 0 和 (右端点-左端点) 之间的一个值，需要二重二分 | [我的提交](https://leetcode.cn/problems/find-k-th-smallest-pair-distance/submissions/468777437/) |
| [875. 爱吃香蕉的珂珂](https://leetcode.cn/problems/koko-eating-bananas/) | 答案是数组最小值和最大值之间的一个                    | [我的提交](https://leetcode.cn/problems/koko-eating-bananas/submissions/466078855/) |
| [1300. 转变数组后最接近目标值的数组和](https://leetcode.cn/problems/sum-of-mutated-array-closest-to-target/description/) | 答案是数组最大值和0之间的一个                      | [我的提交](https://leetcode.cn/problems/sum-of-mutated-array-closest-to-target/submissions/479566095/) |

> 3> 数组无须有序，只要中点两边特性不同
>
| 题目                                                                                                  | 说明               | 实现                                                                            |
|-----------------------------------------------------------------------------------------------------|------------------|-------------------------------------------------------------------------------|
| [153. 寻找旋转排序数组中的最小值](https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/description/) | 左半部分大于右端点，右半部分小于左端点 | [我的提交](https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/submissions/465090277/) |
| [162. 寻找峰值](https://leetcode.cn/problems/find-peak-element/description/) | 两个大于两边的区间中一定存在峰值 | [我的提交](https://leetcode.cn/problems/find-peak-element/submissions/466060281/) |

> 4> 求某一段区间时，分别二分查找区间的left和right点
>
| 题目                                                                                            | 说明                               | 实现                                                                            |
|-----------------------------------------------------------------------------------------------|----------------------------------|-------------------------------------------------------------------------------|
| [1712. 将数组分成三个子数组的方案数](https://leetcode.cn/problems/ways-to-split-array-into-three-subarrays/) | 固定left和mid交点，求mid区间右端点的取值区间，累加区间长度 | [我的提交](https://leetcode.cn/problems/ways-to-split-array-into-three-subarrays/submissions/473290077/) |
| [2563. 统计公平数对的数目](https://leetcode.cn/problems/count-the-number-of-fair-pairs/description/) | i，求j的取值区间，累加区间长度 | [我的提交](https://leetcode.cn/problems/count-the-number-of-fair-pairs/description/) |

> 5> 依次从数组中找出最接近某一目标的值，通过排序优化时间
>
| 题目                                                                                           | 说明                                      | 实现                                                                            |
|----------------------------------------------------------------------------------------------|-----------------------------------------|-------------------------------------------------------------------------------|
| [475. 供暖器](https://leetcode.cn/problems/heaters/description/) | max(某一house左右最近的两个供暖器的最短距离) for 所有house | [我的提交](https://leetcode.cn/problems/heaters/submissions/489172246/) |


### 3. 更多练习

#### 3.1. 二分求下标

在给定数组中查找符合条件的元素下标

| 题目                                                         | 说明                                                         | 题解                                                         |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| [704.二分查找](https://leetcode.cn/problems/binary-search/)  | 简单的二分即可，常规思路                                     |                                                              |
| [35.搜索插入位置](https://leetcode.cn/problems/search-insert-position/) | 找到插入位置的索引，可以是len，因此可以设置right=len         | [解](https://leetcode.cn/problems/search-insert-position/solution/te-bie-hao-yong-de-er-fen-cha-fa-fa-mo-ban-python-/) |
| [300.最长上分子序列](https://leetcode.cn/problems/longest-increasing-subsequence/solution/) | 这题DP思路比较简单，二分查找比较难                           | [解](https://leetcode.cn/problems/longest-increasing-subsequence/solution/dong-tai-gui-hua-er-fen-cha-zhao-tan-xin-suan-fa-p/) |
| [34.排序数组查找第一个和最后一个位置](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/) | 可以总结为两个模板：**查找左边界**和**查找右边界**           | [解](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/solution/tu-jie-er-fen-zui-qing-xi-yi-dong-de-jia-ddvc/) |
| [611.有效三角形的个数](https://leetcode.cn/problems/valid-triangle-number/) | **二分有两种思路**：左边界和右边界，**双指针**：固定最长边逆序 | [解1](https://leetcode.cn/problems/valid-triangle-number/solution/ming-que-tiao-jian-jin-xing-qiu-jie-by-jerring/) [解2](https://leetcode.cn/problems/valid-triangle-number/solution/er-fen-cha-zhao-python-dai-ma-java-dai-ma-by-liwei/) |
| [659.找到K个最接近的元素](https://leetcode.cn/problems/find-k-closest-elements/) | 二分找到左边界，双指针                                       | [解](https://leetcode.cn/problems/find-k-closest-elements/solution/pai-chu-fa-shuang-zhi-zhen-er-fen-fa-python-dai-ma/) |
| [436.寻找右区间](https://leetcode.cn/problems/find-right-interval/) | 使用哈希表记录第一个元素位置之后在[:\][0\]二分查找[:\][1\]   | [解](https://leetcode.cn/problems/find-right-interval/solution/by-fuxuemingzhu-98m1/) |
| [1237.找出给定方程的正整数解](https://leetcode.cn/problems/find-positive-integer-solution-for-a-given-equation/) | 二分利用一个变量递增，双指针利用两个变量都递增               | [解](https://leetcode.cn/problems/find-positive-integer-solution-for-a-given-equation/solution/xiang-jie-bao-li-er-fen-yu-shuang-zhi-zhen-fa-by-q/) |
| [4.寻找两个正序数组的中位数](https://leetcode.cn/problems/median-of-two-sorted-arrays/) | 直接归并比较简单，通过二分找到__分割线__比较难               | [解](https://leetcode.cn/problems/median-of-two-sorted-arrays/solution/he-bing-yi-hou-zhao-gui-bing-guo-cheng-zhong-zhao-/) |
| [6355. 统计公平数对的数目](https://leetcode.cn/problems/count-the-number-of-fair-pairs/) | **排序**之后二分，lower_bound，upper_bound                   | [通过](https://leetcode.cn/submissions/detail/401623604/)    |



**「选择数组」和「山脉数组」：局部单调性**

| 题目                                                         | 说明                                                         | 题解                                                         |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| [33.搜索旋转排序数组](https://leetcode.cn/problems/search-in-rotated-sorted-array/) | 直接在循环里面定位比较直接，**在外面定位[思考](https://leetcode.cn/problems/search-in-rotated-sorted-array/solution/er-fen-fa-python-dai-ma-java-dai-ma-by-liweiwei141/129291)很细节** | [解](https://leetcode.cn/problems/search-in-rotated-sorted-array/solution/er-fen-fa-python-dai-ma-java-dai-ma-by-liweiwei141/) |
| [81.搜索旋转排序树组II](https://leetcode.cn/problems/search-in-rotated-sorted-array-ii/) | 在 33 基础上通过`++left`去重，或者直接[while去重](https://leetcode.cn/submissions/detail/170342574/) | [解](https://leetcode.cn/problems/search-in-rotated-sorted-array-ii/solution/er-fen-cha-zhao-by-liweiwei1419/) |
| [153.旋转排列数组最小值](https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/) | 通过**比较 mid 和 right** 判断最小值在左还是右               | [解](https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/solution/er-fen-cha-zhao-wei-shi-yao-zuo-you-bu-dui-cheng-z/) |
| [154.旋转排序数组最小值II](https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array-ii/) | 在 153 基础上通过 `-- right` 去重                            | [解](https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array-ii/solution/154-find-minimum-in-rotated-sorted-array-ii-by-jyd/) |
| [852.山脉数组的峰顶索引](https://leetcode.cn/problems/peak-index-in-a-mountain-array/) | 找到**最小满足** `nums[i] > nums[i+1]` 的下标                | [解](https://leetcode.cn/problems/peak-index-in-a-mountain-array/solution/shan-mai-shu-zu-de-feng-ding-suo-yin-by-dtqvv/) |
| [1095.山脉数组找目标值](https://leetcode.cn/problems/find-in-mountain-array/) | **三个二分**：找到峰顶下标，升序找target，降序找target       | [解](https://leetcode.cn/problems/find-in-mountain-array/solution/shi-yong-chao-hao-yong-de-er-fen-fa-mo-ban-python-/) |
| [1802. 有界数组中指定下标处的最大值](https://leetcode.cn/problems/maximum-value-at-a-given-index-in-a-bounded-array/) | “山峰数列”求和 sum，找到满足 sum <= maxSum 的**最大峰**      | [解](https://leetcode.cn/problems/maximum-value-at-a-given-index-in-a-bounded-array/solution/by-lcbin-4vp4/) |



#### 3.2. 二分找答案

在给定的序列中找一个满足条件的答案，通过二分查找逐渐缩小范围，最后逼近到一个数

| 题目                                                         | 说明                                                         | 题解                                                         |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| [69.x的平方根](https://leetcode.cn/problems/sqrtx/)          | 使用除法可以避免溢出，另有一个 [牛顿迭代法](https://leetcode.cn/problems/sqrtx/solution/niu-dun-die-dai-fa-by-loafer/) | [解](https://leetcode.cn/problems/sqrtx/solution/er-fen-cha-zhao-niu-dun-fa-python-dai-ma-by-liweiw/) |
| [287.寻找重复数](https://leetcode.cn/problems/find-the-duplicate-number/) | 二分思路有点绕，[循环链表](https://leetcode.cn/problems/find-the-duplicate-number/solution/287xun-zhao-zhong-fu-shu-by-kirsche/)思路比较直接 | [解](https://leetcode.cn/problems/find-the-duplicate-number/solution/er-fen-fa-si-lu-ji-dai-ma-python-by-liweiwei1419/) |
| [374.猜数字大小](https://leetcode.cn/problems/guess-number-higher-or-lower/) | 比较简单的二分，注意一下题意就好                             |                                                              |
| [275.H指数 II](https://leetcode-cn.com/problems/h-index-ii/) | 注意向左找还是向右找的条件                                   | [解](https://leetcode-cn.com/problems/h-index-ii/solution/jian-er-zhi-zhi-er-fen-cha-zhao-by-liweiwei1419-2/) |
| [1283.使结果不超过阈值的最小除数](https://leetcode-cn.com/problems/find-the-smallest-divisor-given-a-threshold/) | 注意不整除才+1                                               | [解](https://leetcode-cn.com/problems/find-the-smallest-divisor-given-a-threshold/solution/er-fen-cha-zhao-ding-wei-chu-shu-by-liweiwei1419/) |
| [1292. 元素和小于等于阈值的正方形的最大边长](https://leetcode.cn/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold/) | 二维前缀和，遍历二分                                         | [解](https://leetcode.cn/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold/solution/yuan-su-he-xiao-yu-deng-yu-yu-zhi-de-zheng-fang-2/) |
| [剑指 Offer 53 - II. 0～n-1中缺失的数字](https://leetcode.cn/problems/que-shi-de-shu-zi-lcof/) | 左子数组 `nums[i]=i`，右子数组 `nums[i]!=i`                  | [通过](https://leetcode.cn/submissions/detail/380479872/)    |
| [1760. 袋子里最少数目的球](https://leetcode.cn/problems/minimum-limit-of-balls-in-a-bag/) | 在\[1, max(nums)\]中二分找到**开销（最大值）的最小值**       | [官解](https://leetcode.cn/problems/minimum-limit-of-balls-in-a-bag/solution/dai-zi-li-zui-shao-shu-mu-de-qiu-by-leet-boay/) |
| [2576. 求出最多标记下标](https://leetcode.cn/problems/find-the-maximum-number-of-marked-indices/) | 问题的**单调性**，二分过程贪心检查                           | [0x3F](https://leetcode.cn/problems/find-the-maximum-number-of-marked-indices/solution/er-fen-da-an-pythonjavacgo-by-endlessche-t9f5/) |



#### 3.3. 最大化最小值

| 题目                                                         | 说明                                                       | 题解                                                         |
| ------------------------------------------------------------ | ---------------------------------------------------------- | ------------------------------------------------------------ |
| [1552. 两球之间的磁力](https://leetcode.cn/problems/magnetic-force-between-two-balls/) | 猜答案 ans 越小，可以放的位置越多，越有可能 >= m，满足题意 | [边界](https://leetcode.cn/problems/magnetic-force-between-two-balls/solution/c-er-fen-sou-suo-ying-gai-neng-gei-ni-jiang-ming-b/) |
| [2517. 礼盒的最大甜蜜度](https://leetcode.cn/problems/maximum-tastiness-of-candy-basket/) | 猜答案 ans 越小，可以选择的数越多，越有可能 >= k，满足题意 | [0x3F](https://leetcode.cn/problems/maximum-tastiness-of-candy-basket/solution/er-fen-da-an-by-endlesscheng-r418/) |



#### 3.4. 最小化最大值

| 题目                                                         | 说明                                           | 题解                                                         |
| ------------------------------------------------------------ | ---------------------------------------------- | ------------------------------------------------------------ |
| [2560. 打家劫舍 IV](https://leetcode.cn/problems/house-robber-iv/) | 问题的**单调性**，二分找答案过程中 DP 进行判断 | [0x3F](https://leetcode.cn/problems/house-robber-iv/solution/er-fen-da-an-dp-by-endlesscheng-m558/) |


### 4. 参考


1. 大部分参考自：[写对二分查找不是套模板并往里面填空，需要仔细分析题意](https://leetcode.cn/problems/search-insert-position/solution/te-bie-hao-yong-de-er-fen-cha-fa-fa-mo-ban-python-/)
2. 另模板1参考自：[代码随想录](https://leetcode.cn/problems/search-insert-position/solution/by-carlsun-2-2dlr/)
3. 另模板2参考自：[图解二分](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/solution/tu-jie-er-fen-zui-qing-xi-yi-dong-de-jia-ddvc/)