

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
        p = i+self.n-1
        self.tree[p] = new_val
        while p>0:
            p = (p-1)//2
            self.tree[p] = max(self.tree[p*2+1], self.tree[p*2+2])

    def query(self, l, r):
        p, q = self.n-1+l, self.n-1+r
        while p<q:
            p = (p-1)//2
            q = (q-1)//2
        return self.tree[p]


data = [i for i in range(1, 5)]
t = SegmentTree(data)

print(t.query(0, 3))
t.update(0, 10)
print(t.query(0, 1))
print(t.query(2, 3))