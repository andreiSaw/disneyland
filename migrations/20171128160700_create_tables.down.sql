ALTER TABLE jobs
  DROP CONSTRAINT IF EXISTS jobs_queues_name_fk;
DROP TABLE IF EXISTS queues;
ALTER TABLE jobs
  RENAME COLUMN queue TO kind;
