CREATE TABLE IF NOT EXISTS customer
(
    id         text NOT NULL,
    name       text NOT NULL,
    sms_number text NOT NULL,
    email      text NOT NULL,
    active     bool NOT NULL,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    updated_at timestamptz NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id)
);