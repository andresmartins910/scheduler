package handler

import (
	m "app/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetTasksHandler(c echo.Context, h *Handler) error {
	var tasks []m.Task

	h.DB.Find(&tasks)

	return c.JSONPretty(http.StatusOK, tasks, "  ")
}

func GetTaskByIdHandler(c echo.Context, h *Handler) error {
	var task m.Task

	h.DB.First(&task, c.Param("id"))

	return c.JSONPretty(http.StatusOK, task, "  ")
}

func CreateTaskHandler(c echo.Context, h *Handler) error {
	var task m.Task

	c.Bind(&task)

	h.DB.Create(&task)

	return c.JSONPretty(http.StatusOK, task, "  ")
}

func UpdateTaskHandler(c echo.Context, h *Handler) error {
	var task m.Task

	h.DB.First(&task, c.Param("id"))

	c.Bind(&task)

	h.DB.Save(&task)

	return c.JSONPretty(http.StatusOK, task, "  ")
}

func DeleteTaskHandler(c echo.Context, h *Handler) error {
	var task m.Task

	h.DB.First(&task, c.Param("id"))

	h.DB.Delete(&task)

	return c.JSONPretty(http.StatusOK, task, "  ")
}
