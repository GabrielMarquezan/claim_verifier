package repository

import (
	"context"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"

	"api/model"
)

type ChunkRepository struct {
	DB *sqlx.DB
}

func vectorToString(vector []float32) string {
	values := make([]string, len(vector))

	for i, value := range vector {
		values[i] = strconv.FormatFloat(
			float64(value),
			'f',
			-1,
			32,
		)
	}

	return "[" + strings.Join(values, ",") + "]"
}

func (repository *ChunkRepository) Create(
	ctx context.Context,
	content string,
	embedding []float32,
	sourceID int64,
) (*model.Chunk, error) {
	const query = `
		INSERT INTO chunk (
			content,
			embedding,
			source_id
		)
		VALUES ($1, $2::vector, $3)
		RETURNING
			id,
			content,
			source_id,
			created_at,
			deleted_at
	`

	var chunk model.Chunk

	err := repository.DB.GetContext(
		ctx,
		&chunk,
		query,
		content,
		vectorToString(embedding),
		sourceID,
	)
	if err != nil {
		return nil, err
	}

	return &chunk, nil
}

func (repository *ChunkRepository) GetByID(
	ctx context.Context,
	id int64,
) (*model.Chunk, error) {
	const query = `
		SELECT
			id,
			content,
			source_id,
			created_at,
			deleted_at
		FROM chunk
		WHERE id = $1
		  AND deleted_at IS NULL
	`

	var chunk model.Chunk

	if err := repository.DB.GetContext(ctx, &chunk, query, id); err != nil {
		return nil, err
	}

	return &chunk, nil
}

func (repository *ChunkRepository) List(
	ctx context.Context,
) ([]model.Chunk, error) {
	const query = `
		SELECT
			id,
			content,
			source_id,
			created_at,
			deleted_at
		FROM chunk
		WHERE deleted_at IS NULL
		ORDER BY id DESC
	`

	var chunks []model.Chunk

	if err := repository.DB.SelectContext(ctx, &chunks, query); err != nil {
		return nil, err
	}

	return chunks, nil
}

func (repository *ChunkRepository) Update(
	ctx context.Context,
	id int64,
	content string,
	embedding []float32,
) (*model.Chunk, error) {
	const query = `
		UPDATE chunk
		SET content = $2,
		    embedding = $3::vector
		WHERE id = $1
		  AND deleted_at IS NULL
		RETURNING
			id,
			content,
			source_id,
			created_at,
			deleted_at
	`

	var chunk model.Chunk

	err := repository.DB.GetContext(
		ctx,
		&chunk,
		query,
		id,
		content,
		vectorToString(embedding),
	)
	if err != nil {
		return nil, err
	}

	return &chunk, nil
}

func (repository *ChunkRepository) Delete(
	ctx context.Context,
	id int64,
) error {
	const query = `
		UPDATE chunk
		SET deleted_at = NOW()
		WHERE id = $1
		  AND deleted_at IS NULL
	`

	_, err := repository.DB.ExecContext(ctx, query, id)
	return err
}
