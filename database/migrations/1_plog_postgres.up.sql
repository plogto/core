-- Extentions
CREATE EXTENSION IF NOT EXISTS "UUID-ossp";

-- Enums
DROP TYPE IF EXISTS "user_role";
CREATE TYPE user_role AS ENUM ('USER', 'ADMIN');

DROP TYPE IF EXISTS "post_status";
CREATE TYPE post_status AS ENUM ('PUBLIC', 'PRIVATE');

DROP TYPE IF EXISTS "background_color";
CREATE TYPE background_color AS ENUM ('LIGHT', 'DIM', 'DARK');

DROP TYPE IF EXISTS "primary_color";
CREATE TYPE primary_color AS ENUM ('BLUE', 'GREEN', 'RED', 'PURPLE', 'ORANGE', 'YELLOW');

-- Tables
CREATE TABLE "user" (
	"id" UUID NOT NULL DEFAULT uuid_generate_v4(),
	"username" VARCHAR(32) NOT NULL UNIQUE,
	"email" VARCHAR(100) NOT NULL UNIQUE,
	"full_name" VARCHAR(64) NOT NULL,
	"avatar" UUID DEFAULT NULL,
	"background_color" background_color DEFAULT 'LIGHT',
	"primary_color" primary_color DEFAULT 'BLUE',
	"background" UUID DEFAULT NULL,
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
	"id" UUID NOT NULL DEFAULT uuid_generate_v4(),
	"user_id" UUID NOT NULL,
	"password" TEXT NOT NULL,
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
 CONSTRAINT "password_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "post" (
	"id" UUID NOT NULL DEFAULT uuid_generate_v4(),
	"user_id" UUID NOT NULL,
	"parent_id" UUID DEFAULT NULL,
	"child_id" UUID DEFAULT NULL,
	"content" TEXT DEFAULT NULL,
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
	"id" UUID NOT NULL DEFAULT uuid_generate_v4(),
	"follower_id" UUID NOT NULL,
	"following_id" UUID NOT NULL,
	"status" INTEGER NOT NULL DEFAULT 1,
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
 CONSTRAINT "connection_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "tag" (
	"id" UUID NOT NULL DEFAULT uuid_generate_v4(),
	"name" VARCHAR(100) NOT NULL UNIQUE,
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
 CONSTRAINT "tag_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "post_tag" (
	"id" UUID NOT NULL DEFAULT uuid_generate_v4(),
	"tag_id" UUID NOT NULL,
	"post_id" UUID NOT NULL,
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
 CONSTRAINT "post_tag_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "post_attachment" (
	"id" UUID NOT NULL DEFAULT uuid_generate_v4(),
	"post_id" UUID NOT NULL,
	"file_id" UUID NOT NULL,
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
 CONSTRAINT "post_attachment_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "post_like" (
	"id" UUID NOT NULL DEFAULT uuid_generate_v4(),
	"user_id" UUID NOT NULL,
	"post_id" UUID NOT NULL,
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
 CONSTRAINT "post_like_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "post_save" (
	"id" UUID NOT NULL DEFAULT uuid_generate_v4(),
	"user_id" UUID NOT NULL,
	"post_id" UUID NOT NULL,
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
 CONSTRAINT "post_save_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "online_user" (
	"id" UUID NOT NULL DEFAULT uuid_generate_v4(),
	"user_id" UUID NOT NULL,
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
	"id" UUID NOT NULL DEFAULT uuid_generate_v4(),
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
	"id" UUID NOT NULL DEFAULT uuid_generate_v4(),
	"notification_type_id" UUID NOT NULL,
	"sender_id" UUID NOT NULL,
	"receiver_id" UUID NOT NULL,
	"post_id" UUID DEFAULT NULL,
	"reply_id" UUID DEFAULT NULL,
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
	"id" UUID NOT NULL DEFAULT uuid_generate_v4(),
	"hash" TEXT NOT NULL UNIQUE,
	"name" TEXT NOT NULL UNIQUE,
	"width" SMALLINT NOT NULL,
	"height" SMALLINT NOT NULL,
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
CREATE TRIGGER update_post_attachment BEFORE UPDATE ON "post_attachment" FOR EACH ROW EXECUTE PROCEDURE  trigger_set_updated_at();
CREATE TRIGGER update_post_like BEFORE UPDATE ON "post_like" FOR EACH ROW EXECUTE PROCEDURE  trigger_set_updated_at();
CREATE TRIGGER update_post_save BEFORE UPDATE ON "post_save" FOR EACH ROW EXECUTE PROCEDURE  trigger_set_updated_at();
CREATE TRIGGER update_online_user BEFORE UPDATE ON "online_user" FOR EACH ROW EXECUTE PROCEDURE  trigger_set_updated_at();
CREATE TRIGGER update_notification_type BEFORE UPDATE ON "notification_type" FOR EACH ROW EXECUTE PROCEDURE  trigger_set_updated_at();
CREATE TRIGGER update_notification BEFORE UPDATE ON "notification" FOR EACH ROW EXECUTE PROCEDURE  trigger_set_updated_at();
CREATE TRIGGER update_file BEFORE UPDATE ON "file" FOR EACH ROW EXECUTE PROCEDURE  trigger_set_updated_at();

-- Foreign keys
ALTER TABLE "user" ADD CONSTRAINT "user_fk0" FOREIGN KEY ("avatar") REFERENCES "file"("id");
ALTER TABLE "user" ADD CONSTRAINT "user_fk1" FOREIGN KEY ("background") REFERENCES "file"("id");
ALTER TABLE "password" ADD CONSTRAINT "password_fk0" FOREIGN KEY ("user_id") REFERENCES "user"("id");
ALTER TABLE "post" ADD CONSTRAINT "post_fk0" FOREIGN KEY ("user_id") REFERENCES "user"("id");
ALTER TABLE "post" ADD CONSTRAINT "post_fk1" FOREIGN KEY ("parent_id") REFERENCES "post"("id");
ALTER TABLE "post" ADD CONSTRAINT "post_fk2" FOREIGN KEY ("child_id") REFERENCES "post"("id");
ALTER TABLE "connection" ADD CONSTRAINT "connection_fk1" FOREIGN KEY ("follower_id") REFERENCES "user"("id");
ALTER TABLE "connection" ADD CONSTRAINT "connection_fk0" FOREIGN KEY ("following_id") REFERENCES "user"("id");
ALTER TABLE "post_tag" ADD CONSTRAINT "post_tag_fk0" FOREIGN KEY ("tag_id") REFERENCES "tag"("id");
ALTER TABLE "post_tag" ADD CONSTRAINT "post_tag_fk1" FOREIGN KEY ("post_id") REFERENCES "post"("id");
ALTER TABLE "post_attachment" ADD CONSTRAINT "post_attachment_fk0" FOREIGN KEY ("post_id") REFERENCES "post"("id");
ALTER TABLE "post_attachment" ADD CONSTRAINT "post_attachment_fk1" FOREIGN KEY ("file_id") REFERENCES "file"("id");
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