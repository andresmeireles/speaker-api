CREATE TABLE invites (
    id SERIAL PRIMARY KEY,
    theme VARCHAR(100),
    date DATE,
    accept BOOLEAN,
    remembered BOOLEAN,
    user_id INT REFERENCES persons(id) ON DELETE RESTRICT,
);