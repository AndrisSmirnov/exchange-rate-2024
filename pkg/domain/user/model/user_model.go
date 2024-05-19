package user_model

type UserDB struct {
	ID   string `db:"id"`
	Mail string `db:"mail"`
}

func (u *UserDB) GetUserID() string { return u.ID }
func (u *UserDB) GetMail() string   { return u.Mail }
