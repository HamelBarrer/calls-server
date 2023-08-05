package followeduser

import "github.com/HamelBarrer/calls-server/internal/storage"

type Repository interface {
	GetById(int) (*FollowedUser, error)
	GetAllFollowedUser() (*[]FollowedUser, error)
	Create(FollowedUser) (*FollowedUser, error)
}

type Service struct {
	s storage.Storage
}

func NewService(s storage.Storage) Repository {
	return &Service{s}
}

func (s *Service) GetById(fi int) (*FollowedUser, error) {
	f := FollowedUser{}

	query := `
		select
			f.followed_user_id,
			f.user_id,
			u.username,
			uf.user_id,
			uf.username
		from users.followed_users f
			join users.users u using(user_id)
			join users.users uf on uf.user_id = f.followed_user_id
		where f.user_id = $1;
	`

	err := s.s.QueryRow(query, fi).Scan(
		&f.FollowedUserId,
		&f.User.UserId,
		&f.User.Username,
		&f.FollowerUser.UserId,
		&f.FollowerUser.Username,
	)
	if err != nil {
		return nil, err
	}

	return &f, nil
}

func (s *Service) GetAllFollowedUser() (*[]FollowedUser, error) {
	fs := []FollowedUser{}

	query := `
		select
			f.followed_user_id,
			f.user_id,
			u.username,
			uf.user_id,
			uf.username
		from users.followed_users f
			join users.users u using(user_id)
			join users.users uf on uf.user_id = f.followed_user_id;
	`

	rows, err := s.s.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		f := FollowedUser{}

		err := rows.Scan(
			&f.FollowedUserId,
			&f.User.UserId,
			&f.User.Username,
			&f.FollowerUser.UserId,
			&f.FollowerUser.Username,
		)
		if err != nil {
			return nil, err
		}

		fs = append(fs, f)
	}

	return &fs, nil
}

func (s *Service) Create(f FollowedUser) (*FollowedUser, error) {
	query := `
		insert into users.followed_users (user_id, follower_user_id)
		values ($1, $2)
		returning user_id;
	`

	registerId := 0

	if err := s.s.QueryRow(query, f.UserId, f.FollowerUserId).Scan(&registerId); err != nil {
		return nil, err
	}

	return s.GetById(registerId)
}
