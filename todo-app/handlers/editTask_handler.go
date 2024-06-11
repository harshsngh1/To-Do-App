package handlers

import (
	"net/http"
	"strconv"
	"todo-app/models"

	"github.com/labstack/echo/v4"
)

func EditTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid task ID"})
	}

	var updatedTask models.Task
	if err := c.Bind(&updatedTask); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	tasksMutex.Lock()
	defer tasksMutex.Unlock()

	if task, exists := tasks[id]; exists {
		updatedTask.Id = task.Id
		tasks[id] = updatedTask
		return c.JSON(http.StatusOK, updatedTask)
	} else {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
	}
}
