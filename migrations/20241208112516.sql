-- Create "debtor_installment_lines" table
CREATE TABLE "public"."debtor_installment_lines" (
  "id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by_id" uuid NOT NULL,
  "updated_by_id" uuid NULL,
  "user_id" uuid NOT NULL,
  "debtor_installment_id" uuid NOT NULL,
  "due_date" date NOT NULL,
  "installment_number" bigint NOT NULL,
  "installment_amount" numeric(18,2) NOT NULL,
  "payment_date" date NULL,
  "status" bigint NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_debtor_installment_lines_created_by" FOREIGN KEY ("created_by_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_debtor_installment_lines_updated_by" FOREIGN KEY ("updated_by_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_debtor_installment_lines_user" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_debtor_installments_installment_lines" FOREIGN KEY ("debtor_installment_id") REFERENCES "public"."debtor_installments" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
