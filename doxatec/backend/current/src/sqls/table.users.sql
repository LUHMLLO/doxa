-- create users table
CREATE TABLE IF NOT EXISTS Users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
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