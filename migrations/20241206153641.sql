-- Modify "debtor_installments" table
ALTER TABLE "public"."debtor_installments" ALTER COLUMN "created_by_id" SET NOT NULL, ALTER COLUMN "debtor_transaction_id" SET NOT NULL, ALTER COLUMN "debtor_id" SET NOT NULL;
-- Modify "debtor_tenor_limits" table
ALTER TABLE "public"."debtor_tenor_limits" ALTER COLUMN "created_by_id" SET NOT NULL, ALTER COLUMN "debtor_id" SET NOT NULL;
-- Modify "debtor_to_users" table
ALTER TABLE "public"."debtor_to_users" ALTER COLUMN "debtor_id" SET NOT NULL, ALTER COLUMN "user_id" SET NOT NULL;
-- Modify "debtor_transactions" table
ALTER TABLE "public"."debtor_transactions" ALTER COLUMN "created_by_id" SET NOT NULL, ALTER COLUMN "debtor_id" SET NOT NULL;
-- Modify "debtors" table
ALTER TABLE "public"."debtors" ALTER COLUMN "created_by_id" SET NOT NULL;
