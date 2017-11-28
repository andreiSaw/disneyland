CREATE TABLE queues
(
  name   TEXT PRIMARY KEY,
  policy SMALLINT
);
ALTER TABLE jobs
  RENAME COLUMN kind TO queue;
INSERT INTO queues (name) SELECT DISTINCT queue
                          FROM jobs;
ALTER TABLE jobs
  ADD CONSTRAINT jobs_queues_name_fk
FOREIGN KEY (queue) REFERENCES queues (name);
INSERT INTO projects (name) VALUES ('docker'), ('test');
