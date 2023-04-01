-- create transfers table
CREATE TABLE IF NOT EXISTS Transfers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    subscription_id UUID,
    FOREIGN KEY (subscription_id) REFERENCES Subscriptions(id),
    amount NUMERIC(12, 2),
    initial_balance NUMERIC(12, 2),
    final_balance NUMERIC(12, 2),
    created TIMESTAMP,
    isolated BOOLEAN DEFAULT FALSE
);