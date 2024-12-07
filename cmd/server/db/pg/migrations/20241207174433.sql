-- Modify "members" table
ALTER TABLE "public"."members" DROP CONSTRAINT "members_pkey", ADD COLUMN "user_id" uuid NOT NULL, ADD PRIMARY KEY ("id", "user_id");
-- Modify "namespace_members" table
ALTER TABLE "public"."namespace_members" DROP CONSTRAINT "namespace_members_pkey", DROP CONSTRAINT "fk_namespace_members_member", ADD COLUMN "member_user_id" uuid NOT NULL, ADD PRIMARY KEY ("namespace_id", "member_id", "member_user_id"), ADD
 CONSTRAINT "fk_namespace_members_member" FOREIGN KEY ("member_id", "member_user_id") REFERENCES "public"."members" ("id", "user_id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Modify "service_members" table
ALTER TABLE "public"."service_members" DROP CONSTRAINT "service_members_pkey", DROP CONSTRAINT "fk_service_members_member", ADD COLUMN "member_user_id" uuid NOT NULL, ADD PRIMARY KEY ("service_id", "member_id", "member_user_id"), ADD
 CONSTRAINT "fk_service_members_member" FOREIGN KEY ("member_id", "member_user_id") REFERENCES "public"."members" ("id", "user_id") ON UPDATE NO ACTION ON DELETE NO ACTION;
