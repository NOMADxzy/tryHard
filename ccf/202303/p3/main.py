from typing import Dict
# 2
# 1 2 1 2 2 3
# 2 2 2 3 3 1
# 4
# 1:2
# 3~1
# &(1:2)(2:3)
# |(1:2)(3:1)

n = int(input())
features = [{} for _ in range(n)]
dn_list = []

for i in range(n):
    arr = list(map(int, input().split()))
    dn_list.append(arr[0])
    for j in range(2, len(arr), 2):
        features[i][arr[j]] = arr[j+1]


class TreeNode:

    def __init__(self, symb: str, expr=None, left=None, right=None):
        self.symb = symb
        self.expr = expr
        self.left = left
        self.right = right

def build(s: str) -> TreeNode:
    if s[0] == '|' or s[0] == '&':
        left_bs, right_bs = [], []
        cnt = 0
        for i in range(len(s)):
            if s[i] == '(':
                cnt += 1
                if cnt == 1:
                    left_bs.append(i)
            elif s[i] == ')':
                cnt -= 1
                if cnt == 0:
                    right_bs.append(i)

        left_s = s[left_bs[0]+1: right_bs[0]]
        right_s = s[left_bs[1]+1: right_bs[1]]
        return TreeNode(s[0], None, build(left_s), build(right_s))
    else:
        mid = 0
        for i in range(len(s)):
            if s[i] == ':' or s[i] == '~':
                mid = i
                break
        expr = (int(s[:mid]), s[mid], int(s[mid+1:]))
        return TreeNode("single", expr)

def find(root: TreeNode)-> Dict:
    ret = {}
    if root.symb == "single":
        for i in range(n):
            query_res = None
            if root.expr[0] in features[i]:
                query_res = features[i][root.expr[0]]
            if not query_res is None:
                if root.expr[1] == ':' and root.expr[2] == query_res or root.expr[1] == '~' and root.expr[2] != query_res:
                    ret[dn_list[i]] = True
    else:
        left_ret, right_ret = find(root.left), find(root.right)
        if root.symb == '|':
            for dn,_ in left_ret.items():
                ret[dn] = True
            for dn,_ in right_ret.items():
                ret[dn] = True
        elif root.symb == '&':
            for dn,_ in left_ret.items():
                if dn in right_ret:
                    ret[dn] = True
    return ret



m = int(input())
for _ in range(m):
    s = input()
    root = build(s)
    ans = sorted(find(root))
    print(" ".join(map(str, ans)))



