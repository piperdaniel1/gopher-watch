SET session_replication_role = replica;

-- Drop existing tables
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS users;

-- Re-enable referential integrity
SET session_replication_role = DEFAULT;

-- Recreate tables
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    product VARCHAR(100) NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Insert initial data
INSERT INTO users (name, email) VALUES
('Alice', 'alice@example.com'),
('Bob', 'bob@example.com');

INSERT INTO orders (user_id, product, amount) VALUES
(1, 'Product A', 9.99),
(2, 'Product B', 19.99);