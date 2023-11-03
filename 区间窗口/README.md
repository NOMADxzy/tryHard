## 区间窗口

### 1. 概念
区间问题，就是线段问题，让你合并所有线段、找出线段的交集。两个技巧，画图+排序！

### 2. 解题技巧（我的总结）

> 1> 区间重叠问题，排序后，用一个变量记录前面最右端rightMax
> 当前区间[l,r] ,
> r<=rightMax{被重叠}
> r>rightMax && l==上一个的左端点 && 上一个区间没被重叠过 {重叠上一个}
> 
| 题目                                                                            | 说明       | 实现                                                                            |
|-------------------------------------------------------------------------------|----------|-------------------------------------------------------------------------------|
| [1288. 删除被覆盖区间](https://leetcode.cn/problems/remove-covered-intervals/description/) | 注意可能重复计算 | [我的提交](https://leetcode.cn/problems/remove-covered-intervals/submissions/479551049/) |




### 3. 更多练习


### 4. 参考
1. 大部分参考自：[算法基础——枚举](https://blog.csdn.net/weixin_45652283/article/details/131244459?utm_medium=distribute.pc_relevant.none-task-blog-2~default~baidujs_baidulandingword~default-1-131244459-blog-129442726.235^v38^pc_relevant_sort_base3&spm=1001.2101.3001.4242.2&utm_relevant_index=4) 