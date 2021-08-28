ALTER TABLE "password" DROP CONSTRAINT IF EXISTS "password_fk0";

ALTER TABLE "connection" DROP CONSTRAINT IF EXISTS "connection_fk0";
ALTER TABLE "connection" DROP CONSTRAINT IF EXISTS "connection_fk1";

DROP TABLE IF EXISTS "user";

DROP TABLE IF EXISTS "password";

DROP TABLE IF EXISTS "post";

DROP TABLE IF EXISTS "connection";

DROP TYPE IF EXISTS "user_roles";
