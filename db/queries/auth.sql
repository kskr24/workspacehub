-- name: InsertSessions :one
INSERT INTO sessions (ua, ip, user_id, token, expires, created, updated)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id;

-- name: GetWorkspacesForUser :many
SELECT w.id, w.name, w.description, w.created, w.updated
FROM workspace w
JOIN user_workspaces uw ON w.id = uw.workspace_id
WHERE uw.user_id = $1;
