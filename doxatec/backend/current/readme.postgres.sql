-- enable the uuid-ossp extension (if not already enabled)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- generate a UUID value
SELECT uuid_generate_v4();
-- create clients table
CREATE TABLE IF NOT EXISTS Clients (
    id UUID PRIMARY KEY,
    name VARCHAR(50) UNIQUE,
    email VARCHAR(250) UNIQUE,
    phone VARCHAR(20) UNIQUE,
    created TIMESTAMP,
    modified TIMESTAMP NULL,
    isolated BOOLEAN DEFAULT FALSE
);
-- create users table
CREATE TABLE IF NOT EXISTS Users (
    id UUID PRIMARY KEY,
    client_id UUID,
    FOREIGN KEY (client_id) REFERENCES Clients(id),
    avatar VARCHAR(50),
    username VARCHAR(50) UNIQUE,
    password VARCHAR(250),
    role VARCHAR(20),
    created TIMESTAMP,
    modified TIMESTAMP NULL,
    accessed TIMESTAMP NULL,
    isolated BOOLEAN DEFAULT FALSE
);
-- create devices table
CREATE TABLE IF NOT EXISTS Devices (
    id UUID PRIMARY KEY,
    user_id UUID,
    FOREIGN KEY (user_id) REFERENCES Users(id),
    name VARCHAR(50),
    pin VARCHAR(50),
    created TIMESTAMP,
    modified TIMESTAMP NULL,
    isolated BOOLEAN DEFAULT FALSE
);
-- create temperatures table
CREATE TABLE IF NOT EXISTS Temperatures (
    id UUID PRIMARY KEY,
    device_id UUID,
    FOREIGN KEY (device_id) REFERENCES Devices(id),
    temp_sup NUMERIC(6, 2),
    temp_mid NUMERIC(6, 2),
    temp_sub NUMERIC(6, 2),
    created TIMESTAMP,
    isolated BOOLEAN DEFAULT FALSE
);
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