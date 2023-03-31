-- enable the uuid-ossp extension (if not already enabled)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- generate a UUID value
SELECT uuid_generate_v4();
-- create transfers table
CREATE TABLE IF NOT EXISTS Transfers (
    id UUID PRIMARY KEY,
    subscription_id UUID,
    FOREIGN KEY (subscription_id) REFERENCES Subscriptions(id),
    amount NUMERIC(12, 2),
    initial_balance: NUMERIC(12, 2),
    final_balance: NUMERIC(12, 2),
    created TIMESTAMP,
    isolated BOOLEAN DEFAULT FALSE
);