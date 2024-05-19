package user_service

import (
	"context"
	user_entity "exchange_rate/pkg/domain/user/entity"
	user_mapper "exchange_rate/pkg/domain/user/mapper"

	"exchange_rate/pkg/packages/errors"
)

func (s *UserService) CreateUserService(ctx context.Context, user *user_entity.User) *errors.Error {
	return s.repository.CreateUser(ctx, user_mapper.ConvertUserToDB(user))
}

func (s *UserService) GetUserByIdService(ctx context.Context, id string) (*user_entity.User, *errors.Error) {
	userDB, err := s.repository.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user_mapper.ConvertUserFromDB(userDB), nil
}

func (s *UserService) GetUserByEmailService(ctx context.Context, email string) (*user_entity.User, *errors.Error) {
	userDB, err := s.repository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user_mapper.ConvertUserFromDB(userDB), nil
}

func (s *UserService) GetUsersService(ctx context.Context) ([]*user_entity.User, *errors.Error) {
	usersDB, err := s.repository.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	users := make([]*user_entity.User, 0, len(usersDB))

	for _, el := range usersDB {
		users = append(users, user_mapper.ConvertUserFromDB(el))
	}

	return users, nil
}

func (s *UserService) DeleteUserByIdService(ctx context.Context, id string) *errors.Error {
	return s.repository.DeleteUserByID(ctx, id)
}

func (s *UserService) DeleteUserByEmailService(ctx context.Context, email string) *errors.Error {
	return s.repository.DeleteUserByEmail(ctx, email)
}
