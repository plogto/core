-- Extentions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Enums
DROP TYPE IF EXISTS "user_role";
CREATE TYPE user_role AS ENUM ('USER', 'ADMIN');

DROP TYPE IF EXISTS "post_status";
CREATE TYPE post_status AS ENUM ('PUBLIC', 'PRIVATE');

-- Tables
CREATE TABLE "user" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"username" VARCHAR(32) NOT NULL UNIQUE,
	"email" VARCHAR(100) NOT NULL UNIQUE,
	"full_name" VARCHAR(64) NOT NULL,
	"avatar" TEXT DEFAULT NULL,
	"background" TEXT DEFAULT NULL,
	"bio" TEXT DEFAULT NULL,
	"role" user_role NOT NULL DEFAULT 'USER',
	"is_private" BOOLEAN NOT NULL DEFAULT FALSE,
  "created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
	CONSTRAINT "user_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "password" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"user_id" uuid NOT NULL,
	"password" TEXT NOT NULL,
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
 CONSTRAINT "password_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "post" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"user_id" uuid NOT NULL,
	"parent_id" uuid DEFAULT NULL,
	"child_id" uuid DEFAULT NULL,
	"content" TEXT NOT NULL,
	"attachment" TEXT DEFAULT NULL,
	"url" VARCHAR(20) NOT NULL UNIQUE,
	"status" post_status NOT NULL DEFAULT 'PUBLIC',
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
 CONSTRAINT "post_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

-- status => 1: pending, 2: following
CREATE TABLE "connection" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"follower_id" uuid NOT NULL,
	"following_id" uuid NOT NULL,
	"status" INTEGER NOT NULL DEFAULT 1,
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
 CONSTRAINT "connection_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "tag" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"name" VARCHAR(100) NOT NULL UNIQUE,
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
 CONSTRAINT "tag_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "post_tag" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"tag_id" uuid NOT NULL,
	"post_id" uuid NOT NULL,
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
 CONSTRAINT "post_tag_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "post_like" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"user_id" uuid NOT NULL,
	"post_id" uuid NOT NULL,
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
 CONSTRAINT "post_like_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "post_save" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"user_id" uuid NOT NULL,
	"post_id" uuid NOT NULL,
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
 CONSTRAINT "post_save_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "online_user" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"user_id" uuid NOT NULL,
	"socket_id" VARCHAR(20) NOT NULL,
	"token" TEXT NOT NULL,
	"user_agent" TEXT NOT NULL,
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
 CONSTRAINT "online_user_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "notification_type" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"name" VARCHAR(100) NOT NULL,
	"template" TEXT NOT NULL,
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
 CONSTRAINT "notification_type_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "notification" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"notification_type_id" uuid NOT NULL,
	"sender_id" uuid NOT NULL,
	"receiver_id" uuid NOT NULL,
	"post_id" uuid DEFAULT NULL,
	"reply_id" uuid DEFAULT NULL,
	"url" TEXT NOT NULL,
	"read" BOOLEAN NOT NULL DEFAULT FALSE,
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
 CONSTRAINT "notification_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "file" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"hash" TEXT NOT NULL UNIQUE,
	"name" TEXT NOT NULL UNIQUE,
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
 CONSTRAINT "file_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

