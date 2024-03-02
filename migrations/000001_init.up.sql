CREATE TABLE IF NOT EXISTS role (
    id serial primary key,
    name varchar not null
);

CREATE TABLE IF NOT EXISTS user (
    id serial primary key,
    email varchar not null,
    password varchar not null,
    created_at timestamp not null
);

CREATE INDEX IF NOT EXISTS idx_role_name_btree ON role(name);
CREATE INDEX IF NOT EXISTS idx_user_email_btree ON user(email);