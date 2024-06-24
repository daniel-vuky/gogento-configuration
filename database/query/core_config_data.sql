-- name: GetConfig :one
SELECT *
FROM core_config_data
Where path = $1;

-- name: UpdateConfig :one
UPDATE core_config_data
SET value = $2
Where path = $1
RETURNING *;

-- name: CreateConfig :one
INSERT
INTO core_config_data (path, value)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteConfig :one
DELETE
FROM core_config_data
Where path = $1
RETURNING *;