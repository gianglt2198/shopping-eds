CREATE TABLE IF NOT EXISTS payment
(
    id          text NOT NULL,
    order_id    text NOT NULL,
    customer_id text NOT NULL,
    amount      double precision NOT NULL,
    status      text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id)
);