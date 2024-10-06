class SegmentTree:

    def __init__(self, data):
        n = len(data)
        self.n = n
        self.tree = [ [0,0] for _ in range(2*n)]

        for i in range(0, len(data)):
            self.tree[n + i][0], self.tree[n + i][1]  = data[i], i

        for i in range(n-1, -1, -1):
            if self.tree[i*2][0] >= self.tree[i*2+1][0]:
                self.tree[i][0], self.tree[i][1]  = self.tree[i*2][0], self.tree[i*2][1]
            else:
                self.tree[i][0], self.tree[i][1] = self.tree[i * 2+1][0], self.tree[i * 2+1][1]

    def update(self, i, delta):
        pos = self.n + i
        self.tree[pos][0] = self.query(i, i+1)[0] + delta
        while pos>0:
            pos //= 2
            left = pos*2
            right = pos*2 + 1
            if self.tree[left][0] >= self.tree[right][0]:
                self.tree[pos][0], self.tree[pos][1] = self.tree[left]
            else:
                self.tree[pos][0], self.tree[pos][1] = self.tree[right]

    def query(self, l, r):
        l += self.n
        r += self.n

        cur_max_idx = l
        while l<r:
            if l%2==1:
                if self.tree[l][0] > self.tree[cur_max_idx][0] or self.tree[l][0] == self.tree[cur_max_idx][0] and self.tree[l][1] < self.tree[cur_max_idx][1]:
                    cur_max_idx = l
                l += 1
            if r%2==1:
                r -= 1
                if self.tree[r][0] > self.tree[cur_max_idx][0] or self.tree[r][0] == self.tree[cur_max_idx][0] and self.tree[r][1] < self.tree[cur_max_idx][1]:
                    cur_max_idx = r
            l //= 2
            r //= 2
        return self.tree[cur_max_idx] if self.tree[cur_max_idx][0]>0 else [0, -1]

d = {'key': []}F

for _ in range(100):
    d['key'].append(0)

print(d)