-- Triggers
CREATE OR REPLACE FUNCTION trigger_set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at= now();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_user BEFORE UPDATE ON "user" FOR EACH ROW EXECUTE PROCEDURE  trigger_set_updated_at();
CREATE TRIGGER update_password BEFORE UPDATE ON "password" FOR EACH ROW EXECUTE PROCEDURE  trigger_set_updated_at();
CREATE TRIGGER update_post BEFORE UPDATE ON "post" FOR EACH ROW EXECUTE PROCEDURE  trigger_set_updated_at();
CREATE TRIGGER update_connection BEFORE UPDATE ON "connection" FOR EACH ROW EXECUTE PROCEDURE  trigger_set_updated_at();
CREATE TRIGGER update_tag BEFORE UPDATE ON "tag" FOR EACH ROW EXECUTE PROCEDURE  trigger_set_updated_at();
CREATE TRIGGER update_post_tag BEFORE UPDATE ON "post_tag" FOR EACH ROW EXECUTE PROCEDURE  trigger_set_updated_at();
CREATE TRIGGER update_post_like BEFORE UPDATE ON "post_like" FOR EACH ROW EXECUTE PROCEDURE  trigger_set_updated_at();
CREATE TRIGGER update_post_save BEFORE UPDATE ON "post_save" FOR EACH ROW EXECUTE PROCEDURE  trigger_set_updated_at();
CREATE TRIGGER update_online_user BEFORE UPDATE ON "online_user" FOR EACH ROW EXECUTE PROCEDURE  trigger_set_updated_at();
CREATE TRIGGER update_notification_type BEFORE UPDATE ON "notification_type" FOR EACH ROW EXECUTE PROCEDURE  trigger_set_updated_at();
CREATE TRIGGER update_notification BEFORE UPDATE ON "notification" FOR EACH ROW EXECUTE PROCEDURE  trigger_set_updated_at();
CREATE TRIGGER update_file BEFORE UPDATE ON "file" FOR EACH ROW EXECUTE PROCEDURE  trigger_set_updated_at();

-- Foreign keys
ALTER TABLE "password" ADD CONSTRAINT "password_fk0" FOREIGN KEY ("user_id") REFERENCES "user"("id");
ALTER TABLE "post" ADD CONSTRAINT "post_fk0" FOREIGN KEY ("user_id") REFERENCES "user"("id");
ALTER TABLE "post" ADD CONSTRAINT "post_fk1" FOREIGN KEY ("parent_id") REFERENCES "post"("id");
ALTER TABLE "post" ADD CONSTRAINT "post_fk2" FOREIGN KEY ("child_id") REFERENCES "post"("id");
ALTER TABLE "connection" ADD CONSTRAINT "connection_fk1" FOREIGN KEY ("follower_id") REFERENCES "user"("id");
ALTER TABLE "connection" ADD CONSTRAINT "connection_fk0" FOREIGN KEY ("following_id") REFERENCES "user"("id");
ALTER TABLE "post_tag" ADD CONSTRAINT "post_tag_fk0" FOREIGN KEY ("tag_id") REFERENCES "tag"("id");
ALTER TABLE "post_tag" ADD CONSTRAINT "post_tag_fk1" FOREIGN KEY ("post_id") REFERENCES "post"("id");
ALTER TABLE "post_like" ADD CONSTRAINT "post_like_fk0" FOREIGN KEY ("user_id") REFERENCES "user"("id");
ALTER TABLE "post_like" ADD CONSTRAINT "post_like_fk1" FOREIGN KEY ("post_id") REFERENCES "post"("id");
ALTER TABLE "post_save" ADD CONSTRAINT "post_save_fk0" FOREIGN KEY ("user_id") REFERENCES "user"("id");
ALTER TABLE "post_save" ADD CONSTRAINT "post_save_fk1" FOREIGN KEY ("post_id") REFERENCES "post"("id");
ALTER TABLE "online_user" ADD CONSTRAINT "online_user_fk0" FOREIGN KEY ("user_id") REFERENCES "user"("id");
ALTER TABLE "notification" ADD CONSTRAINT "notification_fk0" FOREIGN KEY ("notification_type_id") REFERENCES "notification_type"("id");
ALTER TABLE "notification" ADD CONSTRAINT "notification_fk1" FOREIGN KEY ("sender_id") REFERENCES "user"("id");
ALTER TABLE "notification" ADD CONSTRAINT "notification_fk2" FOREIGN KEY ("receiver_id") REFERENCES "user"("id");
ALTER TABLE "notification" ADD CONSTRAINT "notification_fk3" FOREIGN KEY ("post_id") REFERENCES "post"("id");
ALTER TABLE "notification" ADD CONSTRAINT "notification_fk4" FOREIGN KEY ("reply_id") REFERENCES "post"("id");