package routes

import (
	"todo-app/handlers"

	"github.com/labstack/echo/v4"
)

func Router() *echo.Echo {
	e := echo.New()

	e.POST("/tasks", handlers.CreateTask)
	e.PUT("/tasks/:id", handlers.EditTask)
	e.GET("/tasks", handlers.ListTasks)
	e.PUT("/tasks/:id/status", handlers.MarkTask)

	return e
}
