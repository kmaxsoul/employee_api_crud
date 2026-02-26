SET search_path TO crud_project;

ALTER TABLE employees
    DROP COLUMN IF EXISTS department_name;
