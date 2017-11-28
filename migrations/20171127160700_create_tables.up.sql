CREATE TABLE projects
(
  name            TEXT PRIMARY KEY,
  characteristics INT
);
INSERT INTO projects (name) SELECT DISTINCT project
                            FROM jobs;
ALTER TABLE public.jobs
  ADD CONSTRAINT jobs_projects_name_fk
FOREIGN KEY (project) REFERENCES projects (name);
INSERT INTO projects (name) VALUES ('ship-shield'), ('test-project');
