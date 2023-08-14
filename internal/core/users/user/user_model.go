package user

import (
	"github.com/jackc/pgx/v5/pgtype/zeronull"
)

type User struct {
	UserId   int           `json:"user_id,omitempty" required:"false"`
	Username string        `json:"username,omitempty" required:"true"`
	Avatar   zeronull.Text `json:"avatar,omitempty" required:"false"`
}

type UserCreate struct {
	User
	Password        string `json:"password,omitempty" required:"true"`
	PasswordConfirm string `json:"password_confirm,omitempty" required:"true"`
}
