CREATE TABLE auth_codes (
    id SERIAL PRIMARY KEY,
    code VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    user_id INT REFERENCES users(id) ON DELETE RESTRICT
);
