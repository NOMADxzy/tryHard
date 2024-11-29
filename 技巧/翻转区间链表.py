class ListNode:
    def __init__(self, val, next=None):
        self.val, self.next = val, next


head = ListNode(1, ListNode(2, ListNode(3, ListNode(4, ListNode(5)))))
m, n = 0, 2


def reverseList(head: ListNode) -> (ListNode, ListNode):
    if head.next is None:
        return head, head
    next_head, next_tail = reverseList(head.next)
    next_tail.next = head
    head.next = None
    return next_head, head


tmp_head = ListNode(-1, head)
first_pre, second_next = tmp_head, None
p = tmp_head
idx = -1
while idx <= n:
    if idx == m - 1:
        first_pre = p
    elif idx == n:
        second_next = p.next
        p.next = None
    idx += 1
    p = p.next

new_mid_first, new_mid_second = reverseList(first_pre.next)
first_pre.next = new_mid_first
new_mid_second.next = second_next

ans = tmp_head.next
q = ans
while q is not None:
    print(q.val)
    q = q.next
