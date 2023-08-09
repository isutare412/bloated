-- Create "users" table
CREATE TABLE "users" ("id" uuid NOT NULL, "email" character varying NOT NULL, "user_name" character varying NOT NULL, "given_name" character varying NOT NULL, "family_name" character varying NOT NULL, "photo_url" character varying NOT NULL, "origin" character varying NOT NULL, PRIMARY KEY ("id"));
-- Drop index "todo_user_id" from table: "todos"
DROP INDEX "todo_user_id";
-- Modify "todos" table
DELETE FROM "todos";
ALTER TABLE "todos" ALTER COLUMN "user_id" DROP NOT NULL, ADD COLUMN "owner_id" uuid NOT NULL, ADD CONSTRAINT "todos_users_todos" FOREIGN KEY ("owner_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
