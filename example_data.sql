/*
 Navicat Premium Data Transfer

 Source Server         : mnc
 Source Server Type    : PostgreSQL
 Source Server Version : 140016 (140016)
 Source Host           : localhost:5432
 Source Catalog        : mnc
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140016 (140016)
 File Encoding         : 65001

 Date: 14/02/2025 22:56:01
*/


-- ----------------------------
-- Table structure for schema_migrations
-- ----------------------------
DROP TABLE IF EXISTS "public"."schema_migrations";
CREATE TABLE "public"."schema_migrations" (
  "version" int8 NOT NULL,
  "dirty" bool NOT NULL
)
;

-- ----------------------------
-- Records of schema_migrations
-- ----------------------------
INSERT INTO "public"."schema_migrations" VALUES (3, 'f');

-- ----------------------------
-- Table structure for transactions
-- ----------------------------
DROP TABLE IF EXISTS "public"."transactions";
CREATE TABLE "public"."transactions" (
  "transaction_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL,
  "target_user" uuid,
  "transaction_type" varchar(10) COLLATE "pg_catalog"."default" NOT NULL,
  "amount" int8 NOT NULL,
  "remarks" text COLLATE "pg_catalog"."default",
  "balance_before" int8 NOT NULL,
  "balance_after" int8 NOT NULL,
  "created_date" timestamp(6) DEFAULT now()
)
;

-- ----------------------------
-- Records of transactions
-- ----------------------------
INSERT INTO "public"."transactions" VALUES ('248f7eb6-0337-4ae9-85b3-ae86fa8e1a09', 'cd4a9127-b2d5-490c-b856-b519432e015b', NULL, 'CREDIT', 100000, '', 0, 100000, '2025-02-14 22:53:52.849341');
INSERT INTO "public"."transactions" VALUES ('7a4ed63c-7f7d-46fd-90f2-e51ebc983b16', 'cd4a9127-b2d5-490c-b856-b519432e015b', NULL, 'DEBIT', 100000, 'Pulsa Telkomsel 100k', 100000, 0, '2025-02-14 22:54:22.388804');
INSERT INTO "public"."transactions" VALUES ('d6151f84-cf51-402a-b298-c8feff62e5ec', 'cd4a9127-b2d5-490c-b856-b519432e015b', NULL, 'CREDIT', 100000, '', 0, 100000, '2025-02-14 22:55:01.605408');
INSERT INTO "public"."transactions" VALUES ('55fb7ea0-9fab-45d0-8faa-7c03f8259496', 'cd4a9127-b2d5-490c-b856-b519432e015b', '855d6d6f-694c-4071-98d0-187c854204f2', 'DEBIT', 50000, 'Pulsa Telkomsel 50k', 100000, 50000, '2025-02-14 22:55:15.564701');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "first_name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "last_name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "phone_number" varchar(20) COLLATE "pg_catalog"."default" NOT NULL,
  "address" text COLLATE "pg_catalog"."default",
  "pin" text COLLATE "pg_catalog"."default",
  "created_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6),
  "deleted_at" timestamp(6)
)
;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO "public"."users" VALUES ('cd4a9127-b2d5-490c-b856-b519432e015b', 'Bambang', 'Persija', '082240119661', 'Jl. H GG Y No.1', '$2a$10$epnf9NCslIIMo0EIMWtmku8haj1fytre3Q6xbKY/qyV3b3LZtoBGK', '2025-02-14 22:53:01.410339', NULL, NULL);
INSERT INTO "public"."users" VALUES ('855d6d6f-694c-4071-98d0-187c854204f2', 'Atep', 'Persib', '082240119662', 'Jl. H GG Y No.1', '$2a$10$vlVSqkdEMteV9f2Lo2szA.aFpGx3TeTJOvCV8jMR9O/2dUJjnzzQi', '2025-02-14 22:53:39.670197', NULL, NULL);

-- ----------------------------
-- Primary Key structure for table schema_migrations
-- ----------------------------
ALTER TABLE "public"."schema_migrations" ADD CONSTRAINT "schema_migrations_pkey" PRIMARY KEY ("version");

-- ----------------------------
-- Indexes structure for table transactions
-- ----------------------------
CREATE INDEX "idx_transactions_target_user" ON "public"."transactions" USING btree (
  "target_user" "pg_catalog"."uuid_ops" ASC NULLS LAST
);
CREATE INDEX "idx_transactions_transaction_type" ON "public"."transactions" USING btree (
  "transaction_type" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_transactions_user_id" ON "public"."transactions" USING btree (
  "user_id" "pg_catalog"."uuid_ops" ASC NULLS LAST
);

-- ----------------------------
-- Checks structure for table transactions
-- ----------------------------
ALTER TABLE "public"."transactions" ADD CONSTRAINT "transactions_transaction_type_check" CHECK (transaction_type::text = ANY (ARRAY['CREDIT'::character varying, 'DEBIT'::character varying]::text[]));

-- ----------------------------
-- Primary Key structure for table transactions
-- ----------------------------
ALTER TABLE "public"."transactions" ADD CONSTRAINT "transactions_pkey" PRIMARY KEY ("transaction_id");

-- ----------------------------
-- Uniques structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "users_phone_number_key" UNIQUE ("phone_number");

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("id");
