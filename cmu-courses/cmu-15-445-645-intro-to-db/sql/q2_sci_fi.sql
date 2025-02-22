-- Problem

-- Find the 10 `Sci-Fi` works with the longest runtimes.
-- Details: Print the title of the work, the premiere date, and the runtime. 
-- The column listing the runtime should be suffixed with the string " (mins)", 
-- for example, if the runtime_mins value is `12`, you should output 12 (mins). 
-- Note a work is Sci-Fi even if it is categorized in multiple genres, as long as Sci-Fi is one of the genres.
-- Your first row should look like this: Cicak-Man 2: Planet Hitam|2008|999 (mins)


-- My Solution
SELECT primary_title,premiered,runtime_minutes || ' (mins)' 
FROM (titles) 
WHERE genres LIKE '%Sci-Fi%' 
ORDER BY runtime_minutes 
DESC LIMIT 10;