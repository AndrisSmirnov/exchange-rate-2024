package user_service

import (
	"context"
	user_model "exchange_rate/pkg/domain/user/model"
	"exchange_rate/pkg/packages/errors"
)

type (
	Repository interface {
		CreateUser(ctx context.Context, data *user_model.UserDB) *errors.Error
		GetUserByID(ctx context.Context, id string) (*user_model.UserDB, *errors.Error)
		GetUserByEmail(ctx context.Context, email string) (*user_model.UserDB, *errors.Error)
		GetUsers(ctx context.Context) ([]*user_model.UserDB, *errors.Error)
		DeleteUserByID(ctx context.Context, id string) *errors.Error
		DeleteUserByEmail(ctx context.Context, email string) *errors.Error
	}
)
