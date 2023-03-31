-- enable the uuid-ossp extension (if not already enabled)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- generate a UUID value
SELECT uuid_generate_v4();
-- create subscriptions table
CREATE TABLE IF NOT EXISTS Subscriptions (
    id UUID PRIMARY KEY,
    user_id UUID,
    FOREIGN KEY (user_id) REFERENCES Users(id),
    amount NUMERIC(12, 2),
    interest_rate NUMERIC(6, 2),
    term_duration INT,
    term_frequency VARCHAR(250),
    loan_start TIMESTAMP,
    loan_end TIMESTAMP,
    total_interest VARCHAR(250),
    total_loan VARCHAR(250),
    monthly_interest VARCHAR(250),
    monthly_loan VARCHAR(250),
    monthly_fee VARCHAR(250),
    approval BOOLEAN,
    created TIMESTAMP,
    modified TIMESTAMP NULL,
    isolated BOOLEAN DEFAULT FALSE
);