package followeduser

import "github.com/HamelBarrer/calls-server/internal/core/users/user"

type FollowedUser struct {
	FollowedUserId         int       `json:"followed_user_id,omitempty" required:"false"`
	User                   user.User `json:"user,omitempty" required:"false"`
	UserId                 int       `json:"user_id,omitempty" required:"true"`
	FollowerUser           user.User `json:"follower_user,omitempty" required:"false"`
	FollowerUserId         int       `json:"follower_user_id,omitempty" required:"true"`
	CanceledFollowerUser   user.User `json:"canceled_follower_user,omitempty" required:"false"`
	CanceledFollowerUserId int       `json:"canceled_follower_user_id,omitempty" required:"false"`
	ReactiveFollowerUser   user.User `json:"reactive_followeruser,omitempty" required:"false"`
	ReactiveFollowerUserId int       `json:"reactive_follower_user_id,omitempty" required:"false"`
}
