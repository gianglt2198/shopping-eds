CREATE TABLE IF NOT EXISTS ordering
(
    id          text NOT NULL,
    customer_id text NOT NULL,
    payment_id  text NOT NULL,
    items       bytea NOT NULL,
    status      text NOT NULL,
    created_at  timestamptz NOT NULL DEFAULT NOW(),
    updated_at  timestamptz NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id)
);