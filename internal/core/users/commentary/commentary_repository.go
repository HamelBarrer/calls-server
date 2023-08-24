package commentary

import "github.com/HamelBarrer/calls-server/internal/storage"

type Repository interface {
	GetById(int) (*Commentary, error)
	GetAll(int) (*[]Commentary, error)
	Create(*Commentary) (*Commentary, error)
}

type Service struct {
	s storage.Storage
}

func NewService(s storage.Storage) Repository {
	return &Service{s}
}

func (s *Service) GetById(ci int) (*Commentary, error) {
	c := Commentary{}

	query := `
		select
			c.commentary_id,
			u.username,
			c.commentary,
			c.created_at
		from users.commentaries c
			join users.users u using (user_id)
		where c.commentary_id = $1;
	`

	err := s.s.QueryRow(query, ci).Scan(
		&c.CommentaryId,
		&c.User.Username,
		&c.Commentary,
		&c.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (s *Service) GetAll(ci int) (*[]Commentary, error) {
	cs := []Commentary{}

	query := `
		select
			c.commentary_id,
			u.username,
			c.commentary,
			c.created_at
		from users.commentaries c
			join users.users u using (user_id)
		order by c.commentary_id desc;
	`

	rows, err := s.s.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		c := Commentary{}

		err := rows.Scan(
			&c.CommentaryId,
			&c.User.Username,
			&c.Commentary,
			&c.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		cs = append(cs, c)
	}

	return &cs, nil
}

func (s *Service) Create(c *Commentary) (*Commentary, error) {
	query := `
		insert into users.commentaries (user_id, commentary)
		values ($1, $2)
		returning commentary_id;
	`

	registerId := 0

	if err := s.s.QueryRow(query, c.UserId, c.Commentary).Scan(&registerId); err != nil {
		return nil, err
	}

	return s.GetById(registerId)
}
