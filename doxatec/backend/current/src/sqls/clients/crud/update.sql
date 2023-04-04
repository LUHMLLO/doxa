-- update client
UPDATE Clients
SET %s = $2
WHERE id = $1