package user_usecase

import (
	"context"
	user_entity "exchange_rate/pkg/domain/user/entity"
	"exchange_rate/pkg/packages/errors"
	user_dto "exchange_rate/pkg/usecase/user/dto"
)

func (u *UserUC) CreateUser(
	ctx context.Context,
	data *user_dto.CreateUserDto,
) *errors.Error {
	return u.rep.CreateUserService(ctx, user_entity.NewUser(data.Mail))
}

func (u *UserUC) GetUserByEmail(
	ctx context.Context,
	data *user_dto.GetUserByEmailDto,
) (*user_entity.User, *errors.Error) {
	return u.rep.GetUserByEmailService(ctx, data.Mail.ToString())
}

func (u *UserUC) GetUsers(
	ctx context.Context,
) ([]*user_entity.User, *errors.Error) {
	return u.rep.GetUsersService(ctx)
}

func (u *UserUC) DeleteUserByEmail(
	ctx context.Context,
	data *user_dto.DeleteUserByEmailDto,
) *errors.Error {
	return u.rep.DeleteUserByEmailService(ctx, data.Mail.ToString())
}
