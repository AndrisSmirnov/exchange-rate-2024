package user_dto

import "exchange_rate/pkg/domain/vo"

type CreateUserDto struct {
	UserEmail `validate:"required"`
}

type GetUserByEmailDto struct {
	UserEmail `validate:"required"`
}

type DeleteUserByEmailDto struct {
	UserEmail `validate:"required"`
}

type UserEmail struct {
	Mail vo.Email `json:"id" validate:"required,email"`
}
