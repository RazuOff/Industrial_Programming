package taskhandler

import (
	"ginexample.com/pkg/tasks"
	"github.com/gin-gonic/gin"
)

// GetTask godoc
// @Summary      Get task by ID
// @Description  Retrieves the details of a specific task by its ID
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Task ID"
// @Success      200  {object}  tasks.Task  "Task details"
// @Failure      404  {object}  map[string]string  "Task not found"
// @Router       /tasks/{id} [get]
func GetTask(c *gin.Context) {
	taskID := c.Param("id")
	task := tasks.GetTask(taskID)
	if task == nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(200, task)
}
