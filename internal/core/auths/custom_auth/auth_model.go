package customauth

type Auth struct {
	Username string `json:"username,omitempty" required:"true"`
	Password string `json:"password,omitempty" required:"true"`
}

type AuthUser struct {
	UserId   int    `json:"user_id,omitempty" required:"false"`
	Username string `json:"username,omitempty" required:"true"`
	Password string `json:"password,omitempty" required:"true"`
}
