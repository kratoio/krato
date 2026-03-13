DROP TRIGGER IF EXISTS trg_identities_updated_at ON identities;
DROP TRIGGER IF EXISTS trg_organizations_updated_at ON organizations;

DROP FUNCTION IF EXISTS set_updated_at();

DROP TABLE IF EXISTS tokens CASCADE;
DROP TABLE IF EXISTS sessions CASCADE;
DROP TABLE IF EXISTS api_keys CASCADE;
DROP TABLE IF EXISTS identities CASCADE;
DROP TABLE IF EXISTS organizations CASCADE;

DROP EXTENSION IF EXISTS "citext";
