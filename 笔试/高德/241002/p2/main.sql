# 找出所有工资在平均工资以上的员工的名字和部门

select name, department from employees where salary > (select AVG(salary) from employees) order by salary DESC;