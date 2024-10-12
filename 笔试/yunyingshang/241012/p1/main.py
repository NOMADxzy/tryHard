# 找出链表的倒数第k个节点
class ListNode:
    def __init__(self, x):
        self.val, self.next = x, None


class Solution:
    def FindKthToTail(self, pHead: ListNode, k: int) -> ListNode:
        if k == 0:
            return None
        p, q = pHead, pHead
        for i in range(1, k):
            if q is None:
                break
            q = q.next
        if q is None:
            return None
        while q.next:
            p = p.next
            q = q.next
        return p
