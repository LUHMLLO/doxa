# Generate Generic Function

## Postgres tables

CREATE TABLE IF NOT EXISTS Clients (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(250),
    email VARCHAR(250) UNIQUE,
    phone VARCHAR(250) UNIQUE,
    created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    modified TIMESTAMP WITH TIME ZONE NULL
);

CREATE TABLE IF NOT EXISTS Users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    client_id UUID,
    FOREIGN KEY (client_id) REFERENCES Clients(id),
    avatar VARCHAR(250),
    username VARCHAR(250) UNIQUE,
    password VARCHAR(250),
    role VARCHAR(250),
    created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    modified TIMESTAMP WITH TIME ZONE NULL,
    accessed TIMESTAMP WITH TIME ZONE NULL
);

CREATE TABLE IF NOT EXISTS Devices (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID,
    FOREIGN KEY (user_id) REFERENCES Users(id),
    name VARCHAR(250),
    pin VARCHAR(250),
    created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    modified TIMESTAMP WITH TIME ZONE NULL
);

## Go structs

type Client struct {
 ID       uuid.UUID  `json:"id"`
 Name     string     `json:"name"`
 Email    string     `json:"email"`
 Phone    string     `json:"phone"`
 Created  time.Time  `json:"created"`
 Modified *time.Time `json:"modified"`
}

type User struct {
 ID       uuid.UUID  `json:"id"`
 Client   uuid.UUID  `json:"client_id"`
 Avatar   string     `json:"avatar"`
 Username string     `json:"username"`
 Password string     `json:"password"`
 Role     string     `json:"role"`
 Created  time.Time  `json:"created"`
 Modified*time.Time `json:"modified"`
 Accessed*time.Time `json:"accessed"`
}

type Device struct {
 ID       uuid.UUID     `json:"id"`
 User     uuid.UUID     `json:"user_id"`
 Name     string        `json:"name"`
 PIN      string        `json:"pin"`
 TEMPS    []Temperature `json:"temps"`
 Created  time.Time     `json:"created"`
 Modified*time.Time    `json:"modified"`
}

## Goal

Generate a function to update one or more columns, the columns to be updated will be declared by a json request body
