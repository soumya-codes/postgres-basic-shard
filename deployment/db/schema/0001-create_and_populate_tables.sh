#!/bin/bash
set -e

echo "creating schema"

# Connect to PostgreSQL and execute the commands
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL

CREATE TABLE IF NOT EXISTS heartbeat (
    machine_id TEXT,
    last_heartbeat BIGINT
);

ALTER TABLE heartbeat ADD CONSTRAINT unique_machine_id UNIQUE (machine_id);

\copy heartbeat from '/var/lib/data/init/heartbeat.csv' with delimiter E'\t' null ''

EOSQL
echo "Data import complete"