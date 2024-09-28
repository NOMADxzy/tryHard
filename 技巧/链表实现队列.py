
class ListNode:
    def __init__(self, val: int, next = None):
        self.val, self.next = val, next


class Queue:
    def __init__(self):
        self.head = ListNode(-1)
        self.tail = self.head

    def add(self, x: int):
        self.tail.next = ListNode(x)
        self.tail = self.tail.next

    def get(self):
        if self.head == self.tail:
            return None
        if self.head.next == self.tail:
            self.tail = self.head

        r = self.head.next.val
        self.head.next = self.head.next.next
        return r