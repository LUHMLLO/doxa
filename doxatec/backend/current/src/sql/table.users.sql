-- enable the uuid-ossp extension (if not already enabled)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- generate a UUID value
SELECT uuid_generate_v4();
-- create users table
CREATE TABLE IF NOT EXISTS Users (
    id UUID PRIMARY KEY,
    client_id UUID,
    FOREIGN KEY (client_id) REFERENCES Clients(id),
    avatar VARCHAR(50),
    username VARCHAR(50) UNIQUE,
    password VARCHAR(250),
    role VARCHAR(20),
    created TIMESTAMP,
    modified TIMESTAMP NULL,
    accessed TIMESTAMP NULL,
    isolated BOOLEAN DEFAULT FALSE
);