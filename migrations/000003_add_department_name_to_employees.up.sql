SET search_path TO crud_project;

ALTER TABLE employees
    ADD COLUMN IF NOT EXISTS department_name VARCHAR(255);

UPDATE employees e
SET department_name = d.name
FROM departments d
WHERE e.department_id = d.id;
