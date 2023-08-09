-- Modify "users" table
ALTER TABLE "users" ALTER COLUMN "email" DROP NOT NULL, ALTER COLUMN "user_name" DROP NOT NULL, ALTER COLUMN "given_name" DROP NOT NULL, ALTER COLUMN "family_name" DROP NOT NULL, ALTER COLUMN "photo_url" DROP NOT NULL;
