-- Create "sequences" table
CREATE TABLE "public"."sequences" (
  "id" text NOT NULL,
  "prefix" text NULL,
  "last_number" bigint NOT NULL,
  PRIMARY KEY ("id")
);
