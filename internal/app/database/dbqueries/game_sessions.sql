-- name: CreateGameSession :execresult
INSERT INTO game_sessions (
    name, token, credits
) VALUES (?, ?, ?);

-- name: GetGameSessionByToken :one
SELECT * FROM game_sessions
WHERE token = ? AND deleted_at IS NULL LIMIT 1;

-- name: DeleteGameSessionByToken :execresult
UPDATE game_sessions
SET deleted_at = NOW()
WHERE token = ? AND deleted_at IS NULL;