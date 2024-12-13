
# 1
# 3 5
# XOR 2 I1 I2
# XOR 2 O1 I3
# AND 2 O1 I3
# AND 2 I1 I2
# OR 2 O3 O4
# 4
# 0 1 1
# 1 0 1
# 1 1 1
# 0 0 0
# 2 5 2
# 2 5 2
# 2 5 2
# 2 5 2

q = int(input())



class Node:
    def __init__(self, typ="IN", inputs=None):
        self.typ, self.inputs = typ, inputs

    def val(self) -> bool:
        if self.typ == "IN":
            assert len(self.inputs) == 1 and isinstance(self.inputs[0], bool)
            return self.inputs[0]
        elif self.typ == "AND":
            ret = True
            for v in self.inputs:
                ret &= v if isinstance(v, bool) else v.val()
            return ret
        elif self.typ == "OR":
            ret = False
            for v in self.inputs:
                ret |= v if isinstance(v, bool) else v.val()
            return ret
        elif self.typ == "XOR":
            ret = False
            for v in self.inputs:
                ret ^= v if isinstance(v, bool) else v.val()
            return ret
        elif self.typ == "NOT":
            assert len(self.inputs) == 1
            v = self.inputs[0]
            return not v if isinstance(v, bool) else v.val()

class Pres:
    def __init__(self, typ):
        self.typ, self.pre_ins, self.pre_outs = typ, [], []


for _ in range(q):
    m, n = map(int, input().split())

    ins = [Node(inputs=[False]) for _ in range(m)]
    outs = [Node() for _ in range(n)]
    next_out_ids = [[] for _ in range(n)]
    in_degrees = [0]*n
    pres_list = []

    for o_idx in range(n):
        input_strs = input().split()
        cur_opt = input_strs[0]
        cur_input_cnt = int(input_strs[1])
        pres = Pres(cur_opt)
        for v in input_strs[2:]:
            idx = int(v[1:])
            if v[0] == "I":
                pres.pre_ins.append(idx)
            else:
                pres.pre_outs.append(idx)
                in_degrees[o_idx] += 1
                next_out_ids[idx].append(o_idx)
        pres_list.append(pres)
    queue = []
    out_idxs = []
    for i in range(n):
        if in_degrees[i] == 0:
            queue.append(i)
        if len(next_out_ids[i]) == 0:
            out_idxs.append(i)

    flag = 0
    while len(queue)>0:
        next_queue = []
        while len(queue) > 0:
            o_idx = queue[0]
            queue = queue[1:]
            pres = pres_list[o_idx]
            all_pre_nodes = []
            for pre_idx in pres.pre_ins:
                all_pre_nodes.append(ins[pre_idx])
            for pre_idx in pres.pre_outs:
                all_pre_nodes.append(outs[pre_idx])

            outs[o_idx] = Node(pres.typ, all_pre_nodes)
            flag += 1
            for next_id in next_out_ids[o_idx]:
                in_degrees[next_id] -= 1
                if in_degrees[next_id] == 0:
                    next_queue.append(next_id)
        queue = next_queue
    last_out_nodes = [outs[i] for i in out_idxs]
    if flag < n:
        print("")
    else:

        out_cnts = int(input())
        for _ in range(out_cnts):
            input_ints = list(map(int, input().split()))
            for i in range(m):
                ins[i].inputs = [input_ints[i] == 1]

        ans = [last_out_node.val() for last_out_node in last_out_nodes]






