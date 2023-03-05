BEGIN;

CREATE TABLE clients(
    id SERIAL PRIMARY KEY,
    name VARCHAR(64) UNIQUE NOT NULL,
    inn VARCHAR(20) UNIQUE NOT NULL
);

CREATE TABLE contacts(
    id SERIAL PRIMARY KEY,
    fullname VARCHAR(64) NOT NULL,
    phone VARCHAR(64),
    email VARCHAR(64),
    Position VARCHAR(64),
    client_id INTEGER,
    FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE
);

END;