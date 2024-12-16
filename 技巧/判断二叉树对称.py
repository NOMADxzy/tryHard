from typing import List


class TreeNode:
    def __init__(self, val, left=None, right=None):
        self.val, self.left, self.right = val, left, right


root = TreeNode(0, TreeNode(2, TreeNode(3), TreeNode(4)), TreeNode(2, TreeNode(4), TreeNode(3)))
queue = [root]
ans = True


def check(q: List) -> bool:
    n = len(q)
    for i in range(n // 2):
        l, r = i, n - 1 - i
        if q[l] is None:
            if q[r] is not None:
                return False
        else:
            if q[r] is None or q[r].val != q[l].val:
                return False
    return True


while len(queue) > 0:
    if not check(queue):
        ans = False
        break
    new_queue = []
    has_next_layer = False
    while len(queue) > 0:
        e = queue[0]
        queue = queue[1:]
        if e is None:
            new_queue.append(None)
            new_queue.append(None)
        else:
            new_queue.append(e.left)
            new_queue.append(e.right)

    if not has_next_layer:
        break
    queue = new_queue

print(ans)
