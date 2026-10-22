CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    description TEXT NOT NULL,
    price INT NOT NULL check (price>=0),
    currency CHAR(3) NOT NULL DEFAULT 'INR',
);



