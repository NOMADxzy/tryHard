# Floyd-Warshall算法的Python实现
INF = float('inf')


def floyd_warshall(graph):
    # 获取图中节点的数量
    n = len(graph)

    # 初始化距离矩阵，复制原图
    dist = [[graph[i][j] for j in range(n)] for i in range(n)]

    # 使用Floyd-Warshall算法更新距离矩阵
    for k in range(n):
        for i in range(n):
            for j in range(n):
                # 更新dist[i][j]为经过节点k的路径和直接路径的最短者
                dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])

    return dist


# 示例：带权重的有向图
graph = [
    [0, 3, INF, 7],
    [8, 0, 2, INF],
    [5, INF, 0, 1],
    [2, INF, INF, 0]
]

# 运行Floyd-Warshall算法
result = floyd_warshall(graph)

# 输出最短路径矩阵
for row in result:
    print(row)