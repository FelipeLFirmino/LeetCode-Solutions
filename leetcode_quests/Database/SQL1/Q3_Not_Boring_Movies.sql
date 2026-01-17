--  To solve this question we need to know about arithmetic operations, where filter clause and order by

SELECT C.id, C.movie, C.description, C.rating FROM Cinema AS C 
WHERE C.description != "boring" AND (C.id % 2 != 0) 
ORDER BY C.rating DESC