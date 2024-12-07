-- Create "members" table
CREATE TABLE "public"."members" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
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
-- Create "namespace_members" table
CREATE TABLE "public"."namespace_members" (
  "namespace_id" uuid NOT NULL,
  "member_id" uuid NOT NULL,
  PRIMARY KEY ("namespace_id", "member_id"),
  CONSTRAINT "fk_namespace_members_member" FOREIGN KEY ("member_id") REFERENCES "public"."members" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_namespace_members_namespace" FOREIGN KEY ("namespace_id") REFERENCES "public"."namespaces" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "services" table
CREATE TABLE "public"."services" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "app" text NULL,
  "namespace_id" uuid NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_namespaces_services" FOREIGN KEY ("namespace_id") REFERENCES "public"."namespaces" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create index "idx_services_deleted_at" to table: "services"
CREATE INDEX "idx_services_deleted_at" ON "public"."services" ("deleted_at");
-- Create "service_members" table
CREATE TABLE "public"."service_members" (
  "service_id" uuid NOT NULL,
  "member_id" uuid NOT NULL,
  PRIMARY KEY ("service_id", "member_id"),
  CONSTRAINT "fk_service_members_member" FOREIGN KEY ("member_id") REFERENCES "public"."members" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_service_members_service" FOREIGN KEY ("service_id") REFERENCES "public"."services" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
