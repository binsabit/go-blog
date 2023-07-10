ALTER TABLE IF EXISTS users
    ADD CONSTRAINT uq_username UNIQUE (username);