nums = list(map(int, input().split()))
target = int(input())
# -1 0 1 2 -1 -4
# 0

nums.sort()
ans = []


def dfs(remain: int, acc: [], pos: int):
    if len(acc) == 3:
        if remain == 0:
            ans.append(acc)
    else:
        i = pos
        while i < len(nums) and nums[i] <= remain:
            dfs(remain - nums[i], acc + [nums[i]], i + 1)
            while i + 1 < len(nums) and nums[i] == nums[i + 1]:
                i += 1
            i += 1


dfs(target, [], 0)
print(ans)
