package repository

import (
	"github.com/anisetusbambangmanalu19/task-management/internal/config"
	"github.com/anisetusbambangmanalu19/task-management/internal/entity"
)

type UserRepository struct{}

func (r *UserRepository) Create(user *entity.User) error {
	query := `
		INSERT INTO users (name, email, password, role)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	return config.DB.QueryRow(
		query,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
	).Scan(&user.ID)
}

func (r *UserRepository) FindByEmail(email string) (*entity.User, error) {
	query := `
		SELECT id, name, email, password, role, created_at
		FROM users
		WHERE email = $1
	`

	user := &entity.User{}
	err := config.DB.QueryRow(query, email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}
