CREATE TABLE request (
    id SERIAL PRIMARY KEY,
    user_email TEXT,
    url TEXT NOT NULL,
    method TEXT NOT NULL,
    origin TEXT NOT NULL,
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
    history INTEGER[],
    favorites INTEGER[],
    date TEXT NOT NULL,
    deleted BOOLEAN NOT NULL
);

ALTER TABLE request
ADD CONSTRAINT fk_request_user_email
FOREIGN KEY (user_email)
REFERENCES "user"(email);
