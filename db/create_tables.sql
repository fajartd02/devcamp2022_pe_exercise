CREATE TABLE IF NOT EXISTS items (
    item_id SERIAL PRIMARY KEY,
    stock INT,
    discount INT
);

CREATE TABLE IF NOT EXISTS product (
    product_id SERIAL PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    product_description TEXT NULL,
    price NUMERIC NOT NULL,
    item_id INT FOREIGN KEY REFERENCES (item_id)
);