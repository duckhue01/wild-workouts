-- name: GetDemo :one
SELECT * FROM demo
WHERE id = $1 LIMIT 1;