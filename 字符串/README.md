## 字符串

### 1. 概念

字符串就是特殊的字符数组，字符数组末尾的元素为 '\0'。和数组一样可以使用arr[i]或*(arr+i)来访问元素。

        无论是用数组保存字符串（如：char arr[] = "Hello，World";），还是用指针保存字符串（如：char *brr = "Hello，World";），我们都可以使用字符串函数strlen()，来计算字符串长度。因为这里数组名和指针名保存着字符串在内存中的首地址，并且这里是字符串（存储时末尾有一个隐藏的'\0'），以 '\0' 结尾。这样，对于字符数组：使用sizeof(arr)/sizeof(arr[0])，求出来的长度比strlen求出的长度大1，因为sizeof计算数组大小时包含了'\0'，strlen计算时忽略了'\0'；对于字符串指针：sizeof(brr)/sizeof(brr[0])计算结果为4，因为brr是指针名，sizeof(brr)计算的是指针的大小为4，再除以元素大小1，所以结果为4。
————————————————
版权声明：本文为CSDN博主「加油JIAX」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/weixin_50803498/article/details/126528554

### 2. 解题技巧（我的总结）

> 1> 字符串分割
> 
| 题目                                                                           | 说明                   | 实现                                                                            |
|------------------------------------------------------------------------------|----------------------|-------------------------------------------------------------------------------|
| [640. 求解方程](https://leetcode.cn/problems/solve-the-equation/description/) | 通过分割字符串统计x的系数和方程的常数项 | [我的提交](https://leetcode.cn/problems/solve-the-equation/submissions/489859079/) |

> 2> 字符串匹配
> KMP算法：next数组
```go
i, l := 1, 0
	for i < len(s) {
		if s[i] == s[l] {
			l++
			next[i] = l
			i++
		} else {
			if l == 0 {
				next[i] = 0
				i++
			} else {
				l = next[l-1]
			}
		}
	}
```
>
| 题目                                                                                                          | 说明                                         | 实现                                                                               |
|-------------------------------------------------------------------------------------------------------------|--------------------------------------------|----------------------------------------------------------------------------------|
| [686. 重复叠加字符串匹配](https://leetcode.cn/problems/repeated-string-match/description/)                           | KMP算法，如果直到i>ab长度之和都未匹配成功，则一定无法匹配           | [我的提交](https://leetcode.cn/problems/repeated-string-match/submissions/490194187/) |
| [210. 课程表 II](https://leetcode.cn/problems/course-schedule-ii/description/)                                 | KMP算法                                      | [我的提交](https://leetcode.cn/problems/course-schedule-ii/submissions/484717117/)   |
| [2800. 包含三个字符串的最短字符串](https://leetcode.cn/problems/shortest-string-that-contains-three-strings/description/) | 先分解问题为 a,b 两个字符串的匹配，依次在a每个位置尝试匹配   ，最后合并答案 | [我的提交](https://leetcode.cn/problems/shortest-string-that-contains-three-strings/submissions/515538112/) |
| [2844. 生成特殊数字的最少操作](https://leetcode.cn/problems/minimum-operations-to-make-a-special-number/description/) | 4种模式从后向前匹配，注意特殊情况                          | [我的提交](https://leetcode.cn/problems/minimum-operations-to-make-a-special-number/submissions/533442146/) |

> 3> 字符串问题分类讨论
>
| 题目                                                                         | 说明                              | 实现                                                                            |
|----------------------------------------------------------------------------|---------------------------------|-------------------------------------------------------------------------------|
| [809. 情感丰富的文字](https://leetcode.cn/problems/expressive-words/description/) | 使用双指针匹配，分类讨论所有情况                | [我的提交](https://leetcode.cn/problems/expressive-words/submissions/490440072/) |
| [816. 模糊坐标](https://leetcode.cn/problems/ambiguous-coordinates/description/) | 分左右，分长度，再分首尾字母是否为0              | [我的提交](https://leetcode.cn/problems/ambiguous-coordinates/submissions/490614053/) |
| [955. 删列造序 II](https://leetcode.cn/problems/delete-columns-to-make-sorted-ii/) | 从左到右按列讨论，抽象成单个字符串的情形，使用所有字符串做限制 | [我的提交](https://leetcode.cn/problems/delete-columns-to-make-sorted-ii/submissions/490787364/) |
| [2380. 二进制字符串重新安排顺序需要的时间](https://leetcode.cn/problems/time-needed-to-rearrange-a-binary-string/description/) | max(当前1到最终位置的距离，前一个1到最终位置的时间+1) | [我的提交](https://leetcode.cn/problems/time-needed-to-rearrange-a-binary-string/submissions/509767919/) |
| [2981. 找出出现至少三次的最长特殊子字符串 I](https://leetcode.cn/problems/find-longest-special-substring-that-occurs-thrice-i/description/) | 只需存每个类型的前两长的长度和个数               | [我的提交](https://leetcode.cn/problems/find-longest-special-substring-that-occurs-thrice-i/submissions/513334558/) |
| [2825. 循环增长使字符串子序列等于另一个字符串](https://leetcode.cn/problems/make-string-a-subsequence-using-cyclic-increments/description/) | 字符串子序列问题用双指针 一定是没问题的   | [我的提交](https://leetcode.cn/problems/make-string-a-subsequence-using-cyclic-increments/submissions/515895823/) |
| [面试题 01.05. 一次编辑](https://leetcode.cn/problems/one-away-lcci/description/) | 字符串子序列问题用双指针 一定是没问题的   | [我的提交](https://leetcode.cn/problems/one-away-lcci/submissions/533994468/) |

### 3. 更多练习


### 4. 参考
1. [LeetCode（力扣）初级算法 字符串篇](https://blog.csdn.net/weixin_50803498/article/details/126528554) 
2. 总库：[tryHard](https://github.com/NOMADxzy/tryHard)