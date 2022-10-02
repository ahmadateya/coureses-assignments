-- Problem

-- Find the people who appear most frequently as crew members.
-- Details: Print the names and number of appearances of the 20 people with the most crew appearances ordered by their number of appearances in a descending fashion.
-- Your output should look like this: NAME|NUM_APPEARANCES


-- My Solution

SELECT people.name, COUNT(crew.person_id) AS num_appearances FROM crew 
LEFT JOIN people ON crew.person_id=people.person_id
GROUP BY crew.person_id
ORDER BY num_appearances DESC
LIMIT 20;
