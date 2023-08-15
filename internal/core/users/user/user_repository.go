package user

import (
	"github.com/HamelBarrer/calls-server/internal/storage"
	"github.com/HamelBarrer/calls-server/internal/utils"
)

type Repository interface {
	GetById(int) (*User, error)
	GetAllUser() (*[]User, error)
	Create(UserCreate) (*User, error)
	Update(int, string) (*User, error)
}

type repository struct {
	s storage.Storage
}

func Newrepository(s storage.Storage) Repository {
	return &repository{s}
}

func (r *repository) GetById(ui int) (*User, error) {
	u := User{}

	query := `
		select
			u.user_id,
			u.username,
			u.avatar
		from users.users u
		where u.user_id = $1;
	`

	if err := r.s.QueryRow(query, ui).Scan(&u.UserId, &u.Username, &u.Avatar); err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *repository) GetAllUser() (*[]User, error) {
	us := []User{}

	query := `
		select
			u.user_id,
			u.username,
			u.avatar
		from users.users u;
	`

	rows, err := r.s.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		u := User{}

		if err := rows.Scan(&u.UserId, &u.Username, &u.Avatar); err != nil {
			return nil, err
		}

		us = append(us, u)
	}

	return &us, nil
}

func (r *repository) Create(uc UserCreate) (*User, error) {
	p, err := utils.CreationHash(uc.Password)
	if err != nil {
		return nil, err
	}

	query := `
		insert into users.users (username, password)
		values ($1, $2)
		returning user_id;
	`

	registerId := 0

	if err := r.s.QueryRow(query, uc.Username, p).Scan(&registerId); err != nil {
		return nil, err
	}

	return r.GetById(registerId)
}

func (r *repository) Update(ui int, a string) (*User, error) {
	query := `
		update users.users
		set avatar = $1,
			updated_at = now()
		where user_id = $2
		returning user_id;
	`

	updatedId := 0

	err := r.s.QueryRow(query, a, ui).Scan(&updatedId)
	if err != nil {
		return nil, err
	}

	return r.GetById(updatedId)
}
