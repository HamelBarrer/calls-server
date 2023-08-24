package commentary

import (
	"time"

	"github.com/HamelBarrer/calls-server/internal/core/users/user"
)

type Commentary struct {
	CommentaryId int       `json:"commentary_id,omitempty" required:"false"`
	User         user.User `json:"user,omitempty" required:"false"`
	UserId       int       `json:"user_id,omitempty" required:"true"`
	Commentary   string    `json:"commentary,omitempty" required:"true"`
	CreatedAt    time.Time `json:"created_at,omitempty" required:"false"`
}
