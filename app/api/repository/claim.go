package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"api/model"
)

type ClaimRepository struct {
	DB  *sqlx.DB
	Ctx *context.Context
}

func (repository *ClaimRepository) Create(
	claim *model.Claim,
) error {
	const query = `
		INSERT INTO claim (content)
		VALUES ($1)
		RETURNING id, content, verdict, created_at
	`
	err := repository.DB.GetContext(
		*repository.Ctx,
		claim,
		query,
		claim.Content,
	)
	if err != nil {
		return err
	}

	return nil
}

func (repository *ClaimRepository) GetByID(ctx context.Context, id int64) (*model.Claim, error) {
	const query = `
		SELECT id, content, verdict, created_at
		FROM claim
		WHERE id = $1
	`

	var claim model.Claim

	if err := repository.DB.GetContext(*repository.Ctx, &claim, query, id); err != nil {
		return nil, err
	}

	return &claim, nil
}

func (repository *ClaimRepository) List(ctx context.Context) ([]model.Claim, error) {
	const query = `
		SELECT id, content, verdict, created_at
		FROM claim
		ORDER BY id DESC
	`

	var claims []model.Claim

	if err := repository.DB.SelectContext(*repository.Ctx, &claims, query); err != nil {
		return nil, err
	}

	return claims, nil
}

func (repository *ClaimRepository) Update(
	ctx context.Context,
	id int64,
	content string,
	verdict model.ClaimVerdict,
) (*model.Claim, error) {
	const query = `
		UPDATE claim
		SET content = $2,
		    verdict = $3
		WHERE id = $1
		RETURNING id, content, verdict, created_at
	`

	var claim model.Claim

	err := repository.DB.GetContext(
		*repository.Ctx,
		&claim,
		query,
		id,
		content,
		verdict,
	)
	if err != nil {
		return nil, err
	}

	return &claim, nil
}

func (repository *ClaimRepository) Delete(
	ctx context.Context,
	id int64,
) error {
	const query = `
		DELETE FROM claim
		WHERE id = $1
	`

	_, err := repository.DB.ExecContext(*repository.Ctx, query, id)
	return err
}
