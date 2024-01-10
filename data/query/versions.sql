-- name: CreateVersion :exec
INSERT INTO versions (version) VALUES ($1) RETURNING *;

-- name: GetVersion :one
SELECT * FROM versions WHERE id = $1;

-- name: GetLatestVersion :one
SELECT * FROM versions ORDER BY created_at DESC LIMIT 1;

-- name: ListVersions :many
SELECT * FROM versions;

-- name: UpdateVersion :exec
UPDATE versions SET version = $2 WHERE id = $1 RETURNING *;

-- name: DeleteVersion :exec
DELETE FROM versions WHERE id = $1;
