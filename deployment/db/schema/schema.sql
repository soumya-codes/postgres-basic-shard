CREATE TABLE IF NOT EXISTS heartbeat (
machine_id TEXT,
last_heartbeat BIGINT
);

-- Ensure the unique constraint on user_id
ALTER TABLE heartbeat ADD CONSTRAINT unique_machine_id UNIQUE (machine_id);