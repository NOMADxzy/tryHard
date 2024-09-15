from math import inf
from typing import List
n, q = map(int, input().split())

idx_map = {}
exist_map = {}
addr_list = []

def add_addr(addr):
    if addr not in exist_map:
        addr_list.append(addr)
        exist_map[addr] = True

opts = []
for _ in range(q):
    splits = input().split()
    opt = splits[0]
    if opt == '1':
        id, start, end = int(splits[1]), splits[2], splits[3]
        add_addr(start)
        add_addr(end)
        opts.append((opt, (id, start, end)))

    elif opt == '2':
        address = splits[1]
        add_addr(address)
        opts.append((opt, address))
    elif opt == '3':
        start, end = splits[1], splits[2]
        add_addr(start)
        add_addr(end)
        opts.append((opt, (start, end)))

addr_list.sort()
m = len(addr_list)
for i,addr in enumerate(addr_list):
    idx_map[addr] = i+1

# tree = [0 for _ in range(m+1)]

# def low_bit(x):
#     return x & (-x)
#
# def update(x, v):
#     while x<len(tree):
#         tree[x] += v
#         x += low_bit(x)
#
# def query(x):
#     total = 0
#     while x>0:
#         total += tree[x]
#         x -= low_bit(x)

class SegmentTree:
    def __init__(self, data: List, is_max: bool):
        n = len(data)
        self.n = n
        self.is_max = is_max
        init_val = inf
        if is_max:
           init_val = -init_val
        self.tree = [init_val for _ in range(2*n)]
        for i in range(len(data)):
            self.tree[i+n] = data[i]
        for i in range(n-1, -1, -1):
            if is_max:
                val = max(self.tree[2*i], self.tree[2*i+1])
            else:
                val = min(self.tree[2 * i], self.tree[2 * i + 1])
            self.tree[i] = val

    def update(self, pos, new_val):
        p = pos + self.n
        self.tree[p] = new_val
        while p>0:
            p //= 2
            if self.is_max:
                self.tree[p] = max(self.tree[p], new_val)
            else:
                self.tree[p] = min(self.tree[p], new_val)

    def query(self, l, r):
        l += self.n
        r += self.n
        cur_val = -inf if self.is_max else inf
        while l<r:
            if l%2 == 1:
                if self.is_max:
                    cur_val = max(cur_val, self.tree[l])
                else:
                    cur_val = min(cur_val, self.tree[l])
                l += 1
            if r%2 == 1:
                r -= 1
                if self.is_max:
                    cur_val = max(cur_val, self.tree[r])
                else:
                    cur_val = min(cur_val, self.tree[r])
            l //= 2
            r //= 2

        return cur_val



min_tree = SegmentTree([inf for i in range(m)], False)
max_tree = SegmentTree([-inf for _ in range(m)], True)

for opt, param in opts:
    if opt == '1':
        id, start, end = param
        start_idx, end_idx = idx_map[start], idx_map[end]
        max_id = max_tree.query(start_idx, end_idx)
        min_id = min_tree.query(start_idx, end_idx)
        if max_id == min_id and min_id == id or max_id == -inf:
            ok = False
            for i in range(start_idx, end_idx):
                if max_tree.query(i,i+1) == id:
                    continue
                max_tree.update(i, id)
                min_tree.update(i, id)
                ok = True

            print("YES" if ok else "NO")
        else:
            print("NO")

    elif opt == '2':
        addr = param
        idx = idx_map[addr]
        cur_id = max_tree.query(idx, idx+1)
        print(cur_id if cur_id != -inf else 0)
    else:
        start, end = param
        start_idx, end_idx = idx_map[start], idx_map[end]
        should_id = max_tree.query(start_idx, start_idx+1)
        if should_id == -inf:
            print(0)
        else:
            ok = True
            for i in range(start_idx+1, end_idx):
                if max_tree.query(i, i+1) != should_id:
                    ok = False
                    break

            print(should_id if ok else 0)
