package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func MarkTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid task ID"})
	}

	var statusUpdate struct {
		Status string `json:"status"`
	}
	if err := c.Bind(&statusUpdate); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	tasksMutex.Lock()
	defer tasksMutex.Unlock()

	if task, exists := tasks[id]; exists {
		task.Status = statusUpdate.Status
		tasks[id] = task
		return c.JSON(http.StatusOK, task)
	} else {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
	}
}
