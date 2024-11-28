# 二叉树的序列化 和 反序列化
class TreeNode:
    def __init__(self, val, left=None, right=None):
        self.val, self.left, self.right = val, left, right


# 1
# 2      3
# 4 null 5 6



def enc(root: TreeNode) -> str:
    if root is None:
        return "null"
    ret = str(root.val) + "_"
    ret += f"({enc(root.left)})({enc(root.right)})"
    return ret


def dec(s: str) -> TreeNode:
    n = len(s)
    i = 0
    while i<n and s[i] != "_":
        i += 1
    if i == n:
        assert s == "null"
        return None
    ret = TreeNode(int(s[:i]))
    i += 1
    acc_left = 0
    pos1 = i+1
    while i<n:
        if s[i] == '(':
            acc_left += 1
        elif s[i] == ')':
            acc_left -= 1
        if acc_left == 0:
            ret.left = dec(s[pos1: i])
            break
        i += 1
    i += 1
    pos2 = i+1
    while i<n:
        if s[i] == '(':
            acc_left += 1
        elif s[i] == ')':
            acc_left -= 1
        if acc_left == 0:
            ret.right = dec(s[pos2: i])
            break
        i += 1
    return ret


root = TreeNode(1, TreeNode(2, TreeNode(4), None), TreeNode(3, TreeNode(5), TreeNode(6)))
enc_ret = enc(root)
dec_ret = dec(enc_ret)
print(enc_ret)
print(dec_ret.val)

