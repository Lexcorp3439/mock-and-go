-- +goose Up
-- +goose StatementBegin

CREATE
    EXTENSION IF NOT EXISTS pgcrypto;

CREATE
    OR REPLACE FUNCTION set_updated_at_column() RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at
        = now() at time zone 'utc';
    RETURN NEW;
END;
$$
    language 'plpgsql';

CREATE TABLE public.template
(
    name        text UNIQUE                                        NOT NULL PRIMARY KEY,
    status      int                      DEFAULT 0                 NOT NULL,
    description text,
    created_at  timestamp with time zone DEFAULT now()             NOT NULL,
    updated_at  timestamp with time zone DEFAULT now()             NOT NULL
);

CREATE TABLE public.template_binding
(
    id           uuid                     DEFAULT gen_random_uuid() NOT NULL PRIMARY KEY,
    name         text references template (name)                    NOT NULL,
    phone_number text                                               NOT NULL,
    status       int                      DEFAULT 0                 NOT NULL,
    created_at   timestamp with time zone DEFAULT now()             NOT NULL,
    updated_at   timestamp with time zone DEFAULT now()             NOT NULL
);

CREATE TRIGGER update_template_updated_at
    BEFORE UPDATE
    ON template
    FOR EACH ROW
EXECUTE PROCEDURE set_updated_at_column();

CREATE TRIGGER update_template_binding_updated_at
    BEFORE UPDATE
    ON template_binding
    FOR EACH ROW
EXECUTE PROCEDURE set_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists template cascade;
drop table if exists template_binding cascade;
-- +goose StatementEnd
