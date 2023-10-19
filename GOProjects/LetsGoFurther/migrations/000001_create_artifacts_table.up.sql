CREATE TABLE IF NOT EXISTS artifacts (
id bigserial PRIMARY KEY,
created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
name text NOT NULL,
origin text NOT NULL,
year integer NOT NULL,
type text NOT NULL,
version integer NOT NULL DEFAULT 1
);
