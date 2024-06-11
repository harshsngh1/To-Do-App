package handlers

import (
	"net/http"
	"sort"
	"strconv"
	"time"
	"todo-app/models"

	"github.com/labstack/echo/v4"
)

func ListTasks(c echo.Context) error {
	tasksMutex.Lock()
	defer tasksMutex.Unlock()

	var result []models.Task

	breached := c.QueryParam("breached")
	priority := c.QueryParam("priority")
	sortBy := c.QueryParam("sort")

	for _, task := range tasks {

		if breached == "true" {
			dueDate, err := time.Parse("2006-01-02", task.DueDate)
			if err != nil {
				continue
			}
			if time.Now().After(dueDate) {
				result = append(result, task)
			}
		} else {
			result = append(result, task)
		}
	}

	if priority != "" {
		priorityInt, err := strconv.Atoi(priority)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid priority value"})
		}

		var filteredResult []models.Task
		for _, task := range result {
			if task.Priority == priorityInt {
				filteredResult = append(filteredResult, task)
			}
		}
		result = filteredResult
	}

	if sortBy == "due_date" {
		sort.Slice(result, func(i, j int) bool {
			date1, err1 := time.Parse("2006-01-02", result[i].DueDate)
			date2, err2 := time.Parse("2006-01-02", result[j].DueDate)
			if err1 != nil || err2 != nil {
				return false
			}
			return date1.Before(date2)
		})
	} else if sortBy == "priority" {
		sort.Slice(result, func(i, j int) bool {
			return result[i].Priority < result[j].Priority
		})
	}

	return c.JSON(http.StatusOK, result)
}
