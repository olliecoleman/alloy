-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE admin_users (
    id SERIAL PRIMARY KEY,
    name CHARACTER VARYING(255),
    email CHARACTER VARYING(255),
    password_hash CHARACTER VARYING(255),
    inserted_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

CREATE UNIQUE INDEX admin_users_email_index ON admin_users USING btree (email);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE admin_users;