class ListNode:
    def __init__(self, val: int, next):
        self.val, self.next = val, next


def sort(head: ListNode):
    tail_next = None
    while head.next != tail_next:

        p = head
        swap = False
        while p.next != tail_next:
            if p.val > p.next.val:
                p.val, p.next.val = p.next.val, p.val
                swap = True
            p = p.next
        if not swap:
            break


head = ListNode(3, ListNode(1, ListNode(5, ListNode(2, None))))
sort(head)
pass
