-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE support_messages (
    id SERIAL PRIMARY KEY,
    name CHARACTER VARYING(255),
    email CHARACTER VARYING(255),
    subject CHARACTER VARYING(255),
    content TEXT,
    inserted_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE support_messages;