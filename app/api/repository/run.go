package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"api/model"
)

type RunRepository struct {
	DB *sqlx.DB
}

func (repository *RunRepository) Create(ctx context.Context, claimID int64, modelID string) (*model.Run, error) {
	const query = `
		INSERT INTO run (claim_id, model_id)
		VALUES ($1, $2)
		RETURNING
			id,
			claim_id,
			status,
			model_id,
			created_at,
			finished_at
	`

	var run model.Run

	err := repository.DB.GetContext(
		ctx,
		&run,
		query,
		claimID,
		modelID,
	)
	if err != nil {
		return nil, err
	}

	return &run, nil
}

func (repository *RunRepository) GetByID(
	ctx context.Context,
	id int64,
) (*model.Run, error) {
	const query = `
		SELECT
			id,
			claim_id,
			status,
			model_id,
			created_at,
			finished_at
		FROM run
		WHERE id = $1
	`

	var run model.Run

	if err := repository.DB.GetContext(ctx, &run, query, id); err != nil {
		return nil, err
	}

	return &run, nil
}

func (repository *RunRepository) List(
	ctx context.Context,
) ([]model.Run, error) {
	const query = `
		SELECT
			id,
			claim_id,
			status,
			model_id,
			created_at,
			finished_at
		FROM run
		ORDER BY id DESC
	`

	var runs []model.Run

	if err := repository.DB.SelectContext(ctx, &runs, query); err != nil {
		return nil, err
	}

	return runs, nil
}

func (repository *RunRepository) UpdateStatus(
	ctx context.Context,
	id int64,
	status model.RunStatus,
) (*model.Run, error) {
	const query = `
		UPDATE run
		SET
			status = $2,
			finished_at = CASE
				WHEN $2 IN ('SUCCESSFUL', 'FAILED')
					THEN NOW()
				ELSE NULL
			END
		WHERE id = $1
		RETURNING
			id,
			claim_id,
			status,
			model_id,
			created_at,
			finished_at
	`

	var run model.Run

	if err := repository.DB.GetContext(
		ctx,
		&run,
		query,
		id,
		status,
	); err != nil {
		return nil, err
	}

	return &run, nil
}

func (repository *RunRepository) Delete(
	ctx context.Context,
	id int64,
) error {
	const query = `
		DELETE FROM run
		WHERE id = $1
	`

	_, err := repository.DB.ExecContext(ctx, query, id)
	return err
}
