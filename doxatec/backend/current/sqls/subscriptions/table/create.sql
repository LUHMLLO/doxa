-- create subscriptions table
CREATE TABLE IF NOT EXISTS Subscriptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(250),
    description VARCHAR(250),
    price NUMERIC(12, 2),
    created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    modified TIMESTAMP WITH TIME ZONE NULL
);