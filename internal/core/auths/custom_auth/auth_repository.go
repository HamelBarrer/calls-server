package customauth

import "github.com/HamelBarrer/calls-server/internal/storage"

type Repository interface {
	GetUserByUsername(string) (*AuthUser, error)
}

type Service struct {
	s storage.Storage
}

func NewService(s storage.Storage) Repository {
	return &Service{s}
}

func (s *Service) GetUserByUsername(u string) (*AuthUser, error) {
	a := AuthUser{}

	query := `
		select
			u.user_id,
			u.username,
			u.password,
			u.avatar
		from users.users u
		where u.username = $1;
	`

	if err := s.s.QueryRow(query, u).Scan(&a.UserId, &a.Username, &a.Password, &a.Avatar); err != nil {
		return nil, err
	}

	return &a, nil
}
