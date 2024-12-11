from typing import List
from math import inf


class SegmentTree:
    def __init__(self, data: List):
        n = len(data)
        self.n = n
        self.tree = [inf] * (2 * n)

        for i in range(n):
            self.tree[i + n] = data[i]

        for i in range(n - 1, 0, -1):
            l, r = 2 * i, 2 * i + 1
            self.tree[i] = min(self.tree[l], self.tree[r])

    def update(self, pos, new_val):
        p = pos + self.n
        self.tree[p] = new_val
        while p > 0:
            p //= 2
            self.tree[p] = min(self.tree[p], new_val)

    def query(self, l, r):
        l += self.n
        r += self.n
        cur_min = inf
        while r > l:
            if l % 2 == 1:
                cur_min = min(cur_min, self.tree[l])
                l += 1
            if r % 2 == 0:
                cur_min = min(cur_min, self.tree[r])
                r -= 1
            l //= 2
            r //= 2
        return min(self.tree[l], cur_min)
