DO $$ 
BEGIN
    IF NOT EXISTS(SELECT 1 FROM pg_database WHERE datname = 'postwoman') THEN
        CREATE DATABASE postwoman;
        \c postwoman;
    END IF;
END $$;

DO $$ 
BEGIN
    IF NOT EXISTS(SELECT 1 FROM information_schema.tables WHERE table_name = 'request') THEN
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
    END IF;
END $$;

DO $$ 
BEGIN
    IF NOT EXISTS(SELECT 1 FROM information_schema.tables WHERE table_name = 'user') THEN
        CREATE TABLE "user" (
            id SERIAL PRIMARY KEY,
            username TEXT NOT NULL,
            email TEXT NOT NULL UNIQUE,
            password TEXT NOT NULL,
            favorites INTEGER[],
            date TEXT NOT NULL,
            deleted BOOLEAN NOT NULL
        );
    END IF;
END $$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.table_constraints WHERE constraint_name = 'fk_request_user_email') THEN
        ALTER TABLE request
        ADD CONSTRAINT fk_request_user_email
        FOREIGN KEY (user_email)
        REFERENCES "user"(email);
    END IF;
END $$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM "user" WHERE email = 'null') THEN
        INSERT INTO "user" (username, email, password, date, deleted) VALUES ('anon', 'null', 'anon', CAST(CAST(EXTRACT(EPOCH FROM NOW()) AS INTEGER) AS TEXT), false);
    END IF;
END $$;
