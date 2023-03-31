-- enable the uuid-ossp extension (if not already enabled)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- generate a UUID value
SELECT uuid_generate_v4();
-- create clients table
CREATE TABLE IF NOT EXISTS Clients (
    id UUID PRIMARY KEY,
    name VARCHAR(50) UNIQUE,
    email VARCHAR(250) UNIQUE,
    phone VARCHAR(20) UNIQUE,
    created TIMESTAMP,
    modified TIMESTAMP NULL,
    isolated BOOLEAN DEFAULT FALSE
);
