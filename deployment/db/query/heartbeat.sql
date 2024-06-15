-- name: UpdateLastHeartbeat :one
UPDATE heartbeat SET last_heartbeat = $2 WHERE machine_id = $1 RETURNING *;

-- name: GetLastHeartbeat :one
SELECT last_heartbeat FROM heartbeat WHERE machine_id = $1;
