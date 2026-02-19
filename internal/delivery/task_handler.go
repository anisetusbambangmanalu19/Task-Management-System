package delivery

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/anisetusbambangmanalu19/task-management/internal/usecase"
)

type TaskHandler struct {
	TaskUsecase usecase.TaskUsecase
}

func (h *TaskHandler) Create(c *gin.Context) {

	userID := int(c.GetFloat64("user_id"))
	projectID, _ := strconv.Atoi(c.Param("project_id"))

	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.TaskUsecase.Create(userID, projectID, req.Title, req.Description)
	if err != nil {
		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, gin.H{"error": "not your project"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}


func (h *TaskHandler) GetByProject(c *gin.Context) {

	userID := int(c.GetFloat64("user_id"))
	projectID, _ := strconv.Atoi(c.Param("project_id"))

	tasks, err := h.TaskUsecase.GetByProject(userID, projectID)
	if err != nil {
		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, gin.H{"error": "not your project"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

