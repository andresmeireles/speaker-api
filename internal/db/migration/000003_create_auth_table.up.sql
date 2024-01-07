CREATE TABLE auths (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES "users"(id) ON DELETE CASCADE,
    hash TEXT NOT NULL,
    expired BOOLEAN NOT NULL
);