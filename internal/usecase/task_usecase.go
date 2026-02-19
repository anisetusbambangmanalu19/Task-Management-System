package usecase

import (
	"github.com/anisetusbambangmanalu19/task-management/internal/entity"
	"github.com/anisetusbambangmanalu19/task-management/internal/repository"
	"errors"
)

type TaskUsecase struct {
	TaskRepo repository.TaskRepository
	ProjectRepo repository.ProjectRepository
}

func (u *TaskUsecase) Create(userID int, projectID int, title, desc string) (*entity.Task, error) {

	isOwner, err := u.ProjectRepo.IsOwner(projectID, userID)
	if err != nil {
		return nil, err
	}

	if !isOwner {
		return nil, errors.New("forbidden")
	}

	task := &entity.Task{
		ProjectID:  projectID,
		Title:      title,
		Description: desc,
		Status:     "todo",
	}

	err = u.TaskRepo.Create(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (u *TaskUsecase) GetByProject(userID int, projectID int) ([]entity.Task, error) {

	isOwner, err := u.ProjectRepo.IsOwner(projectID, userID)
	if err != nil {
		return nil, err
	}

	if !isOwner {
		return nil, errors.New("forbidden")
	}

	return u.TaskRepo.GetByProject(projectID)
}
