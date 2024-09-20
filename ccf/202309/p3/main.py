from typing import List, Tuple
from math import pow


class TreeNode:

    def __init__(self, val, left = None, right = None):
        self.val = val
        self.left = left
        self.right = right

    def is_symbol(self):
        return self.val == "*" or self.val == "+" or self.val == "-"


n, m = map(int, input().split())
elems = list(map(str, input().split()))

# 2 2
# x1 x1 x1 * x2 + *
# 1 2 3
# 2 3 4

stack = []
for e in elems:
    if e == "*" or e == "+" or e == "-":
        new_node = TreeNode(e, stack[-2], stack[-1])
        stack = stack[:-2]
    else:
        new_node = TreeNode(e)
    stack.append(new_node)

assert len(stack) == 1
root = stack[0]


def cpt_f(root: TreeNode, target: str, vals: List) -> List[Tuple[int, int]]:
    if not root.is_symbol():
        if root.val == target:
            return [(1,  0)]
        else:
            return [(0, 0)]
    else:
        ret = []
        tmp_map = {}
        left_f_ret = cpt_f(root.left, target, vals)
        right_f_ret = cpt_f(root.right, target, vals)
        if root.val == "*":
            left_ret = cpt(root.left, target, vals)
            right_ret = cpt(root.right, target, vals)

            for x in left_f_ret:
                for y in right_ret:
                    ex = x[1] + y[1]
                    if ex not in tmp_map:
                        tmp_map[ex] = 0
                    tmp_map[ex] += x[0]*y[0]
            for x in left_ret:
                for y in right_f_ret:
                    ex = x[1] + y[1]
                    if ex not in tmp_map:
                        tmp_map[ex] = 0
                    tmp_map[ex] += x[0]*y[0]
        else:
            right_symbol = -1 if root.val == "-" else 1
            for x in left_f_ret:
                if x[1] not in tmp_map:
                    tmp_map[x[1]] = 0
                tmp_map[x[1]] += x[0]

            for y in right_f_ret:
                if y[1] not in tmp_map:
                    tmp_map[y[1]] = 0
                tmp_map[y[1]] += y[0] * right_symbol
        for k,v in tmp_map.items():
            ret.append((v, k))
        return ret


def cpt(root: TreeNode, target: str, vals: List) -> List[Tuple[int, int]]:
    if not root.is_symbol():
        if root.val == target:
            return [(1,  1)]
        else:
            if root.val[:1] == "x":
                return [(vals[int(root.val[1:]) - 1], 0)]
            return [(int(root.val), 0)]
    else:
        ret = []
        left_ret = cpt(root.left, target, vals)
        right_ret = cpt(root.right, target, vals)
        tmp_map = {}
        if root.val == "*":
            for x in left_ret:
                for y in right_ret:
                    ex = x[1] + y[1]
                    if ex not in tmp_map:
                        tmp_map[ex] = 0
                    tmp_map[ex] += x[0]*y[0]
        else :
            right_symbol = -1 if root.val == "-" else 1
            for x in left_ret:
                if x[1] not in tmp_map:
                    tmp_map[x[1]] = 0
                tmp_map[x[1]] += x[0]

            for y in right_ret:
                if y[1] not in tmp_map:
                    tmp_map[y[1]] = 0
                tmp_map[y[1]] += y[0] * right_symbol
        for k,v in tmp_map.items():
            ret.append((v, k))
        return ret


MOD = 1000000007

for _ in range(m):
    row_elems = list(map(int, input().split()))
    idx = row_elems[0]
    vals = row_elems[1:]

    expresstion = cpt_f(root, f'x{idx}', vals)
    ans = 0
    for co, ex in expresstion:
        ans = (ans + co * pow(vals[idx-1], ex) % MOD) % MOD
    print(int(ans))
