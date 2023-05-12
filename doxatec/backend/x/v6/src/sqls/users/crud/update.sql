-- update client 
UPDATE Users
SET %s = $2
WHERE id = $1