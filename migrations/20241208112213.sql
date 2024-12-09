-- Modify "debtor_installments" table
ALTER TABLE "public"."debtor_installments" DROP COLUMN "tenor_limit_type", DROP COLUMN "tenor_duration", ADD COLUMN "tenor_limit_id" uuid NOT NULL, ADD
 CONSTRAINT "fk_debtor_installments_debtor_tenor_limit" FOREIGN KEY ("tenor_limit_id") REFERENCES "public"."debtor_tenor_limits" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
