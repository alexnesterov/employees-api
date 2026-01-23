-- Create employee schema
CREATE SCHEMA IF NOT EXISTS employees;

-- Set search path to employee schema
SET search_path TO employees;

-- Create employee table
CREATE TABLE IF NOT EXISTS employees (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    sex VARCHAR(10) NOT NULL,
    age INTEGER NOT NULL,
    salary INTEGER NOT NULL,
    department_code VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create departments table
CREATE TABLE IF NOT EXISTS departments (
    code VARCHAR(50) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Add foreign key constraint
ALTER TABLE employees 
ADD CONSTRAINT fk_employees_department 
FOREIGN KEY (department_code) 
REFERENCES departments(code)
ON DELETE SET NULL;

-- Insert sample data
INSERT INTO departments (code, name) VALUES
    ('IT', 'Information Technology'),
    ('HR', 'Human Resources'),
    ('FIN', 'Finance'),
    ('OPS', 'Operations')
ON CONFLICT (code) DO NOTHING;

INSERT INTO employees (name, sex, age, salary, department_code) VALUES
    ('John Doe', 'M', 30, 75000, 'IT'),
    ('Jane Smith', 'F', 28, 82000, 'HR'),
    ('Bob Johnson', 'M', 35, 95000, 'FIN'),
    ('Alice Brown', 'F', 32, 88000, 'OPS')
ON CONFLICT DO NOTHING;
