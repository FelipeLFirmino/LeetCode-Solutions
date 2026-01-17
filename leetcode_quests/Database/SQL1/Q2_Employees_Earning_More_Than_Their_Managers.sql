--write your MySQL query statement below
-- To solve this question we need to self join the table so we can compare the values
SELECT 
    e.name AS Employee
FROM Employee AS e
JOIN Employee AS m 
    ON e.managerId = m.id
WHERE e.salary > m.salary; 
