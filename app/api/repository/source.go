package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	"api/model"
)

type SourceRepository struct {
	DB *sqlx.DB
}

func (repository *SourceRepository) Create(
	ctx context.Context,
	url string,
	title *string,
	publishedAt *time.Time,
) (*model.Source, error) {
	const query = `
		INSERT INTO source (url, title, published_at)
		VALUES ($1, $2, $3)
		RETURNING
			id,
			url,
			title,
			created_at,
			published_at,
			deleted_at
	`

	var source model.Source

	err := repository.DB.GetContext(
		ctx,
		&source,
		query,
		url,
		title,
		publishedAt,
	)
	if err != nil {
		return nil, err
	}

	return &source, nil
}

func (repository *SourceRepository) GetByID(
	ctx context.Context,
	id int64,
) (*model.Source, error) {
	const query = `
		SELECT
			id,
			url,
			title,
			created_at,
			published_at,
			deleted_at
		FROM source
		WHERE id = $1
		  AND deleted_at IS NULL
	`

	var source model.Source

	if err := repository.DB.GetContext(ctx, &source, query, id); err != nil {
		return nil, err
	}

	return &source, nil
}

func (repository *SourceRepository) List(
	ctx context.Context,
) ([]model.Source, error) {
	const query = `
		SELECT
			id,
			url,
			title,
			created_at,
			published_at,
			deleted_at
		FROM source
		WHERE deleted_at IS NULL
		ORDER BY id DESC
	`

	var sources []model.Source

	if err := repository.DB.SelectContext(ctx, &sources, query); err != nil {
		return nil, err
	}

	return sources, nil
}

func (repository *SourceRepository) Update(
	ctx context.Context,
	id int64,
	url string,
	title *string,
	publishedAt *time.Time,
) (*model.Source, error) {
	const query = `
		UPDATE source
		SET url = $2,
		    title = $3,
		    published_at = $4
		WHERE id = $1
		  AND deleted_at IS NULL
		RETURNING
			id,
			url,
			title,
			created_at,
			published_at,
			deleted_at
	`

	var source model.Source

	err := repository.DB.GetContext(
		ctx,
		&source,
		query,
		id,
		url,
		title,
		publishedAt,
	)
	if err != nil {
		return nil, err
	}

	return &source, nil
}

func (repository *SourceRepository) Delete(
	ctx context.Context,
	id int64,
) error {
	const query = `
		UPDATE source
		SET deleted_at = NOW()
		WHERE id = $1
		  AND deleted_at IS NULL
	`

	_, err := repository.DB.ExecContext(ctx, query, id)
	return err
}
