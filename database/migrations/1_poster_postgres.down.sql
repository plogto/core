-- Foreign keys
ALTER TABLE "password" DROP CONSTRAINT IF EXISTS "password_fk0";
ALTER TABLE "connection" DROP CONSTRAINT IF EXISTS "connection_fk0";
ALTER TABLE "connection" DROP CONSTRAINT IF EXISTS "connection_fk1";
ALTER TABLE "post_tag" DROP CONSTRAINT IF EXISTS "post_tag_fk0";
ALTER TABLE "post_tag" DROP CONSTRAINT IF EXISTS "post_tag_fk1";
ALTER TABLE "post_like" DROP CONSTRAINT IF EXISTS "post_like_fk0";
ALTER TABLE "post_like" DROP CONSTRAINT IF EXISTS "post_like_fk1";
ALTER TABLE "post_save" DROP CONSTRAINT IF EXISTS "post_save_fk0";
ALTER TABLE "post_save" DROP CONSTRAINT IF EXISTS "post_save_fk1";
ALTER TABLE "post_comment" DROP CONSTRAINT IF EXISTS "post_comment_fk0";
ALTER TABLE "post_comment" DROP CONSTRAINT IF EXISTS "post_comment_fk1";
ALTER TABLE "post_comment" DROP CONSTRAINT IF EXISTS "post_comment_fk2";
ALTER TABLE "post_comment_like" DROP CONSTRAINT IF EXISTS "post_comment_like_fk0";
ALTER TABLE "post_comment_like" DROP CONSTRAINT IF EXISTS "post_comment_like_fk1";

-- Tables
DROP TABLE IF EXISTS "user";
DROP TABLE IF EXISTS "password";
DROP TABLE IF EXISTS "post";
DROP TABLE IF EXISTS "connection";
DROP TABLE IF EXISTS "tag";
DROP TABLE IF EXISTS "post_tag";
DROP TABLE IF EXISTS "post_like";
DROP TABLE IF EXISTS "post_save";
DROP TABLE IF EXISTS "post_comment";
DROP TABLE IF EXISTS "post_comment_like";

-- Types
DROP TYPE IF EXISTS "user_roles";
