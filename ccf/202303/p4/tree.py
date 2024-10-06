
class SegmentTree:
    def __init__(self, data, is_max=True):
        self.n = len(data)
        self.tree = [0] * (2 * self.n)
        self.is_max = is_max  # 是否求最大值
        self.build(data)

    def build(self, data):
        # 将原始数组放在线段树数组的后半部分
        for i in range(self.n):
            self.tree[self.n + i] = data[i]
        # 构建线段树的父节点
        for i in range(self.n - 1, 0, -1):
            if self.is_max:
                self.tree[i] = max(self.tree[2 * i], self.tree[2 * i + 1])
            else:
                self.tree[i] = min(self.tree[2 * i], self.tree[2 * i + 1])

    def update(self, pos, value):
        # 更新叶子节点
        pos += self.n
        self.tree[pos] = value
        # 递归更新父节点
        while pos > 1:
            pos //= 2
            if self.is_max:
                self.tree[pos] = max(self.tree[2 * pos], self.tree[2 * pos + 1])
            else:
                self.tree[pos] = min(self.tree[2 * pos], self.tree[2 * pos + 1])

    def query(self, l, r):
        # 查询区间 [l, r) 的最值
        l += self.n
        r += self.n
        result = float('-inf') if self.is_max else float('inf')
        while l < r:
            if l % 2 == 1:
                result = max(result, self.tree[l]) if self.is_max else min(result, self.tree[l])
                l += 1
            if r % 2 == 1:
                r -= 1
                result = max(result, self.tree[r]) if self.is_max else min(result, self.tree[r])
            l //= 2
            r //= 2
        return result

t = SegmentTree([1,2,3], True)
print(t.query(0, 2))