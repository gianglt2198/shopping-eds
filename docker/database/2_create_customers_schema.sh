#!/bin/sh
set -e 

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "shopping" <<-EOSQL
    CREATE SCHEMA customers;
    CREATE TABLE customers.customer
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

    CREATE TRIGGER created_at_customers_trg BEFORE UPDATE ON customers.customer FOR EACH ROW EXECUTE PROCEDURE created_at_trigger();
    CREATE TRIGGER updated_at_customers_trg BEFORE UPDATE ON customers.customer FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();
EOSQL