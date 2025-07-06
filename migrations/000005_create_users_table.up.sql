CREATE TABLE IF not exists users (
	id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
	name TEXT NOT NULL,
	email citext NOT NULL UNIQUE,
    password_hash bytea NOT NULL,
    activated BOOLEAN NOT NULL DEFAULT FALSE,
    version INTEGER NOT NULL DEFAULT 1
);
