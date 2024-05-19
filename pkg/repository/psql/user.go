package pstgrs

import (
	"context"
	user_model "exchange_rate/pkg/domain/user/model"
	"exchange_rate/pkg/packages/errors"
	"fmt"
)

func (r *Repository) CreateUser(
	ctx context.Context, data *user_model.UserDB,
) *errors.Error {
	fmt.Printf("--> CREATE USER %+v\n", data)
	query := `
		INSERT INTO "user"
			(
				"id", "mail"
			)
			VALUES (
				$1, $2
			)
		`

	_, err := r.DB.ExecContext(
		ctx,
		query,
		data.GetUserID(),
		data.GetMail(),
	)

	if err != nil {
		return r.ErrorHandler(err, User)
	}

	return nil
}

func (r *Repository) GetUserByID(
	ctx context.Context, id string,
) (*user_model.UserDB, *errors.Error) {
	query := `
		SELECT
			id, mail
		FROM "user"
		WHERE id = $1;
		`
	var userDB = user_model.UserDB{}

	err := r.DB.GetContext(
		ctx,
		&userDB,
		query,
		id,
	)

	if err != nil {
		return nil, r.ErrorHandler(err, User)
	}

	return &userDB, nil
}

func (r *Repository) GetUserByEmail(
	ctx context.Context, mail string,
) (*user_model.UserDB, *errors.Error) {
	query := `
		SELECT
			id, mail
		FROM "user"
		WHERE mail = $1;
		`
	var userDB = user_model.UserDB{}

	err := r.DB.GetContext(
		ctx,
		&userDB,
		query,
		mail,
	)

	if err != nil {
		return nil, r.ErrorHandler(err, User)
	}

	return &userDB, nil
}

func (r *Repository) GetUsers(
	ctx context.Context,
) ([]*user_model.UserDB, *errors.Error) {
	query := `
		SELECT
			id, mail
		FROM "user";
		`
	dest := []*user_model.UserDB{}

	err := r.DB.SelectContext(
		ctx,
		&dest,
		query,
	)

	if err != nil {
		return nil, r.ErrorHandler(err, User)
	}

	return dest, nil
}

func (r *Repository) DeleteUserByID(ctx context.Context, id string) *errors.Error {
	query := `
	DELETE
	FROM "user"
	WHERE "id" = $1
	`

	if _, err := r.DB.ExecContext(ctx, query, id); err != nil {
		return r.ErrorHandler(err, User)
	}

	return nil
}

func (r *Repository) DeleteUserByEmail(ctx context.Context, mail string) *errors.Error {
	query := `
	DELETE
	FROM "user"
	WHERE "mail" = $1
	`

	if _, err := r.DB.ExecContext(ctx, query, mail); err != nil {
		return r.ErrorHandler(err, User)
	}

	return nil
}
