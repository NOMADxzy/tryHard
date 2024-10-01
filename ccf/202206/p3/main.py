from copy import deepcopy
n, m, q = map(int, input().split())
role_opt_dict = {}

for _ in range(n):
    splits = input().split()
    role = splits[0]
    ptr = 2
    num1 = int(splits[ptr-1])
    opts = splits[ptr: ptr + int(splits[ptr-1])]
    ptr += len(opts) + 1
    num2 = int(splits[ptr - 1])
    r_types = splits[ptr: ptr + int(splits[ptr-1])]
    ptr += len(r_types) + 1
    num3 = int(splits[ptr - 1])
    r_names = splits[ptr: ptr + int(splits[ptr-1])]

    assert len(opts) == num1 and len(r_types) == num2 and len(r_names) == num3


    tmp = [set(opts), set(r_types), None]
    if '*' in tmp[0]:
        tmp[0] = None
    if '*' in tmp[1]:
        tmp[1] = None
    if len(r_names) > 0:
        tmp[2] = set(r_names)
    role_opt_dict[role] = tmp

u_roles = {}
g_roles = {}

def add_role(roles, name, role):
    if name not in roles:
        roles[name] = {role}
    else:
        roles[name].add(role)

for _ in range(m):
    splits = input().split()
    role = splits[0]
    for i in range(2, len(splits), 2):
        x_type, x_name = splits[i], splits[i+1]
        if x_type == 'u':
            add_role(u_roles, x_name, role)
        else:
            add_role(g_roles, x_name, role)

for _ in range(q):
    splits = input().split()
    u_name = splits[0]

    groups = splits[2: 2 + int(splits[1])]
    opt, r_type, r_name = splits[-3:]
    role_set = set(u_roles.get(u_name, set()))
    for g_name in groups:
        for role in g_roles.get(g_name, set()):
            role_set.add(role)

    msg = '0'
    for role in role_set:
        tmp = role_opt_dict[role]
        if (tmp[0] is None or opt in tmp[0]) and (tmp[1] is None or r_type in tmp[1]) and (tmp[2] is None or r_name in tmp[2]):
            msg = '1'
            break

    print(msg)
