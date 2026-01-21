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

-- Insert sample data
INSERT INTO employees (name, sex, age, salary) VALUES
    ('John Doe', 'M', 30, 75000),
    ('Jane Smith', 'F', 28, 82000),
    ('Bob Johnson', 'M', 35, 95000),
    ('Alice Brown', 'F', 32, 88000)
ON CONFLICT DO NOTHING;


INSERT INTO departments (code, name) VALUES
    ('IT', 'Information Technology'),
    ('HR', 'Human Resources'),
    ('FIN', 'Finance'),
    ('OPS', 'Operations')
ON CONFLICT (code) DO NOTHING;

