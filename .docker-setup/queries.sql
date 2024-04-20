CREATE TABLE request (
    id SERIAL PRIMARY KEY,
    user_email TEXT NOT NULL,
    url TEXT NOT NULL,
    method TEXT NOT NULL,
    origin TEXT,
    headers TEXT,
    body TEXT,
    status TEXT NOT NULL,
    date TEXT NOT NULL,
    hidden BOOLEAN NOT NULL
);

CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    favorites INTEGER[],
    date TEXT NOT NULL,
    deleted BOOLEAN NOT NULL
);

ALTER TABLE request
ADD CONSTRAINT fk_request_user_email
FOREIGN KEY (user_email)
REFERENCES "user"(email);

INSERT INTO "user" (username, email, password, date, deleted) VALUES ('anon', 'null', 'anon', CAST(CAST(EXTRACT(EPOCH FROM NOW()) AS INTEGER) AS TEXT), false);

-- psql -U postgres -c "CREATE DATABASE postwoman;"
-- psql -U postgres -d postwoman -f './queries.sql'
