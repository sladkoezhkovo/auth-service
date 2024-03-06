CREATE TABLE IF NOT EXISTS roles (
    id serial primary key,
    name varchar not null
);

CREATE TABLE IF NOT EXISTS users (
    id serial primary key,
    email varchar not null,
    password varchar not null,
    created_at timestamp not null default now()
);

CREATE INDEX IF NOT EXISTS idx_roles_name_btree ON roles(name);
CREATE INDEX IF NOT EXISTS idx_users_email_btree ON users(email);