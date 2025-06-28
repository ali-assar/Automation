CREATE TABLE IF NOT EXISTS users(
    id BIGSERIAL PRIMARY KEY,
    national_id_number TEXT NOT NULL REFERENCES person(national_id_number) UNIQUE,
    username TEXT NOT NULL UNIQUE,
    hash_password TEXT NOT NULL,
    is_admin BOOLEAN
);

CREATE INDEX IF NOT EXISTS users_username_idx ON users(username);
CREATE INDEX IF NOT EXISTS users_national_id_number_idx ON users(national_id_number);