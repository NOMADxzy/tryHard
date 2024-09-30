from math import inf

n = int(input())
risk_start_day_map = {} # 这个地区 风险开始时间
risk_over_day_map = {} # 这个地区 风险结束时间
latest_visit_day = {} # 用户u访问r的最近一天


for cur_day in range(n):
    splits = list(map(int, input().split()))
    risk_n, m, risks = splits[0], splits[1], splits[2:]
    for risk_r in risks:
        risk_start_day, risk_over_day = risk_start_day_map.get(risk_r, -inf), risk_over_day_map.get(risk_r, -inf)
        if cur_day <= risk_over_day:
            risk_over_day_map[risk_r] = cur_day + 7
        else:
            risk_start_day_map[risk_r] = cur_day
            risk_over_day_map[risk_r] = cur_day + 7

    risk_u = set()

    for _ in range(m):
        d, u, r = map(int, input().split())
        if r not in latest_visit_day:
            latest_visit_day[r] = {}
        val = max(latest_visit_day[r].get(u, -inf), d)
        latest_visit_day[r][u] = val
        if type(val) == float or val + 7 <= cur_day:
            latest_visit_day[r].pop(u)


    for r, tmp_dict in latest_visit_day.items():
        day_start = risk_start_day_map.get(r, -inf)
        day_end = risk_over_day_map.get(r,-inf)
        if day_end <= cur_day:
            continue
        for u, latest_visit_day_val in tmp_dict.items():
            if latest_visit_day_val > cur_day - 7 and day_start <= latest_visit_day_val:
                risk_u.add(u)

    ans = list(risk_u)
    ans.sort()
    print(f'{cur_day} {" ".join(map(str, ans))}')
