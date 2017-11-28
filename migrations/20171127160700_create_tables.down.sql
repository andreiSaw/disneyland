ALTER TABLE jobs
  DROP CONSTRAINT IF EXISTS jobs_projects_name_fk;
DROP TABLE IF EXISTS projects;
