package usecase

import (
	"github.com/anisetusbambangmanalu19/task-management/internal/entity"
	"github.com/anisetusbambangmanalu19/task-management/internal/repository"
)

type ProjectUsecase struct {
	ProjectRepo repository.ProjectRepository
}

func (u *ProjectUsecase) Create(userID int, name, desc string) (*entity.Project, error) {

	project := &entity.Project{
		UserID:      userID,
		Name:        name,
		Description: desc,
	}

	err := u.ProjectRepo.Create(project)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (u *ProjectUsecase) GetMyProjects(userID int) ([]entity.Project, error) {
	return u.ProjectRepo.GetByUser(userID)
}
