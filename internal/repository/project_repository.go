package repository

import (
	"github.com/anisetusbambangmanalu19/task-management/internal/config"
	"github.com/anisetusbambangmanalu19/task-management/internal/entity"
)

type ProjectRepository struct{}

func (r *ProjectRepository) Create(project *entity.Project) error {
	query := `
		INSERT INTO projects (user_id, name, description)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	return config.DB.QueryRow(
		query,
		project.UserID,
		project.Name,
		project.Description,
	).Scan(&project.ID)
}

func (r *ProjectRepository) GetByUser(userID int) ([]entity.Project, error) {

	query := `
		SELECT id, user_id, name, description, created_at
		FROM projects
		WHERE user_id = $1
	`

	rows, err := config.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []entity.Project

	for rows.Next() {
		var p entity.Project
		err := rows.Scan(&p.ID, &p.UserID, &p.Name, &p.Description, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	return projects, nil
}

func (r *ProjectRepository) IsOwner(projectID int, userID int) (bool, error) {

	query := `
		SELECT COUNT(*) 
		FROM projects 
		WHERE id = $1 AND user_id = $2
	`

	var count int
	err := config.DB.QueryRow(query, projectID, userID).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

