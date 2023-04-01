-- create devices table
CREATE TABLE IF NOT EXISTS Devices (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID,
    FOREIGN KEY (user_id) REFERENCES Users(id),
    name VARCHAR(50),
    pin VARCHAR(50),
    created TIMESTAMP,
    modified TIMESTAMP NULL,
    isolated BOOLEAN DEFAULT FALSE
);