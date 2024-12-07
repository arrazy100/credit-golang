-- Create "seed_versions" table
CREATE TABLE "public"."seed_versions" (
  "version" bigint NOT NULL
);
-- Create "users" table
CREATE TABLE "public"."users" (
  "id" uuid NOT NULL,
  "email" character varying(255) NOT NULL,
  "password" character varying(255) NOT NULL,
  "role" bigint NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_users_email" UNIQUE ("email")
);
-- Create "debtor_transactions" table
CREATE TABLE "public"."debtor_transactions" (
  "id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by_id" uuid NOT NULL,
  "updated_by_id" uuid NULL,
  "user_id" uuid NOT NULL,
  "contract_number" character varying(255) NOT NULL,
  "otr" numeric(18,2) NOT NULL,
  "admin_fee" numeric(18,2) NOT NULL,
  "total_loan" numeric(18,2) NOT NULL,
  "total_interest" numeric(18,2) NOT NULL,
  "asset_name" character varying(255) NOT NULL,
  "status" bigint NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_debtor_transactions_created_by" FOREIGN KEY ("created_by_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_debtor_transactions_updated_by" FOREIGN KEY ("updated_by_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_debtor_transactions_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "debtor_installments" table
CREATE TABLE "public"."debtor_installments" (
  "id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by_id" uuid NOT NULL,
  "updated_by_id" uuid NULL,
  "user_id" uuid NOT NULL,
  "debtor_transaction_id" uuid NOT NULL,
  "tenor_limit_type" smallint NOT NULL,
  "tenor_duration" bigint NOT NULL,
  "monthly_installment" numeric(18,2) NOT NULL,
  "total_installment" numeric(10,2) NOT NULL,
  "start_date_period" date NOT NULL,
  "end_date_period" date NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_debtor_installments_created_by" FOREIGN KEY ("created_by_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_debtor_installments_debtor_transaction" FOREIGN KEY ("debtor_transaction_id") REFERENCES "public"."debtor_transactions" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_debtor_installments_updated_by" FOREIGN KEY ("updated_by_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_debtor_installments_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "debtors" table
CREATE TABLE "public"."debtors" (
  "id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by_id" uuid NOT NULL,
  "updated_by_id" uuid NULL,
  "user_id" uuid NOT NULL,
  "nik" text NOT NULL,
  "full_name" text NOT NULL,
  "legal_name" text NOT NULL,
  "place_of_birth" text NOT NULL,
  "date_of_birth" timestamptz NOT NULL,
  "salary" numeric(18,2) NOT NULL,
  "identity_picture_url" text NOT NULL,
  "selfie_picture_url" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_debtors_nik" UNIQUE ("nik"),
  CONSTRAINT "fk_debtors_created_by" FOREIGN KEY ("created_by_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_debtors_updated_by" FOREIGN KEY ("updated_by_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_debtors_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "debtor_tenor_limits" table
CREATE TABLE "public"."debtor_tenor_limits" (
  "id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by_id" uuid NOT NULL,
  "updated_by_id" uuid NULL,
  "user_id" uuid NOT NULL,
  "debtor_id" uuid NOT NULL,
  "tenor_limit_type" bigint NOT NULL,
  "tenor_duration" bigint NOT NULL,
  "limit_amount" numeric(18,2) NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_debtor_tenor_limits_created_by" FOREIGN KEY ("created_by_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_debtor_tenor_limits_updated_by" FOREIGN KEY ("updated_by_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_debtor_tenor_limits_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_debtors_tenor_limits" FOREIGN KEY ("debtor_id") REFERENCES "public"."debtors" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
