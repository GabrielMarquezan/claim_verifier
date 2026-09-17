package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"api/model"
)

type ClaimChunkRepository struct {
	DB *sqlx.DB
}

func (repository *ClaimChunkRepository) Add(
	ctx context.Context,
	claimID int64,
	chunkID int64,
) error {
	const query = `
		INSERT INTO claim_chunk (claim_id, chunk_id)
		VALUES ($1, $2)
		ON CONFLICT DO NOTHING
	`

	_, err := repository.DB.ExecContext(
		ctx,
		query,
		claimID,
		chunkID,
	)

	return err
}

func (repository *ClaimChunkRepository) Remove(
	ctx context.Context,
	claimID int64,
	chunkID int64,
) error {
	const query = `
		DELETE FROM claim_chunk
		WHERE claim_id = $1
		  AND chunk_id = $2
	`

	_, err := repository.DB.ExecContext(
		ctx,
		query,
		claimID,
		chunkID,
	)

	return err
}

func (repository *ClaimChunkRepository) GetChunks(
	ctx context.Context,
	claimID int64,
) ([]model.Chunk, error) {
	const query = `
		SELECT
			c.id,
			c.content,
			c.source_id,
			c.created_at,
			c.deleted_at
		FROM chunk c
		INNER JOIN claim_chunk cc
			ON cc.chunk_id = c.id
		WHERE cc.claim_id = $1
		  AND c.deleted_at IS NULL
		ORDER BY c.id
	`

	var chunks []model.Chunk

	if err := repository.DB.SelectContext(
		ctx,
		&chunks,
		query,
		claimID,
	); err != nil {
		return nil, err
	}

	return chunks, nil
}
