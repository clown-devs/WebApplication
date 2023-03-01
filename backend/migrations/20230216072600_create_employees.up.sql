BEGIN;
CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(32),
    second_name VARCHAR(32),
    third_name VARCHAR(32),

    username VARCHAR(32) UNIQUE,
    encrypted_password VARCHAR(128) NOT NULL
);

CREATE TABLE tokens (
    id SERIAL PRIMARY KEY,
    employee_id INT NOT NULL,
    token VARCHAR(64) UNIQUE NOT NULL,
    CONSTRAINT fk_employee
        FOREIGN KEY(employee_id)
        REFERENCES employees(id)   
        ON DELETE CASCADE
);
END;