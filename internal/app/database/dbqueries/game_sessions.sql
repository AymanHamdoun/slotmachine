-- name: GetGameSession :one
SELECT * FROM game_sessions
WHERE id = ? AND deleted_at IS NULL LIMIT 1;

-- name: CreateGameSession :execresult
INSERT INTO game_sessions (
    name, token, credits
) VALUES (?, ?, ?);

-- name: DeleteGameSession :exec
UPDATE game_sessions
SET deleted_at = NOW()
WHERE id = ?;