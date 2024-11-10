# arr1中从小到大取数入栈，每次从栈中取数得分为arr2[栈元素数量-1]*栈顶元素，问所有中操作的得分之和

from copy import deepcopy


n = int(input())
arr1 = list(map(int, input().split()))
arr2 = list(map(int, input().split()))

arr1.sort()
ans = 0

# 1 2 3 4
# 4 4 4 4


def dfs(cur_stack, cur_arr):
    global ans
    global arr2
    if len(cur_stack) > 0:
        new_stack = deepcopy(cur_stack)
        new_arr1 = deepcopy(cur_arr)
        ans += arr2[len(new_stack)-1] * new_stack[-1]
        new_stack = new_stack[:len(new_stack)-1]
        dfs(new_stack, new_arr1)
    if len(cur_arr) == 0:
        return
    new_stack2 = deepcopy(cur_stack)
    new_arr2 = deepcopy(cur_arr)
    x = cur_arr[0]
    new_arr2 = new_arr2[1:]
    new_stack2.append(x)
    dfs(new_stack2, new_arr2)

dfs([],arr1)
print(ans)