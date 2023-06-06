-- name: ListUserNotifications :many
SELECT * FROM  notif.notification WHERE receiver_id = $1 AND id > $2 LIMIT $3;

-- name: InsertNotification :exec
INSERT INTO notif.notification(id, receiver_id, data, is_seen, event, sender_id)
VALUES ($1, $2, $3, $4, $5, $6);


-- name: GetNotification :one
SELECT * FROM  notif.notification WHERE id = $1;

-- name: UpdateNotification :one
UPDATE notif.notification SET receiver_id = $2, data = $3, is_seen = $4, event = $5, sender_id = $6, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;