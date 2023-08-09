-- Modify "users" table
DELETE FROM "users";
ALTER TABLE "users" ADD COLUMN "create_time" timestamptz NOT NULL, ADD COLUMN "update_time" timestamptz NOT NULL;
