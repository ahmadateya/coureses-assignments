-- Problem

-- Determine the oldest people in the dataset who were born in or after 1900. You should assume that a person without a known death year is still alive.
-- Details: Print the name and age of each person. People should be ordered by a compound value of their age and secondly their name in alphabetical order.
-- Return the first 20 results.
-- Your output should have the format: NAME|AGE


-- My Solution

SELECT name,
    CASE WHEN died IS NULL -- is alive
    THEN strftime('%Y','now') - born -- calc year between now (in year) and born
    ELSE died - born 
    END AS age 
FROM people 
WHERE born >= 1900 
ORDER BY age DESC ,name 
LIMIT 20;