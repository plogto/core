-- Extentions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Enums 
DROP TYPE IF EXISTS "user_roles";
CREATE TYPE user_roles AS ENUM ('USER', 'ADMIN');

-- Tables
CREATE TABLE "user" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"username" VARCHAR(32) NOT NULL UNIQUE,
	"email" VARCHAR(100) NOT NULL UNIQUE,
	"fullname" VARCHAR(64) NOT NULL,
	"role" user_roles DEFAULT 'USER',
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
	"content" TEXT NOT NULL,
	"status" INTEGER NOT NULL DEFAULT 0,
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
 CONSTRAINT "post_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);

CREATE TABLE "follower" (
	"id" uuid NOT NULL DEFAULT uuid_generate_v4(),
	"user_id" uuid NOT NULL,
	"follower_id" uuid NOT NULL,
	"status" INTEGER NOT NULL DEFAULT 0,
	"created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
	"deleted_at" TIMESTAMP,
 CONSTRAINT "follower_pk" PRIMARY KEY ("id")
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
CREATE TRIGGER update_follower BEFORE UPDATE ON "follower" FOR EACH ROW EXECUTE PROCEDURE  trigger_set_updated_at();

-- Foreign keys
ALTER TABLE "password" ADD CONSTRAINT "password_fk0" FOREIGN KEY ("user_id") REFERENCES "user"("id");
ALTER TABLE "post" ADD CONSTRAINT "post_fk0" FOREIGN KEY ("user_id") REFERENCES "user"("id");
ALTER TABLE "follower" ADD CONSTRAINT "follower_fk0" FOREIGN KEY ("user_id") REFERENCES "user"("id");
ALTER TABLE "follower" ADD CONSTRAINT "follower_fk1" FOREIGN KEY ("follower_id") REFERENCES "user"("id");