-- create subscriptions table
CREATE TABLE IF NOT EXISTS Subscriptions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
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