-- The purpose of this query is to make sure that the formatting of your output matches exactly the formatting of our auto-grading script.
-- Details: List all Category Names ordered alphabetically.

SELECT DISTINCT(language)
FROM akas
ORDER BY language 
LIMIT 10;