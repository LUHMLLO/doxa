-- create transfers table
CREATE TABLE IF NOT EXISTS Transfers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    client_id UUID,
    FOREIGN KEY (client_id) REFERENCES Clients(id),
    subscription_id UUID,
    FOREIGN KEY (subscription_id) REFERENCES Subscriptions(id),
    amount NUMERIC(12, 2),
    created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);