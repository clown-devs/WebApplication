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

INSERT INTO roles(id, name, is_super, can_see_meetings) VALUES
(1, 'Обычный пользователь', false, false),
(2, 'Менеджер', true, false),
(3, 'Администратор', false, true),
(4, 'Главный администратор', true, true);

/*
    1 - Обычный пользователь. Не может смотреть чужие встречи и не может бронировать 5-6 этаж
    2 - Менеджер. Не может смотреть и редачить чужие встречи(Может бронить 5-6 этаж. is_super=True)
    3 - Администратор. Не может бронировать 5-6 этаж, но умеет редачить и смотреть чужие встречи
    4 - Главный администратор. Может делать всё.
*/
CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    fullname VARCHAR(128),
    username VARCHAR(32) UNIQUE NOT NULL,
    encrypted_password VARCHAR(128) NOT NULL,
    direction_id INT,
    role_id INT NOT NULL DEFAULT 1,

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