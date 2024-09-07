
class TreeNode:
    def __init__(self, val, left = None, right = None):
        self.val = val
        self.left = left
        self.right = right

dp = [{}, {}]
def dfs(root: TreeNode, is_child: bool)->int:
    ret = 0
    if root.left is None and root.right is None:
        ret = 0
    elif root.left is None:
        ret = dfs(root.right, False)
    elif root.right is None:
        ret = dfs(root.left, False)
    else:
        ret = dfs(root.left, False) + dfs(root.right, False)
        if not is_child:
            ret = max(ret, dfs(root.left, True) + dfs(root.right, True) + 1)

    dp[1 if is_child else 0][root] = ret
    return ret

root = TreeNode(1, TreeNode(2, TreeNode(4), TreeNode(5)),  TreeNode(3, TreeNode(6), TreeNode(7)))
print(dfs(root, False))

