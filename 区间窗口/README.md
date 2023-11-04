## 区间窗口

### 1. 概念
区间问题，就是线段问题，让你合并所有线段、找出线段的交集。两个技巧，画图+排序！

### 2. 解题技巧（我的总结）

> 1> 区间重叠问题，排序后，用一个变量记录前面最右端rightMax及其对应的索引idx
> 
| 题目                                                                            | 说明       | 实现                                                                            |
|-------------------------------------------------------------------------------|----------|-------------------------------------------------------------------------------|
| [1288. 删除被覆盖区间](https://leetcode.cn/problems/remove-covered-intervals/description/) | 注意可能重复计算 | [我的提交](https://leetcode.cn/problems/remove-covered-intervals/submissions/479551049/) |

> 1> 使用优先堆解决区间问题
>
| 题目                                                                            | 说明                        | 实现                                                                            |
|-------------------------------------------------------------------------------|---------------------------|-------------------------------------------------------------------------------|
| [1353. 最多可以参加的会议数目](https://leetcode.cn/problems/maximum-number-of-events-that-can-be-attended/description/) | 第i天应当参加所有当天可参加的会议中结束时间最早的 | [我的提交](https://leetcode.cn/problems/maximum-number-of-events-that-can-be-attended/submissions/479644501/) |




### 3. 更多练习


### 4. 参考
1. 大部分参考自：[算法基础——枚举](https://blog.csdn.net/weixin_45652283/article/details/131244459?utm_medium=distribute.pc_relevant.none-task-blog-2~default~baidujs_baidulandingword~default-1-131244459-blog-129442726.235^v38^pc_relevant_sort_base3&spm=1001.2101.3001.4242.2&utm_relevant_index=4) 