ALTER TABLE jobs ADD last_modified TIMESTAMP WITHOUT TIME ZONE DEFAULT (now() AT TIME ZONE 'utc'); ALTER TABLE jobs ALTER COLUMN status SET DEFAULT 0;
