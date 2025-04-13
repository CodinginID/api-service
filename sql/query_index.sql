-- point 2 strategy for creating indexes
-- This SQL script creates indexes on the users table to improve query performance.
-- The indexes are created on the username, email, and created_at columns.
-- These indexes will help speed up queries that filter or sort by these columns.

CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_created_at ON users(created_at);

-- point 3 strategy for creating indexes
