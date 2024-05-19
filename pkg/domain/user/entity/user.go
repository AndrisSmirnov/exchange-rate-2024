package user_entity

import "exchange_rate/pkg/domain/vo"

type User struct {
	ID   vo.UUID  `json:"id"`
	Mail vo.Email `json:"mail"`
}

func NewUser(email vo.Email) *User {
	return &User{
		ID:   vo.NewID(),
		Mail: email,
	}
}
