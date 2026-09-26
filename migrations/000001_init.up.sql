CREATE SCHEMA todoapp;

CREATE TABLE todoapp.user (
    id            BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    version       BIGINT NOT NULL DEFAULT 1,
    full_name     TEXT NOT NULL CHECK(length(full_name) BETWEEN 3 AND 100),
    phone_number  VARCHAR(15) CHECK (
        phone_number ~ '^\+[1-9]\d{1,14}$'
    )
);

CREATE TABLE todoapp.tasks (
    id            BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_id       BIGINT NOT NULL REFERENCES todoapp.user(id),
    version       BIGINT NOT NULL DEFAULT 1,
    title         TEXT NOT NULL CHECK (length(title) BETWEEN 1 AND 100),
    description   TEXT CHECK (length(description) BETWEEN 1 AND 1000),
    completed     BOOLEAN NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    completed_at  TIMESTAMPTZ,

    CHECK (
        (completed = FALSE AND completed_at IS NULL)
        OR
        (completed = TRUE  AND completed_at IS NOT NULL
                           AND completed_at >= created_at)
    )
);