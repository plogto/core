-- Foreign keys
ALTER TABLE "user" DROP CONSTRAINT IF EXISTS "user_fk0";
ALTER TABLE "user" DROP CONSTRAINT IF EXISTS "user_fk1";
ALTER TABLE "password" DROP CONSTRAINT IF EXISTS "password_fk0";
ALTER TABLE "connection" DROP CONSTRAINT IF EXISTS "connection_fk0";
ALTER TABLE "connection" DROP CONSTRAINT IF EXISTS "connection_fk1";
ALTER TABLE "post" DROP CONSTRAINT IF EXISTS "post_fk0";
ALTER TABLE "post" DROP CONSTRAINT IF EXISTS "post_fk1";
ALTER TABLE "post" DROP CONSTRAINT IF EXISTS "post_fk2";
ALTER TABLE "post_tag" DROP CONSTRAINT IF EXISTS "post_tag_fk0";
ALTER TABLE "post_tag" DROP CONSTRAINT IF EXISTS "post_tag_fk1";
ALTER TABLE "post_attachment" DROP CONSTRAINT IF EXISTS "post_attachment_fk0";
ALTER TABLE "post_attachment" DROP CONSTRAINT IF EXISTS "post_attachment_fk1";
ALTER TABLE "post_like" DROP CONSTRAINT IF EXISTS "post_like_fk0";
ALTER TABLE "post_like" DROP CONSTRAINT IF EXISTS "post_like_fk1";
ALTER TABLE "post_save" DROP CONSTRAINT IF EXISTS "post_save_fk0";
ALTER TABLE "post_save" DROP CONSTRAINT IF EXISTS "post_save_fk1";
ALTER TABLE "online_user" DROP CONSTRAINT IF EXISTS "online_user_fk0";
ALTER TABLE "notification" DROP CONSTRAINT IF EXISTS "notification_fk0";
ALTER TABLE "notification" DROP CONSTRAINT IF EXISTS "notification_fk1";
ALTER TABLE "notification" DROP CONSTRAINT IF EXISTS "notification_fk2";
ALTER TABLE "notification" DROP CONSTRAINT IF EXISTS "notification_fk3";
ALTER TABLE "notification" DROP CONSTRAINT IF EXISTS "notification_fk4";

-- Tables
DROP TABLE IF EXISTS "user";
DROP TABLE IF EXISTS "password";
DROP TABLE IF EXISTS "post";
DROP TABLE IF EXISTS "connection";
DROP TABLE IF EXISTS "tag";
DROP TABLE IF EXISTS "post_tag";
DROP TABLE IF EXISTS "post_attachment";
DROP TABLE IF EXISTS "post_like";
DROP TABLE IF EXISTS "post_save";
DROP TABLE IF EXISTS "online_user";
DROP TABLE IF EXISTS "notification_type";
DROP TABLE IF EXISTS "notification";
DROP TABLE IF EXISTS "file";

-- Types
DROP TYPE IF EXISTS "user_roles";
