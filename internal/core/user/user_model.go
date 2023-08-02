package user

type User struct {
	UserId   int    `json:"user_id,omitempty" required:"false"`
	Username string `json:"username,omitempty" required:"true"`
}

type UserCreate struct {
	User
	Password        string `json:"password,omitempty" required:"true"`
	PasswordConfirm string `json:"password_confirm,omitempty" required:"true"`
}
