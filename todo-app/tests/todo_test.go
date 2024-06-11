package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo-app/handlers"
	"todo-app/models"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBufferString(`{"title": "Test Task", "priority": 1, "due_date": "2024-06-15", "status": "open"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handlers.CreateTask(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		var task models.Task
		err := json.Unmarshal(rec.Body.Bytes(), &task)
		assert.NoError(t, err)

	}
}

func TestListTasks(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handlers.ListTasks(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var tasks []models.Task
		err := json.Unmarshal(rec.Body.Bytes(), &tasks)
		assert.NoError(t, err)

	}
}

func TestEditTask(t *testing.T) {

	e := echo.New()

	createReq := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBufferString(`{"title": "Test Task", "priority": 1, "due_date": "2024-06-15", "status": "open"}`))
	createReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	createRec := httptest.NewRecorder()
	createC := e.NewContext(createReq, createRec)
	assert.NoError(t, handlers.CreateTask(createC))

	editReq := httptest.NewRequest(http.MethodPut, "/tasks/1", bytes.NewBufferString(`{"title": "Edited Task", "priority": 2, "due_date": "2024-06-16", "status": "open"}`))
	editReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	editRec := httptest.NewRecorder()
	editC := e.NewContext(editReq, editRec)
	editC.SetParamNames("id")
	editC.SetParamValues("1")

	if assert.NoError(t, handlers.EditTask(editC)) {
		assert.Equal(t, http.StatusOK, editRec.Code)

		var task models.Task
		err := json.Unmarshal(editRec.Body.Bytes(), &task)
		assert.NoError(t, err)

	}
}

func TestMarkTask(t *testing.T) {

	e := echo.New()

	createReq := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBufferString(`{"title": "Test Task", "priority": 1, "due_date": "2024-06-15", "status": "open"}`))
	createReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	createRec := httptest.NewRecorder()
	createC := e.NewContext(createReq, createRec)
	assert.NoError(t, handlers.CreateTask(createC))

	markReq := httptest.NewRequest(http.MethodPut, "/tasks/1/status", bytes.NewBufferString(`{"status": "completed"}`))
	markReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	markRec := httptest.NewRecorder()
	markC := e.NewContext(markReq, markRec)
	markC.SetParamNames("id")
	markC.SetParamValues("1")

	if assert.NoError(t, handlers.MarkTask(markC)) {
		assert.Equal(t, http.StatusOK, markRec.Code)

		var task models.Task
		err := json.Unmarshal(markRec.Body.Bytes(), &task)
		assert.NoError(t, err)

	}
}
