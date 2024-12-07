-- Modify "debtor_tenor_limits" table
ALTER TABLE "public"."debtor_tenor_limits" ADD COLUMN "current_limit" numeric(18,2) NOT NULL;
-- Rename a column from "limit_amount" to "total_limit"
ALTER TABLE "public"."debtor_tenor_limits" RENAME COLUMN "limit_amount" TO "total_limit";
