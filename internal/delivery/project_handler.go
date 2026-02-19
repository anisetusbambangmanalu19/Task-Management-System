package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/anisetusbambangmanalu19/task-management/internal/usecase"
)

type ProjectHandler struct {
	ProjectUsecase usecase.ProjectUsecase
}

func (h *ProjectHandler) Create(c *gin.Context) {

	userID := int(c.GetFloat64("user_id"))

	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project, err := h.ProjectUsecase.Create(userID, req.Name, req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, project)
}

func (h *ProjectHandler) MyProjects(c *gin.Context) {

	userID := int(c.GetFloat64("user_id"))

	projects, err := h.ProjectUsecase.GetMyProjects(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, projects)
}
