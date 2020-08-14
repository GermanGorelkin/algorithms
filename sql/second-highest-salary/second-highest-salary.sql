/*
Write a SQL query to get the second highest salary from the Employee table.

+----+--------+
| Id | Salary |
+----+--------+
| 1  | 100    |
| 2  | 200    |
| 3  | 300    |
+----+--------+
For example, given the above Employee table, the query should return 200 as the second highest salary. If there is no second highest salary, then the query should return null.

+---------------------+
| SecondHighestSalary |
+---------------------+
| 200                 |
+---------------------+
 */

 declare @Employee table(Id int, Salary int)
insert into @Employee (Id, Salary) values ('1', '100')
insert into @Employee (Id, Salary) values ('2', '200')
insert into @Employee (Id, Salary) values ('3', '300')

-----------
select (select distinct Salary
from @Employee
order by Salary desc
OFFSET 1 rows
FETCH FIRST 1 ROWS ONLY) as 'SecondHighestSalary'

--------------
SELECT MAX(Salary) as 'SecondHighestSalary'
FROM @Employee
WHERE Salary < (SELECT MAX(Salary) FROM @Employee);

----------------
;with src as(
select Id, Salary, ROW_NUMBER() over(order by Salary desc) RN
from @Employee)
select Salary as 'SecondHighestSalary'
from src
where RN=2