CREATE TABLE invites (
    id SERIAL PRIMARY KEY,
    theme VARCHAR(100),
    date TIMESTAMP,
    time INTEGER,
    accepted BOOLEAN,
    remembered BOOLEAN,
    person INT REFERENCES persons(id) ON DELETE RESTRICT
);