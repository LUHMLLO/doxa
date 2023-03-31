-- enable the uuid-ossp extension (if not already enabled)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- generate a UUID value
SELECT uuid_generate_v4();
-- create devices table
CREATE TABLE IF NOT EXISTS Devices (
    id UUID PRIMARY KEY,
    user_id UUID,
    FOREIGN KEY (user_id) REFERENCES Users(id),
    name VARCHAR(50),
    pin VARCHAR(50),
    created TIMESTAMP,
    modified TIMESTAMP NULL,
    isolated BOOLEAN DEFAULT FALSE
);