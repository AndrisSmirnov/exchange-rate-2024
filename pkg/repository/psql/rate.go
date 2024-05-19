package pstgrs

import (
	"context"
	rate_model "exchange_rate/pkg/domain/rate/model"
	"exchange_rate/pkg/packages/errors"
)

func (r *Repository) CreateRate(
	ctx context.Context, data *rate_model.RateDB,
) *errors.Error {
	query := `
		INSERT INTO "rate"
			(
				"id", "date", "rate", "val_code", "code"
			)
			VALUES (
				$1, $2, $3, $4, $5
			)
		`

	_, err := r.DB.ExecContext(
		ctx,
		query,
		data.GetID(),
		data.GetDate(),
		data.GetRate(),
		data.GetValCode(),
		data.GetCode(),
	)

	if err != nil {
		return r.ErrorHandler(err, Rate)
	}

	return nil
}

func (r *Repository) GetRateByID(
	ctx context.Context, id string,
) (*rate_model.RateDB, *errors.Error) {
	query := `
		SELECT
			id, date, rate, val_code, code
		FROM rate
		WHERE id = $1;
		`
	var rateDB = rate_model.RateDB{}

	err := r.DB.GetContext(
		ctx,
		&rateDB,
		query,
		id,
	)

	if err != nil {
		return nil, r.ErrorHandler(err, Rate)
	}

	return &rateDB, nil
}

func (r *Repository) GetRateByValCode(
	ctx context.Context, valCode string,
) (*rate_model.RateDB, *errors.Error) {
	query := `
		SELECT
			id, date, rate, val_code, code
		FROM rate
		WHERE id = $1;
		`
	var rateDB = rate_model.RateDB{}

	err := r.DB.GetContext(
		ctx,
		&rateDB,
		query,
		valCode,
	)

	if err != nil {
		return nil, r.ErrorHandler(err, Rate)
	}

	return &rateDB, nil
}

func (r *Repository) GetRates(
	ctx context.Context,
) ([]*rate_model.RateDB, *errors.Error) {
	query := `
		SELECT
			id, date, rate, val_code, code
		FROM rate;
		`
	dest := []*rate_model.RateDB{}

	err := r.DB.SelectContext(
		ctx,
		&dest,
		query,
	)

	if err != nil {
		return nil, r.ErrorHandler(err, Rate)
	}

	return dest, nil
}

func (r *Repository) UpdateRate(ctx context.Context, data *rate_model.RateDB) *errors.Error {
	query := `
	UPDATE "rate"
	SET
		"date" = $2,
		"rate" = $3,
	WHERE "id" = $1
	`
	if _, err := r.DB.ExecContext(
		ctx, query,
		data.GetID(),
		data.GetDate(),
		data.GetRate(),
	); err != nil {
		return r.ErrorHandler(err, Rate)
	}
	return nil
}

func (r *Repository) DeleteRateByID(ctx context.Context, id string) *errors.Error {
	query := `
	DELETE
	FROM "rate"
	WHERE "id" = $1
	`

	if _, err := r.DB.ExecContext(ctx, query, id); err != nil {
		return r.ErrorHandler(err, Rate)
	}

	return nil
}

func (r *Repository) DeleteRateByValCode(ctx context.Context, vadCode string) *errors.Error {
	query := `
	DELETE
	FROM "rate"
	WHERE "vad_code" = $1
	`

	if _, err := r.DB.ExecContext(ctx, query, vadCode); err != nil {
		return r.ErrorHandler(err, Rate)
	}

	return nil
}
