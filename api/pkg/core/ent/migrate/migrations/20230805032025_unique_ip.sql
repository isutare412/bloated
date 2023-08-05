-- Create index "banned_ips_ip_key" to table: "banned_ips"
CREATE UNIQUE INDEX "banned_ips_ip_key" ON "banned_ips" ("ip");
