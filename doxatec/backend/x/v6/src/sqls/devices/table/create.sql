-- create devices table
CREATE TABLE IF NOT EXISTS Devices (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID,
    FOREIGN KEY (user_id) REFERENCES Users(id),
    name VARCHAR(250),
    pin VARCHAR(250),
    created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    modified TIMESTAMP WITH TIME ZONE NULL
);