## 多路归并

### 1. 概念

一、多路归并算法的由来

        假定现在有一包含大量整数的文本文件存放于磁盘中，其文件大小为10GB，而本机内存只有4GB。此时若我们要对该文件中的所有整数进行升序排序，肯定不能直接将文件中的所有数据一次性读入内存中，再使用快速、归并等排序算法对这么大规模的整数进行排序。
    
        好像陷入了难题？ 我们不妨换一个思路，为何不将10GB大文件拆分为10个1GB的小文件呢？ 逐个对10个文件进行排序后，再将其写入磁盘中，此时就得到了10份已排序后的临时文件。
    
        每一份文件都是一个升序序列，这时问题就转换为如何合并这10路升序序列为1路升序序列。正因为待合并的数据路数比较多，所以才有了多路归并这一说法。
    
        还是有些抽象？那不妨举个具体的例子来瞅瞅，假定有nums1、num2、nums3三路升序序列，先打算将他们合并为一路升序序列：

nums1 = [1, 2, 3]

nums2 = [2, 4, 6]

nums3 = [3, 5 ,8] =》 合并结果: res = [1, 2, 2, 3, 3, 4, 5, 6, 8]

二、多路归并算法的最简编程模型

        现在我们将刚才提到的合并k路升序序列问题，转化为具体的C++代码，也即是建立最简单的多路归并算法编程模型。以k = 3，即3路为例：
    
        为了便于理解，我们将三路序列存储为矩阵的形式（如上图），横坐标x的取值范围为0,1,2代表有三列，纵坐标的取值范围也为0,1,2代表有三行。第一行(y=0)是nums1，第二行(y=1)是nums2，第三行（y=2）是nums3。

多路归并算法的基本思想如下：

1. 首先建立一个小顶堆；

2. 将每一路的最小元素（即第1列元素）都加入小顶堆中，此时堆顶就是k路中全局的最小值；

3. 将堆顶元素弹出，并将堆顶元素所在数组的下一个元素加入堆中。

4. 重复第2)和第3)步，直至每一路数据都读取结束。

   ![img](https://gitee.com/xu_zuyun/picgo/raw/master/img/e113f5e6718845a2a58de7f3a64b57e0.png)


### 2. 解题技巧（我的总结）

>  多路归并
> 
| 题目                                                                     | 说明                                | 实现                                                                            |
|------------------------------------------------------------------------|-----------------------------------|-------------------------------------------------------------------------------|
| [313. 超级丑数](https://leetcode.cn/problems/super-ugly-number/description/) | 多路归并                              | [我的提交](https://leetcode.cn/problems/super-ugly-number/submissions/468697927/) |
| [786. 第 K 个最小的质数分数](https://leetcode.cn/problems/k-th-smallest-prime-fraction/description/) | 多路归并                              | [我的提交](https://leetcode.cn/problems/k-th-smallest-prime-fraction/submissions/468733547/) |
| [1508. 子数组和排序后的区间和](https://leetcode.cn/problems/range-sum-of-sorted-subarray-sums/description/) | 多路归并                              | [我的提交](https://leetcode.cn/problems/range-sum-of-sorted-subarray-sums/submissions/468741186/) |
| [LCR 061. 查找和最小的 K 对数字](https://leetcode.cn/problems/qn8gGX/description/) | ptr记录每个nums1中的元素匹配到nums2中的序号，m路归并 | [我的提交](https://leetcode.cn/problems/qn8gGX/submissions/538547705/) |

### 3. 更多练习


### 4. 参考
1. [多路归并算法从理论到应用（易懂）](https://blog.csdn.net/a574780196/article/details/122646309) 
2. 总库：[tryHard](https://github.com/NOMADxzy/tryHard)