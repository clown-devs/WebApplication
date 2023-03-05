BEGIN;

CREATE TABLE directions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(64) UNIQUE NOT NULL
);

CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(64) UNIQUE NOT NULL,
    is_super BOOL NOT NULL,
    can_see_meetings BOOL NOT NULL
);

CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    fullname VARCHAR(128),
    username VARCHAR(32) UNIQUE NOT NULL,
    encrypted_password VARCHAR(128) NOT NULL,
    direction_id INT NOT NULL,
    role_id INT NOT NULL,

    CONSTRAINT fk_direction
        FOREIGN KEY(direction_id)
        REFERENCES directions(id)   
        ON DELETE RESTRICT,

    CONSTRAINT fk_role 
        FOREIGN KEY(role_id)
        REFERENCES roles(id)   
        ON DELETE RESTRICT
    
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

CREATE TABLE clients_employee (
    id SERIAL PRIMARY KEY,
    client_id INTEGER,
    employee_id INTEGER,
    FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE,
    FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE CASCADE
);
END;