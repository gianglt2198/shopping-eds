#!/bin/sh
set -e 

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "shopping" <<-EOSQL
    CREATE SCHEMA products;
    CREATE TABLE products.product
    (
        id         text NOT NULL,
        name       text NOT NULL,
        description text NOT NULL,
        price      double precision NOT NULL,
        created_at timestamptz NOT NULL DEFAULT NOW(),
        updated_at timestamptz NOT NULL DEFAULT NOW(),
        PRIMARY KEY (id)
    );

    CREATE TRIGGER created_at_products_trg BEFORE UPDATE ON products.product FOR EACH ROW EXECUTE PROCEDURE created_at_trigger();
    CREATE TRIGGER updated_at_products_trg BEFORE UPDATE ON products.product FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();

    CREATE TABLE products.events (
        stream_id       text NOT NULL,
        stream_name     text NOT NULL,
        stream_version  int NOT NULL,
        event_id        text NOT NULL,
        event_name      text NOT NULL,
        event_data      bytea NOT NULL,
        occurred_at     timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (stream_id, stream_name, stream_version)
    );

    CREATE TABLE products.snapshots (
        stream_id        text        NOT NULL,
        stream_name      text        NOT NULL,
        stream_version   int         NOT NULL,
        snapshot_name    text        NOT NULL,
        snapshot_data    bytea       NOT NULL,
        updated_at       timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (stream_id, stream_name)
    );

    CREATE TRIGGER updated_at_snapshots_trg BEFORE UPDATE ON products.snapshots FOR EACH ROW EXECUTE PROCEDURE updated_at_trigger();
EOSQL