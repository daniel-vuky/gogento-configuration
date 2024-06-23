CREATE TABLE core_config_data (
    "config_id" bigSerial PRIMARY KEY,
    "path" varchar(255) UNIQUE NOT NULL,
    "value" text,
    "created_at" timestamptz NOT NULL DEFAULT 'NOW()',
    "updated_at" timestamptz NOT NULL DEFAULT 'NOW()'
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_updated_at
    BEFORE UPDATE ON core_config_data
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

ALTER TABLE core_config_data ADD CONSTRAINT CORE_CONFIG_DATA_PATH_NOT_EMPTY CHECK (path <> '');
