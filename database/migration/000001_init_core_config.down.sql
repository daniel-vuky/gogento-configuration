DROP TABLE IF EXISTS core_config_data;

DROP TRIGGER IF EXISTS "update_updated_at" ON core_config_data;

DROP FUNCTION IF EXISTS update_updated_at_column;