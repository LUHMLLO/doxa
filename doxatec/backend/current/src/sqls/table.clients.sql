-- create clients table
CREATE TABLE IF NOT EXISTS Clients (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50) UNIQUE,
    email VARCHAR(250) UNIQUE,
    phone VARCHAR(20) UNIQUE,
    created TIMESTAMP,
    modified TIMESTAMP NULL,
    isolated BOOLEAN DEFAULT FALSE
);