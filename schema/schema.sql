CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users
(
    id         UUID PRIMARY KEY             DEFAULT uuid_generate_v4(),
    email      VARCHAR(255) UNIQUE NOT NULL,
    password   VARCHAR(255)        NOT NULL,
    role       VARCHAR(255)        NOT NULL,
    created_at TIMESTAMPTZ         NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS pvz
(
    id         UUID PRIMARY KEY      DEFAULT uuid_generate_v4(),
    city       VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS receptions
(
    id         UUID PRIMARY KEY      DEFAULT uuid_generate_v4(),
    pvz_id     UUID      NOT NULL REFERENCES pvz (id) ON DELETE CASCADE,
    status     VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS products
(
    id           UUID PRIMARY KEY      DEFAULT uuid_generate_v4(),
    reception_id UUID      NOT NULL REFERENCES receptions (id) ON DELETE CASCADE,
    created_at   TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    type         VARCHAR(255) NOT NULL
);
