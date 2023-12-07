CREATE INDEX IF NOT EXISTS artifacts_name_idx ON artifacts USING GIN (to_tsvector('simple', name));
CREATE INDEX IF NOT EXISTS artifacts_origin_idx ON artifacts USING GIN (origin);
CREATE INDEX IF NOT EXISTS artifacts_type_idx ON artifacts USING GIN (type);