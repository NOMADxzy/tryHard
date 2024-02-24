## 拓扑排序

### 1. 概念
- 对一个有向无环图(Directed Acyclic Graph简称DAG)G进行拓扑排序，是将G中所有顶点排成一个线性序列，使得图中任意一对顶点u和v，若边<u,v>∈E(G)，则u在线性序列中出现在v之前。通常，这样的线性序列称为满足拓扑次序(Topological Order)的序列，简称拓扑序列。
- 简单的说，由某个集合上的一个偏序得到该集合上的一个全序，这个操作称之为拓扑排序。

过程：
```markdown
在一个有向图中，对所有的节点进行排序，要求没有一个节点指向它前面的节点。
先统计所有节点的入度，对于入度为0的节点就可以分离出来，然后把这个节点指向的节点的入度减一。
一直做改操作，直到所有的节点都被分离出来。
如果最后不存在入度为0的节点，但还存在节点，那就说明有环，不存在拓扑排序，也就是很多题目的无解的情况。
```
![](https://gitee.com/xu_zuyun/picgo/raw/master/img/20210421164716512.png)


算法：
```markdown
第一种方式

是遍历整个图中的顶点，找出入度为0的顶点，然后标记删除该顶点，更新相关顶点的入度，由于图中有V个顶点，每次找出入度为0的顶点后会更新相关顶点的入度，因此下一次又要重新扫描图中所有的顶点。故时间复杂度为O(V^2)

问题：由于删除入度为0的顶点时，只会更新与它邻接的顶点的入度，即只会影响与之邻接的顶点。但是上面的方式却遍历了图中所有的顶点的入度。
第二种方式（更优的算法）

先将入度为0的顶点放在栈或者队列中。当队列不空时，删除一个顶点v，然后更新与顶点v邻接的顶点的入度。只要有一个顶点的入度降为0，则将之入队列。此时，拓扑排序就是顶点出队的顺序。该算法的时间复杂度为O（V+E）
```

### 2. 解题技巧（我的总结）

> 1> 拓扑排序
> 
| 题目                                                                          | 说明   | 实现                                                                            |
|-----------------------------------------------------------------------------|------|-------------------------------------------------------------------------------|
| [210. 课程表 II](https://leetcode.cn/problems/course-schedule-ii/description/) | 拓扑排序 | [我的提交](https://leetcode.cn/problems/course-schedule-ii/submissions/484717117/) |
| [310. 最小高度树](https://leetcode.cn/problems/minimum-height-trees/description/) | 拓扑排序 | [我的提交](https://leetcode.cn/problems/minimum-height-trees/submissions/484940211/) |
| [684. 冗余连接](https://leetcode.cn/problems/redundant-connection/description/) | 拓扑排序 | [我的提交](https://leetcode.cn/problems/redundant-connection/submissions/490074846/) |


### 3. 更多练习


### 4. 参考
1. [拓扑排序及算法实现](https://blog.csdn.net/ShyTan/article/details/115962715)
2. 总库：[tryHard](https://github.com/NOMADxzy/tryHard)