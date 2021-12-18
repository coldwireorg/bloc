SELECT 'CREATE DATABASE bloc ENCODING ''UTF8'''
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'bloc')\gexec

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TYPE access AS ENUM ('PRIVATE', 'SHARED', 'PUBLIC');

CREATE TABLE IF NOT EXISTS users (
    username    VARCHAR(25) PRIMARY KEY NOT NULL UNIQUE,
    password    VARCHAR(256) NOT NULL,
    public_key  BYTEA NOT NULL,
    private_key BYTEA NOT NULL,
    quota       BIGINT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS files (
    id          UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name        VARCHAR(128) NOT NULL,
    size        BIGINT NOT NULL,
    type        VARCHAR(128) NOT NULL,
    chunk       BIGINT NOT NULL,
    f_owner     VARCHAR(25) NOT NULL,
    last_edit   TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),
    CONSTRAINT fk_owner 
        FOREIGN KEY(f_owner) 
            REFERENCES users(username)
);

CREATE TABLE IF NOT EXISTS file_access (
    id              UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    access_state    access NOT NULL DEFAULT 'PRIVATE',
    f_shared_to     VARCHAR(25) DEFAULT NULL,
    f_shared_by     VARCHAR(25) NOT NULL,
    f_file          UUID NOT NULL,
    favorite        BOOLEAN NOT NULL DEFAULT FALSE,
    encryption_key  BYTEA NOT NULL,
    CONSTRAINT fk_shared_to
        FOREIGN KEY(f_shared_to)
            REFERENCES users(username),
    CONSTRAINT fk_shared_by
        FOREIGN KEY(f_shared_by)
            REFERENCES users(username),
    CONSTRAINT fk_file
        FOREIGN KEY(f_file)
            REFERENCES files(id)
);
