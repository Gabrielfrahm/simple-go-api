SELECT 'CREATE DATABASE goapidb'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'goapidb');

-- Create Table Users
DROP TABLE IF EXISTS users;
CREATE TABLE users(
   	"id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "name" VARCHAR(50) NOT NULL,
    "nick" VARCHAR(50) NOT NULL,
    "email" VARCHAR(50) NOT NULL,
    "password" VARCHAR(250) NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- CreateIndex
CREATE UNIQUE INDEX "users_email_key" ON "users"("email");
CREATE UNIQUE INDEX "users_nick_key" ON "users"("nick");
