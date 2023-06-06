CREATE TABLE notif.notification (
  "id" UUID PRIMARY KEY,
  "receiver_id" UUID,
  "data" JSON,
  "is_seen" boolean,
  "event" CHAR(50),
  "sender_id" UUID,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMP DEFAULT NULL
);

CREATE INDEX receiver_id_index ON notif.notification (receiver_id);
