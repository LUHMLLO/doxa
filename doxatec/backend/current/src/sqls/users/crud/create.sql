-- insert client
INSERT INTO Users (client_id, avatar, username, password, role)
VALUES ($1, $2, $3, $4, $5)