-- name: CreateWorkspace :one
INSERT INTO workspaces (id, name, created_by)
VALUES ($1, $2, $3)
RETURNING *;

-- name: AddWorkspaceMember :exec
INSERT INTO workspace_members (workspace_id, user_id, role)
VALUES ($1, $2, $3);

-- name: ListUserWorkspaces :many
SELECT w.*
FROM workspaces w
JOIN workspace_members wm ON w.id = wm.workspace_id
WHERE wm.user_id = $1
ORDER BY w.created_at DESC;
