package user_mapper

import (
	user_entity "exchange_rate/pkg/domain/user/entity"
	user_model "exchange_rate/pkg/domain/user/model"
	"exchange_rate/pkg/domain/vo"
)

func ConvertUserToDB(user *user_entity.User) *user_model.UserDB {
	return &user_model.UserDB{
		ID:   user.ID.String(),
		Mail: user.Mail.ToString(),
	}
}

func ConvertUserFromDB(userDB *user_model.UserDB) *user_entity.User {
	return &user_entity.User{
		ID:   vo.UUID(userDB.ID),
		Mail: vo.Email(userDB.Mail),
	}
}
