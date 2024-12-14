-- Create "members" table
CREATE TABLE "public"."members" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "user_id" uuid NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_members_deleted_at" to table: "members"
CREATE INDEX "idx_members_deleted_at" ON "public"."members" ("deleted_at");
-- Create "namespaces" table
CREATE TABLE "public"."namespaces" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "title" text NULL,
  "slug" text NULL,
  "description" text NULL,
  "owner" uuid NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_namespaces_deleted_at" to table: "namespaces"
CREATE INDEX "idx_namespaces_deleted_at" ON "public"."namespaces" ("deleted_at");
-- Create "ns_members" table
CREATE TABLE "public"."ns_members" (
  "id" text NOT NULL,
  "user_id" text NULL,
  PRIMARY KEY ("id")
);
-- Create "namespace_members" table
CREATE TABLE "public"."namespace_members" (
  "namespace_id" uuid NOT NULL,
  "ns_member_id" text NOT NULL,
  PRIMARY KEY ("namespace_id", "ns_member_id"),
  CONSTRAINT "fk_namespace_members_namespace" FOREIGN KEY ("namespace_id") REFERENCES "public"."namespaces" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_namespace_members_ns_member" FOREIGN KEY ("ns_member_id") REFERENCES "public"."ns_members" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
