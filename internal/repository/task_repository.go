package repository

import (
	"github.com/anisetusbambangmanalu19/task-management/internal/config"
	"github.com/anisetusbambangmanalu19/task-management/internal/entity"
)

type TaskRepository struct{}

func (r *TaskRepository) Create(task *entity.Task) error {

	query := `
		INSERT INTO tasks (project_id, title, description, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	return config.DB.QueryRow(
		query,
		task.ProjectID,
		task.Title,
		task.Description,
		task.Status,
	).Scan(&task.ID)
}

func (r *TaskRepository) GetByProject(projectID int) ([]entity.Task, error) {

	query := `
		SELECT id, project_id, title, description, status, created_at
		FROM tasks
		WHERE project_id = $1
	`

	rows, err := config.DB.Query(query, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []entity.Task

	for rows.Next() {
		var t entity.Task
		err := rows.Scan(&t.ID, &t.ProjectID, &t.Title, &t.Description, &t.Status, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}
