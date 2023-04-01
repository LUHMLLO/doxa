-- create temperatures table
CREATE TABLE IF NOT EXISTS Temperatures (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    device_id UUID,
    FOREIGN KEY (device_id) REFERENCES Devices(id),
    temp_sup NUMERIC(6, 2),
    temp_mid NUMERIC(6, 2),
    temp_sub NUMERIC(6, 2),
    created TIMESTAMP,
    isolated BOOLEAN DEFAULT FALSE
);