CREATE EXTENSION IF NOT EXISTS "citext";

CREATE TABLE "organizations" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "owner_identity_id" uuid,
  "display_name" varchar,
  "slug" citext,
  "settings" jsonb,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "identities" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "org_id" uuid,
  "email" citext,
  "display_name" varchar,
  "password_hash" varchar,
  "verified" bool DEFAULT false,
  "active" bool DEFAULT true,
  "failed_login_attempts" int DEFAULT 0,
  "locked_until" timestamptz,
  "last_login_at" timestamptz,
  "metadata" jsonb,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "api_keys" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "org_id" uuid,
  "name" varchar,
  "key_prefix" varchar,
  "key_hash" varchar,
  "last_used_at" timestamptz,
  "expires_at" timestamptz,
  "revoked_at" timestamptz,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "identity_id" uuid,
  "device_info" varchar,
  "user_agent" varchar,
  "token_hash" varchar,
  "ip_address" varchar,
  "last_active_at" timestamptz,
  "expires_at" timestamptz,
  "revoked_at" timestamptz,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "tokens" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "identity_id" uuid,
  "token_type" varchar,
  "token_hash" varchar,
  "expires_at" timestamptz,
  "used_at" timestamptz,
  "created_at" timestamptz DEFAULT (now())
);

CREATE UNIQUE INDEX ON "organizations" ("slug");
CREATE UNIQUE INDEX ON "identities" ("org_id", "email");
CREATE UNIQUE INDEX ON "api_keys" ("key_hash");
CREATE UNIQUE INDEX ON "sessions" ("token_hash");
CREATE UNIQUE INDEX ON "tokens" ("token_hash");

ALTER TABLE "organizations" ADD FOREIGN KEY ("owner_identity_id") REFERENCES "identities" ("id") DEFERRABLE INITIALLY DEFERRED;
ALTER TABLE "identities" ADD FOREIGN KEY ("org_id") REFERENCES "organizations" ("id") DEFERRABLE INITIALLY IMMEDIATE;
ALTER TABLE "api_keys" ADD FOREIGN KEY ("org_id") REFERENCES "organizations" ("id") DEFERRABLE INITIALLY IMMEDIATE;
ALTER TABLE "sessions" ADD FOREIGN KEY ("identity_id") REFERENCES "identities" ("id") DEFERRABLE INITIALLY IMMEDIATE;
ALTER TABLE "tokens" ADD FOREIGN KEY ("identity_id") REFERENCES "identities" ("id") DEFERRABLE INITIALLY IMMEDIATE;

CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER LANGUAGE plpgsql AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$;

CREATE TRIGGER trg_organizations_updated_at
    BEFORE UPDATE ON organizations
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER trg_identities_updated_at
    BEFORE UPDATE ON identities
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();
