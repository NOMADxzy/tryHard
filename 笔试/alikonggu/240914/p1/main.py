# 给n个不重复的数字
# q次操作 (opt, x)，操作0 表示把x移动到队头，操作1 表示把x移动到队尾
# 输出最终的队列

n, q = map(int, input().split())
arr = list(map(int, input().split()))

class ListNode:
    def __init__(self, val, pre=None, next=None):
        self.val = val
        self.pre = pre
        self.next = next

head, tail = ListNode(-1), ListNode(-1)
head.next = tail
tail.pre = head

def print_list(h: ListNode):
    tmp = []
    while h != tail:
        tmp.append(h.val)
        h = h.next
    print(" ".join(map(str, tmp)))

def move_to_end(node:ListNode):
    if node.pre == None and node.next == None:
        node.pre = tail.pre
        tail.pre.next = node
        node.next = tail
        tail.pre = node
        return

    if node.next != tail:
        node.next.pre = node.pre
        node.pre.next = node.next
        node.pre = tail.pre
        tail.pre.next = node
        node.next = tail
        tail.pre = node
        return

def move_to_start(node:ListNode):
    if node.pre == None and node.next == None:
        node.next = head.next
        head.next.pre = node
        head.next = node
        node.pre = head
        return

    if node.pre != head:
        node.next.pre = node.pre
        node.pre.next = node.next
        node.next = head.next
        head.next.pre = node
        head.next = node
        node.pre = head
        return

idx_map = {}
for v in arr:
    new_node = ListNode(v)
    idx_map[v] = new_node
    move_to_end(new_node)


for _ in range(q):
    opt, no = map(int, input().split())
    if opt == 0:
        move_to_start(idx_map[no])
    else:
        move_to_end(idx_map[no])
    # print_list(head.next)



print_list(head.next)
