package handlers

import (
	"net/http"
	"sync"
	"todo-app/models"

	"github.com/labstack/echo/v4"
)

var (
	tasks      = make(map[int]models.Task)
	tasksMutex = &sync.Mutex{}
	idCounter  = 1
)

func CreateTask(c echo.Context) error {
	var task models.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if task.Priority == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Priority must be provided"})
	}

	tasksMutex.Lock()
	defer tasksMutex.Unlock()
	task.Id = idCounter
	idCounter++
	task.Status = "open"

	tasks[task.Id] = task

	return c.JSON(http.StatusCreated, task)
}
