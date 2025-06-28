CREATE TABLE IF NOT EXISTS role_access (
    id BIGSERIAL PRIMARY KEY,
    role_id BIGINT REFERENCES roles(id) NOT NULL,
    resource_key TEXT NOT NULL
);


CREATE INDEX IF NOT EXISTS role_access_res_key_idx ON role_access(resource_key);
CREATE INDEX IF NOT EXISTS role_access_role_id_idx ON role_access(role_id);
CREATE INDEX IF NOT EXISTS role_access_role_id_res_key_idx ON role_access(role_id, resource_key);