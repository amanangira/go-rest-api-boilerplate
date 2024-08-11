-- migrate:up
CREATE EXTENSION "uuid-ossp";
CREATE TABLE
    users (
        id  UUID PRIMARY KEY DEFAULT UUID_GENERATE_V4(),
        name TEXT NOT NULL,
        email TEXT NOT NULL,
        password TEXT NOT NULL,
        token TEXT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP DEFAULT NULL
    );

CREATE UNIQUE INDEX users_unique_email ON users(email) WHERE deleted_at IS NULL;

-- migrate:down

DROP INDEX users_unique_email;
DROP TABLE users;
DROP EXTENSION "uuid-ossp";