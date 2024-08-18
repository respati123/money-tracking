CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE "typetransaction" AS ENUM (
  'debit',
  'credit'
);

CREATE TABLE public."users" (
  "id" SERIAL PRIMARY KEY,
  "uuid" UUID NOT NULL DEFAULT uuid_generate_v4(),
  "user_code" integer UNIQUE NOT NULL,
  "email" varchar(100) UNIQUE NOT NULL,
  "password" text,
  "phone_number" varchar(20) UNIQUE NOT NULL,
  "is_active" boolean DEFAULT false,
  "created_at" timestamp,
  "created_by" varchar(100),
  "updated_at" timestamp,
  "updated_by" varchar(100),
  "deleted_at" timestamp,
  "deleted_by" varchar(100)
);

CREATE TABLE "user_profile" (
  "id" SERIAL PRIMARY KEY,
  "uuid" UUID NOT NULL DEFAULT uuid_generate_v4(),
  "user_code" integer UNIQUE NOT NULL,
  "name" varchar(100),
  "birth_place" varchar(100),
  "birth_date" timestamp,
  "photo" text,
  "address" text,
  "created_at" timestamp NOT NULL,
  "created_by" varchar(100),
  "updated_at" timestamp,
  "updated_by" varchar(100),
  "deleted_at" timestamp,
  "deleted_by" varchar(100)
);

CREATE TABLE "category_type" (
  "id" SERIAL PRIMARY KEY,
  "uuid" UUID NOT NULL DEFAULT uuid_generate_v4(),
  "transaction_type_code" integer UNIQUE NOT NULL,
  "name" varchar(50) UNIQUE NOT NULL,
  "alias" varchar(50) UNIQUE NOT NULL,
  "created_at" timestamp NOT NULL,
  "created_by" varchar(100),
  "updated_at" timestamp,
  "updated_by" varchar(100),
  "deleted_at" timestamp,
  "deleted_by" varchar(100)
);

CREATE TABLE "transaction" (
  "id" SERIAL PRIMARY KEY,
  "uuid" UUID NOT NULL DEFAULT uuid_generate_v4(),
  "user_code" integer NOT NULL,
  "transaction_code" integer UNIQUE NOT NULL,
  "category_type_code" integer,
  "transaction_type" typetransaction NOT NULL,
  "description" text,
  "title" varchar(255),
  "created_at" timestamp NOT NULL,
  "created_by" varchar(100),
  "updated_at" timestamp,
  "updated_by" varchar(100),
  "deleted_at" timestamp,
  "deleted_by" varchar(100)
);

CREATE TABLE "otp" (
  "id" SERIAL PRIMARY KEY,
  "user_code" integer NOT NULL,
  "otp" integer UNIQUE NOT NULL,
  "expired_at" timestamp,
  "is_active" boolean,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE INDEX "user_uuid_index" ON "users" ("uuid");
CREATE INDEX "user_user_code_index" ON "users" ("user_code");
CREATE INDEX "user_email_index" ON "users" ("email");
CREATE INDEX "user_phone_number_index" ON "users" ("phone_number");
CREATE INDEX "user_is_active_index" ON "users" ("is_active");

CREATE INDEX "user_profile_user_code_index" ON "user_profile" ("user_code");

CREATE INDEX "category_type_uuid_index" ON "category_type" ("uuid");
CREATE INDEX "category_type_transaction_type_code_index" ON "category_type" ("transaction_type_code");
CREATE INDEX "category_type_name_index" ON "category_type" ("name");

CREATE INDEX "transaction_uuid_index" ON "transaction" ("uuid");
CREATE INDEX "transaction_transaction_code_index" ON "transaction" ("transaction_code");
CREATE INDEX "transaction_category_type_code_index" ON "transaction" ("category_type_code");
CREATE INDEX "transaction_transaction_type_index" ON "transaction" ("transaction_type");
CREATE INDEX "transaction_created_at_index" ON "transaction" ("created_at");

CREATE INDEX "otp_user_code_index" ON "otp" ("user_code");
CREATE INDEX "otp_otp_index" ON "otp" ("otp");
CREATE INDEX "otp_is_active_index" ON "otp" ("is_active");

COMMENT ON COLUMN "users"."uuid" IS 'this key for update, deleted method';
COMMENT ON COLUMN "users"."password" IS 'Password would be hashed';
COMMENT ON COLUMN "user_profile"."uuid" IS 'this key for update, deleted method';
COMMENT ON COLUMN "category_type"."uuid" IS 'this key for updated, delete method';
COMMENT ON COLUMN "transaction"."uuid" IS 'this key for updated, delete method';
COMMENT ON COLUMN "transaction"."transaction_code" IS 'this key for get transaction by code';

ALTER TABLE "transaction" ADD FOREIGN KEY ("category_type_code") REFERENCES "category_type" ("transaction_type_code");
ALTER TABLE "user_profile" ADD FOREIGN KEY ("user_code") REFERENCES "users" ("user_code") ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE "transaction" ADD FOREIGN KEY ("user_code") REFERENCES "users" ("user_code") ON DELETE CASCADE ON UPDATE NO ACTION;
ALTER TABLE "otp" ADD FOREIGN KEY ("user_code") REFERENCES "users" ("user_code") ON DELETE CASCADE ON UPDATE NO ACTION;