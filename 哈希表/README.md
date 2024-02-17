## 哈希表

### 1. 概念

哈希函数也叫散列函数，它对不同的输出值得到一个固定长度的消息摘要。

1>散列结果应当具有同一性（输出值尽量均匀）

2>雪崩效应（微小的输入值变化使得输出值发生巨大的变化）



**通常有以下几种构造 Hash 函数的方法：**

- 直接定址法：取关键字或关键字的某个线性函数值为散列地址。例如以年龄为关键字的散列表
- 随机数法：选择一个随机函数，把关键字的随机函数值作为它的哈希值。通常用于关键字长度不等时采用此法构造哈希函数。
- 折叠法：将关键字分为位数相同的几部分，然后取这几部分的叠加和（舍去进位）作为散列地址。
- 平方取中法：先计算出关键字值的平方，然后取平方值中间几位作为散列地址。
- 除留余数法（最常用的）：取关键字被某个不大于散列表长度 m 的数 p 求余，得到的作为散列地址。

- 数字分析法：当关键字的位数大于地址的位数，对关键字的各位分布进行分析，选出分布均匀的任意几位作为散列地址。仅适用于所有关键字都已知的情况下，根据实际应用确定要选取的部分，尽量避免发生冲突。



#### Hash表大小的确定
Hash 表的空间如果远远大于实际存储的记录数据的个数，则造成空间浪费；如果过小，则容易造成冲突。Hash 表大小确定通常有这两种思路：
如果最初知道存储的数据量，则需要根据存储个数 和 关键字的分布特点来确定 Hash 表的大小。
事先不知道最终需要存储的记录个数，需要动态维护Hash表的容量，此时可能需要重新计算 Hash 地址。

#### Hash冲突解决

```
(1)拉链法
在重复下表的下面又开了一个数组，在这个数组将重复的全部装进去，这样在查找的时候只需要遍历这个数组就ok了；
(2)开放地址法
这种方法简单来说就是当元素下标值发生冲突时，寻找空白的位置插入数据，三种方法，分别是 线性探测 、二次探测 、再哈希法，其中再哈希法就是再将我们传入的值进行一次 哈希化，获得一个新的探测步数 step，然后按照这个步数进行探测，找到第一个空着的位置插入数据。这在很大的程度上解决了 聚集 的问题。
```



#### Hash 实现方案

这几个都是集合类都是基于散列表，分析它们可以从如下几个点出发：

    线程安全：HashTable是线程安全，HashMap和HashSet不是；
    实现方式：HashMap基于拉链法的散列表，链过长会自动转为红黑树，HashSet底层采用HashMap实现的；
    初始大小：HashTable初始大小是11，HashMap初始大小是16
    空值：HashMap可以将空值作为key（一条：键不能重复）或者value（多条），HashTable不允许null值（键与值都不行）,HashSet多个null只会有一个null存在。
    扩容方式：HashTable采用：(oldCapacity << 1) + 1，HashMap采用oldCap << 1
    哈希值：HashTable直接使用对象的hashCode，而HashTable采用在对象的hashCode上还进行的处理变化；
#### Hash 表的及优劣

（1）优点

    对于一些大数据多数据，hash表处理起来比较轻松
    能够快速的 查，改，插，删元素 等操作
    代码简单（自己想好hash函数就完成啦）

（2）缺点

    在hash函数处理某些元素的时候，不免出现下标重复相同的情况，这种情况可以称作为冲突
    哈希表中的数据都是没有顺序的

### 2. 解题技巧（我的总结）

> 1> 原地哈希（题目限制空间复杂度时使用）
> 
| 题目                                                                            | 说明           | 实现                                                                            |
|-------------------------------------------------------------------------------|--------------|-------------------------------------------------------------------------------|
| [41. 缺失的第一个正数](https://leetcode.cn/problems/first-missing-positive/description/) | 数值n存到数组中的n-1位置 | [我的提交](https://leetcode.cn/problems/first-missing-positive/submissions/485428083/) |
| [LCR 120. 寻找文件副本](https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/description/) | 数值n存到数组中的n位置 | [我的提交](https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/submissions/485438142/) |
| [442. 数组中重复的数据](https://leetcode.cn/problems/find-all-duplicates-in-an-array/description/) | 数值n存到数组中的n-1位置 | [我的提交](https://leetcode.cn/problems/find-all-duplicates-in-an-array/submissions/488940179/) |

> 2> 大量无规则数据问题的简化
>
| 题目                                                                          | 说明                            | 实现                                                                            |
|-----------------------------------------------------------------------------|-------------------------------|-------------------------------------------------------------------------------|
| [554. 砖墙](https://leetcode.cn/problems/brick-wall/description/) | 前缀和的方式记录每块砖的右端点，hash找出最大出现次数的 | [我的提交](https://leetcode.cn/problems/brick-wall/submissions/489660313/) |
| [1711. 大餐计数](https://leetcode.cn/problems/count-good-meals/description/) | 记录所有数的出现次数，对每个数枚举其对面数的可能性     | [我的提交](https://leetcode.cn/problems/count-good-meals/submissions/494379244/) |

> 3> 集合交并/子集
>
| 题目                                                                          | 说明                                                         | 实现                                                                            |
|-----------------------------------------------------------------------------|------------------------------------------------------------|-------------------------------------------------------------------------------|
| [1452. 收藏清单](https://leetcode.cn/problems/people-whose-list-of-favorite-companies-is-not-a-subset-of-another-list/description/) | 记录每个字符串出现的idx列表，i位置的字符串j：j出现的位置cnt都+1，看最后是否有个位置cnt和i字符串数相同 | [我的提交](https://leetcode.cn/problems/people-whose-list-of-favorite-companies-is-not-a-subset-of-another-list/submissions/493450888/) |


### 3. 更多练习


### 4. 参考
1. [❤数据结构入门❤——哈希表](https://blog.csdn.net/bt_giegie/article/details/120572850) 
2. 总库：[tryHard](https://github.com/NOMADxzy/tryHard)