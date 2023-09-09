-- create user table
CREATE TABLE IF NOT EXISTS "user" (
	"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"name" TEXT,
	"password" TEXT,
	"email" TEXT,
	"username" TEXT,
	"avatar" TEXT,
	"status" integer,
	"last_login_ip" TEXT,
	"last_login_time" DATE 
);
-- insert default data
INSERT OR IGNORE INTO "main"."user" ( "id", "name", "password", "email", "username", "avatar", "status", "last_login_ip", "last_login_time" ) VALUES ( 1, 'dbland', '$2a$10$kX5f2MW1W9kkia0p4WVjq.bJ6nfOAT0TYU0HQ1S.TKr4OSWg2HQNO', NULL, 'dbland', NULL, 1, NULL, NULL );

-- create config table
CREATE TABLE IF NOT EXISTS "config" (
	"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"type" TEXT,
	"name" TEXT,
	"host" TEXT,
	"port" TEXT,
	"username" TEXT,
	"password" TEXT,
	"db_file" TEXT
);

-- create execution log table
CREATE TABLE IF NOT EXISTS "execution_log" (
  "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  "source" TEXT,
  "database" TEXT,
  "sql" TEXT,
  "status" integer,
  "cost" integer,
  "created_time" DATE
);