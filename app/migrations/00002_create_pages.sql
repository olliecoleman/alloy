-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE pages (
    id SERIAL PRIMARY KEY,
    title CHARACTER VARYING(255),
    page_title CHARACTER VARYING(255),
    meta_description TEXT,
    content TEXT,
    slug CHARACTER VARYING(255),
    inserted_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    layout CHARACTER VARYING(255) DEFAULT 'two-col'
);

CREATE UNIQUE INDEX pages_slug_index ON pages USING btree (slug);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE pages;