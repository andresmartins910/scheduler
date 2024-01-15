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

// func GetTask(c echo.Context) error {
// 	var task m.Task
// 	DB.First(&task, c.Param("id"))

// 	return c.JSONPretty(http.StatusOK, task, "  ")
// }

// func CreateTask(c echo.Context) error {
// 	var task m.Task
// 	c.Bind(&task)

// 	result := DB.Create(&task)

// 	if result.Error != nil {
// 		return result.Error
// 	}

// 	return c.JSONPretty(http.StatusOK, task, "  ")
// }

// func UpdateTask(c echo.Context) error {
// 	var task m.Task
// 	DB.First(&task, c.Param("id"))
// 	c.Bind(&task)

// 	DB.Save(&task)

// 	return c.JSONPretty(http.StatusOK, task, "  ")
// }

// func DeleteTask(c echo.Context) error {
// 	var task m.Task
// 	DB.First(&task, c.Param("id"))
// 	DB.Delete(&task)

// 	return c.JSONPretty(http.StatusOK, task, "  ")
// }
