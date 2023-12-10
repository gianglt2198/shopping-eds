#!/bin/sh
set -e 

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "shopping" <<-EOSQL
    CREATE SCHEMA payments;
    CREATE TABLE payments.payment
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

    CREATE TRIGGER created_at_payments_trg BEFORE UPDATE ON payments.payment FOR EACH ROW EXECUTE PROCEDURE created_at_trigger();
    CREATE TRIGGER updated_at_payments_trg BEFORE UPDATE ON payments.payment FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();
EOSQL