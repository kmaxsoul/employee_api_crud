SET search_path TO crud_project;

CREATE TABLE IF NOT EXISTS attendance (
    id              SERIAL PRIMARY KEY,
    employee_id     INT NOT NULL REFERENCES employees(id),
    first_name      VARCHAR(255) NOT NULL,
    last_name       VARCHAR(255) NOT NULL,
    department_name VARCHAR(255),
    check_in        TIMESTAMP NOT NULL,
    check_out       TIMESTAMP
);
