BEGIN;
CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(32),
    second_name VARCHAR(32),
    third_name VARCHAR(32),

    username VARCHAR(32),
    encrypted_password VARCHAR(128) NOT NULL
);
END;