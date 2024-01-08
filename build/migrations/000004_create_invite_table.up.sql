CREATE TABLE invites (
    id SERIAL PRIMARY KEY,
    theme VARCHAR(100) NOT NULL,
    "references" TEXT NOT NULL,
    date TIMESTAMP NOT NULL,
    time INTEGER NOT NULL,
    accepted BOOLEAN NOT NULL,
    remembered BOOLEAN NOT NULL,
    person_id INT REFERENCES persons(id) ON DELETE RESTRICT
);