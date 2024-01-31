## 枚举

### 1. 概念
之所以称它们为高级的数据结构，是因为它们的实现要比那些常用的数据结构要复杂很多，能够让我们在处理复杂问题的过程中，
多拥有一把利器，同时掌握好它们的性质，以及所适应的场合，在分析问题的时候回归本质，那么很多问题都能迎刃而解了。

### 2. 总结

> 1> [前缀树](https://blog.csdn.net/DeveloperFire/article/details/128861092)
> 
| 题目                                                                          | 说明                                                    | 实现                                                                            |
|-----------------------------------------------------------------------------|-------------------------------------------------------|-------------------------------------------------------------------------------|
| [1268. 搜索推荐系统](https://leetcode.cn/problems/search-suggestions-system/description/) | 一眼前缀树                    | [我的提交](https://leetcode.cn/problems/search-suggestions-system/submissions/479543984/) |
| [1233. 删除子文件夹](https://leetcode.cn/problems/remove-sub-folders-from-the-filesystem/description/) | 使用哈希表存储子前缀树，ref存储在列表中的位置 | [我的提交](https://leetcode.cn/problems/remove-sub-folders-from-the-filesystem/submissions/482986495/) |
| [140. 单词拆分 II](https://leetcode.cn/problems/word-break-ii/description/) | 前缀树                      | [我的提交](https://leetcode.cn/problems/word-break-ii/submissions/487470460/) |
| [212. 单词搜索 II](https://leetcode.cn/problems/word-search-ii/description/) | 对已搜索到的单词剪去，大大降低搜索时间      | [我的提交](https://leetcode.cn/problems/word-search-ii/submissions/487614649/) |
| [648. 单词替换](https://leetcode.cn/problems/replace-words/description/) | 前缀树                      | [我的提交](https://leetcode.cn/problems/replace-words/submissions/489921442/) |


> 2> [树状数组](https://zhuanlan.zhihu.com/p/546893960)
> ①普通
> ②离散化
>
| 题目                                                                        | 说明                                            | 实现                                                                            |
|---------------------------------------------------------------------------|-----------------------------------------------|-------------------------------------------------------------------------------|
| [315. 计算右侧小于当前元素的个数](https://leetcode.cn/problems/count-of-smaller-numbers-after-self/) | 前缀和                                           | [我的提交](https://leetcode.cn/problems/count-of-smaller-numbers-after-self/submissions/488075059/) |
| [2250. 统计包含每个点的矩形数目](https://leetcode.cn/problems/count-number-of-rectangles-containing-each-point/description/) | 前缀和                                           | [我的提交](https://leetcode.cn/problems/count-number-of-rectangles-containing-each-point/submissions/469494008/) |
| [775. 全局倒置与局部倒置](https://leetcode.cn/problems/global-and-local-inversions/description/) | 前缀和                                           | [我的提交](https://leetcode.cn/problems/global-and-local-inversions/submissions/490980495/) |
| [327. 区间和的个数](https://leetcode.cn/problems/count-of-range-sum/description/) | 记录前缀和，每个i结尾的合法区间数 = 特定区间内前缀和的的数目，数据太大所以预处理离散化 | [我的提交](https://leetcode.cn/problems/count-of-range-sum/submissions/499683923/) |

> 3> [跳表](https://zhuanlan.zhihu.com/p/637407262?utm_id=0)


### 3. 更多练习


### 4. 参考
1. 大部分参考自：[算法基础——枚举](https://blog.csdn.net/weixin_45652283/article/details/131244459?utm_medium=distribute.pc_relevant.none-task-blog-2~default~baidujs_baidulandingword~default-1-131244459-blog-129442726.235^v38^pc_relevant_sort_base3&spm=1001.2101.3001.4242.2&utm_relevant_index=4) 