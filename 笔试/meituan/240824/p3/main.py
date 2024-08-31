
n, m = map(int, input().split())
dates = list(map(int, input().split()))
types = list(map(int, input().split()))


class SegmentTree():

    def __init__(self, data):
        n = 1
        while n<len(data):
            n *= 2
        self.n = n
        self.tree = [-1 for _ in range(2*n-1)]
        self.build(data)

    def build(self, data):
        for i in range(len(data)):
            self.tree[self.n-1+i] = data[i]

        for i in range(self.n-2, -1, -1):
            max_val = -1
            if 2*i + 1 <= self.n-1+len(data):
                max_val = max(max_val, data[2*i+1 - (self.n-1)])
            if 2*i + 2 < self.n-1+len(data):
                max_val = max(max_val, data[2*i+2 - (self.n-1)])
            self.tree[i] = max_val

    def update(self, i, new_val):
        p = i + self.n - 1
        self.tree[p] = new_val
        while p>0:
            p = (p-1)//2
            self.tree[p] = max(self.tree[p*2+1], self.tree[p*2+2])

    def query(self, l, r):
        p, q = self.n-1+l, self.n-1+r
        while p<q:
            p = (p-1)//2
            q = (q-1)//2
        max_val = self.tree[p]
        return max_val, self.query_place(p) if max_val>=0 else -1

    def query_place(self, i):
        while i<self.n-1:
            l,r = 2*i+1, 2*i+2
            if self.tree[l] < self.tree[r]:
                i = r
            else:
                i = l
        return i - (self.n - 1)


arr0, arr1 = [-1 for _ in range(m)], [-1 for _ in range(m)]
for i in range(m):
    if types[i] == 0:
        arr0[i] = dates[i]
    else:
        arr1[i] = dates[i]

trees = [SegmentTree(arr0), SegmentTree(arr1)]

for _ in range(n):
    l, r, t, k = map(int, input().split())
    tree = trees[t]
    idxs = []
    max_val, max_idx = tree.query(l, r)
    while max_idx >= 0 and len(idxs) < k:
        idxs.append(max_idx)
        tree.update(max_idx, -1)
        max_val, max_idx = tree.query(l, r)

    print(" ".join(map(str, idxs))) if len(idxs) == k else print(-1)



