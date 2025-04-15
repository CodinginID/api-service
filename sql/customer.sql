CREATE TABLE orders (
  id INT PRIMARY KEY,
  customer_id INT,
  product_id INT,
  order_date TIMESTAMP,
  amount DECIMAL(10, 2)
);

-- This SQL script creates indexes on the orders table to improve query performance.
-- The indexes are created on the customer_id, product_id, and order_date columns.
-- These indexes will help speed up queries that filter or sort by these columns.
CREATE INDEX idx_orders_customer_id ON orders(customer_id);
CREATE INDEX idx_orders_product_id ON orders(product_id);
CREATE INDEX idx_orders_order_date ON orders(order_date);