-- Modify "banned_ips" table
ALTER TABLE "banned_ips" ALTER COLUMN "country" DROP NOT NULL;
