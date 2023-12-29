#!/bin/sh
set -e 

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "shopping" <<-EOSQL
    CREATE SCHEMA ordering;
    CREATE TABLE ordering.order
    (
        order_id              text NOT NULL,
        customer_id     text NOT NULL,
        customer_name   text NOT NULL,
        items           bytea,
        status          text NOT NULL,
        product_ids     text,
        created_at      timestamptz NOT NULL DEFAULT NOW(),
        updated_at      timestamptz NOT NULL DEFAULT NOW(),
        PRIMARY KEY (order_id)
    );

    CREATE TRIGGER created_at_ordering_trg BEFORE UPDATE ON ordering.order FOR EACH ROW EXECUTE PROCEDURE created_at_trigger();
    CREATE TRIGGER updated_at_ordering_trg BEFORE UPDATE ON ordering.order FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();

    CREATE TABLE ordering.products_cache
    (
        id         text NOT NULL,
        name       text NOT NULL,
        price      decimal(9,4) NOT NULL,
        created_at timestamptz NOT NULL DEFAULT NOW(),
        updated_at timestamptz NOT NULL DEFAULT NOW(),
        PRIMARY KEY (id)
    );

    CREATE TRIGGER created_at_products_trgr BEFORE UPDATE ON ordering.products_cache FOR EACH ROW EXECUTE PROCEDURE created_at_trigger();
    CREATE TRIGGER updated_at_products_trgr BEFORE UPDATE ON ordering.products_cache FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();

    CREATE TABLE ordering.customers_cache
    (
        id         text NOT NULL,
        name       text NOT NULL,
        created_at timestamptz NOT NULL DEFAULT NOW(),
        updated_at timestamptz NOT NULL DEFAULT NOW(),
        PRIMARY KEY (id)
    );

    CREATE TRIGGER created_at_customer_trgr BEFORE UPDATE ON ordering.customers_cache FOR EACH ROW EXECUTE PROCEDURE created_at_trigger();
    CREATE TRIGGER updated_at_customer_trgr BEFORE UPDATE ON ordering.customers_cache FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();


    CREATE TABLE ordering.events (
        stream_id       text NOT NULL,
        stream_name     text NOT NULL,
        stream_version  int NOT NULL,
        event_id        text NOT NULL,
        event_name      text NOT NULL,
        event_data      bytea NOT NULL,
        occurred_at     timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (stream_id, stream_name, stream_version)
    );

    CREATE TABLE ordering.snapshots (
        stream_id        text        NOT NULL,
        stream_name      text        NOT NULL,
        stream_version   int         NOT NULL,
        snapshot_name    text        NOT NULL,
        snapshot_data    bytea       NOT NULL,
        updated_at       timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (stream_id, stream_name)
    );

    CREATE TRIGGER updated_at_snapshots_trg BEFORE UPDATE ON ordering.snapshots FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();
EOSQL