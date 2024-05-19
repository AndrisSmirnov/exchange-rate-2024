package user_usecase

import (
	"context"
	user_entity "exchange_rate/pkg/domain/user/entity"
	user_model "exchange_rate/pkg/domain/user/model"
	"exchange_rate/pkg/packages/errors"
)

type (
	Store interface {
		UserStore
		UserService
	}

	UserStore interface {
		CreateUser(ctx context.Context, data *user_model.UserDB) *errors.Error
		GetUserByID(ctx context.Context, id string) (*user_model.UserDB, *errors.Error)
		GetUserByEmail(ctx context.Context, email string) (*user_model.UserDB, *errors.Error)
		GetUsers(ctx context.Context) ([]*user_model.UserDB, *errors.Error)
		DeleteUserByID(ctx context.Context, id string) *errors.Error
		DeleteUserByEmail(ctx context.Context, email string) *errors.Error
	}

	UserService interface {
		CreateUserService(ctx context.Context, user *user_entity.User) *errors.Error
		GetUserByIdService(ctx context.Context, id string) (*user_entity.User, *errors.Error)
		GetUserByEmailService(ctx context.Context, email string) (*user_entity.User, *errors.Error)
		GetUsersService(ctx context.Context) ([]*user_entity.User, *errors.Error)
		DeleteUserByIdService(ctx context.Context, id string) *errors.Error
		DeleteUserByEmailService(ctx context.Context, email string) *errors.Error
	}
)
