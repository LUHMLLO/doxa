-- create users table
CREATE TABLE IF NOT EXISTS Users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    client_id UUID,
    FOREIGN KEY (client_id) REFERENCES Clients(id),
    avatar VARCHAR(250),
    username VARCHAR(250) UNIQUE,
    password VARCHAR(250),
    role VARCHAR(250),
    created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    modified TIMESTAMP WITH TIME ZONE NULL,
    accessed TIMESTAMP WITH TIME ZONE NULL
);