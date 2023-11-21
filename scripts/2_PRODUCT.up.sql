CREATE TABLE IF NOT EXISTS product
(
    id         text NOT NULL,
    name       text NOT NULL,
    description text NOT NULL,
    price      double precision NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id)
);