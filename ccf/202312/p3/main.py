from copy import deepcopy
from math import inf

from networkx import min_weight_matching

n, m = map(int, input().split())
probs = list(map(int, input().split()))
fathers = list(map(int, input().split()))

class Node:
    def __init__(self, prob: int, id: int, father=None):
        self.prob = prob
        self.father = father
        self.total_prob = -1
        self.id = id
        self.childrens = []

nodes = []
for i, prob in enumerate(probs):
    nodes.append(Node(prob, i+1))

for i,f in enumerate(fathers):
    nodes[f-1].childrens.append(nodes[i+1])
    nodes[i+1].father = nodes[f-1]

def build(node: Node) -> int:
    if node.total_prob < 0:
        node.total_prob = node.prob
        for child in node.childrens:
            node.total_prob += build(child)
    return node.total_prob



def remove(node: Node):
    p = node
    val = p.total_prob
    node.total_prob = 0
    while p.father:
        p = p.father
        p.total_prob -= val


def belong_to(node1: Node, node2: Node)->bool:
    while node1:
        if node1.id == node2.id:
            return True
        node1 = node1.father
    return False



build(nodes[0])

for _ in range(m):
    target = nodes[int(input())-1]
    ans = []

    root = deepcopy(nodes[0])
    while True:
        min_w = inf
        q_node = Node(-1, -1)
        queue = [root]
        remain_cnt = 0
        while len(queue) > 0:
            e = queue[0]
            queue = queue[1:]
            if e.total_prob > 0:
                remain_cnt += 1
                cur_w = abs(root.total_prob - e.total_prob - e.total_prob)
                if cur_w < min_w or cur_w == min_w and e.id < q_node.id:
                    min_w = cur_w
                    q_node = e
                for p in e.childrens:
                    if p.total_prob > 0:
                        queue.append(p)

        if remain_cnt<=1:
            break
        ans.append(q_node.id)
        if belong_to(target, q_node):
            root = q_node
        else:
            remove(q_node)

    print(" ".join(map(str, ans)))




