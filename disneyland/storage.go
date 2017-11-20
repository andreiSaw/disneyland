package disneyland

import (
	"database/sql"
	_ "github.com/lib/pq"
	"time"
)

type DisneylandStorageConfig struct {
	DatabaseURI string `json:"db_uri"`
}

type DisneylandStorage struct {
	db     *sql.DB
	Config DisneylandStorageConfig
}

func NewDisneylandStorage(dbUri string) (*DisneylandStorage, error) {
	ret := &DisneylandStorage{
		Config: DisneylandStorageConfig{DatabaseURI: dbUri},
	}

	err := ret.Connect()

	return ret, err
}

func (storage *DisneylandStorage) Connect() error {
	db, err := sql.Open("postgres", storage.Config.DatabaseURI)
	storage.db = db
	return err
}

func (storage *DisneylandStorage) CreateJob(job *Job, creator User) (*Job, error) {
	tx, err := storage.db.Begin()
	if err != nil {
		return nil, err
	}

	createdJob := &Job{}
	err = tx.QueryRow(`
		INSERT INTO jobs (project, status, metadata, creator, input, output, kind)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, project, status, metadata, input, output, kind;`,
		job.Project, job.Status, job.Metadata, creator.Username, job.Input, job.Output, job.Kind,
	).Scan(
		&createdJob.Id,
		&createdJob.Project,
		&createdJob.Status,
		&createdJob.Metadata,
		&createdJob.Input,
		&createdJob.Output,
		&createdJob.Kind,
	)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}
	return createdJob, err
}

func (storage *DisneylandStorage) GetJob(id uint64) (*Job, error) {
	tx, err := storage.db.Begin()
	if err != nil {
		return nil, err
	}
	job := &Job{}

	strQuery := `SELECT id, project, status,  metadata, input, output, kind
					FROM jobs
					WHERE id=$1;`
	err = tx.QueryRow(strQuery, id).Scan(
		&job.Id,
		&job.Project,
		&job.Status,
		&job.Metadata,
		&job.Input,
		&job.Output,
		&job.Kind,
	)

	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}
	return job, err
}

func queryJobs(rows *sql.Rows) (*ListOfJobs, error) {
	ret := &ListOfJobs{Jobs: []*Job{}}
	var err error

	for rows.Next() {
		job := &Job{}

		err := rows.Scan(
			&job.Id,
			&job.Project,
			&job.Status,
			&job.Metadata,
			&job.Input,
			&job.Output,
			&job.Kind,
		)

		if err != nil {
			return nil, err
		}
		ret.Jobs = append(ret.Jobs, job)
	}
	return ret, err
}

func (storage *DisneylandStorage) ListJobs(howmany uint32, project string, kind string) (*ListOfJobs, error) {
	strQuery := `SELECT id, project, status, metadata, input, output, kind
				 FROM jobs
				 WHERE project LIKE '%' || $1 || '%' AND kind LIKE '%' || $2 || '%'
				 LIMIT $3;`

	rows, err := storage.db.Query(strQuery, project, kind, howmany)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ret, err := queryJobs(rows)
	if err != nil {
		return nil, err
	}
	err = rows.Err()
	return ret, err
}

func (storage *DisneylandStorage) UpdateJob(job *Job) (*Job, error) {
	tx, err := storage.db.Begin()
	if err != nil {
		return nil, err
	}
	curTime:=time.Now().UTC()

	resultJob := &Job{}
	err = tx.QueryRow(`
		UPDATE jobs
		SET
			status=$1,
			metadata=$2,
			output=$3,
			kind=$4,
			last_modified=$5
		WHERE id=$6
		RETURNING id, project, status, metadata, input, output, kind;`,
		job.Status,
		job.Metadata,
		job.Output,
		job.Kind,
		curTime,
		job.Id,
	).Scan(
		&resultJob.Id,
		&resultJob.Project,
		&resultJob.Status,
		&resultJob.Metadata,
		&resultJob.Input,
		&resultJob.Output,
		&resultJob.Kind,
	)
	if err != nil {
		return resultJob, err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return resultJob, err
	}

	return resultJob, err
}

func (storage *DisneylandStorage) PullJobs(how_many uint32, project string, kind string) (*ListOfJobs, error) {
	tx, err := storage.db.Begin()
	if err != nil {
		return nil, err
	}
	curTime:=time.Now().UTC()
	strQuery := `WITH updatedPts AS (
					WITH pulledPts AS (
						SELECT id, project, kind
						FROM jobs
						WHERE status=$1 AND project LIKE '%' || $2 || '%' AND kind LIKE '%' || $3 || '%'
						LIMIT $4
						FOR UPDATE SKIP LOCKED)
					UPDATE jobs pts
					SET status=$5, last_modified=$6
					FROM pulledPts
					WHERE pulledPts.id=pts.id AND pulledPts.project=pts.project AND pulledPts.kind=pts.kind
					RETURNING pts.id, pts.project, pts.status, pts.metadata, pts.input, pts.output, pts.kind)
				SELECT *
				FROM updatedPts
				ORDER BY id ASC;`

	rows, err := tx.Query(strQuery, Job_PENDING, project, kind, how_many, Job_PULLED, curTime)

	if err != nil {
		return nil, err
	}
	ret, err := queryJobs(rows)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return ret, err
}

func (storage *DisneylandStorage) DeleteJob(id uint64, userProject string) (*Job, error) {
	tx, err := storage.db.Begin()
	if err != nil {
		return nil, err
	}
	resultJob := &Job{}
	err = tx.QueryRow(`
		DELETE FROM jobs
		WHERE id=$1 AND project=$2
		RETURNING id, project, kind;`, id, userProject,
	).Scan(
		&resultJob.Id,
		&resultJob.Project,
		&resultJob.Kind,
	)
	if err != nil {
		return resultJob, err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}
	return resultJob, err
}
