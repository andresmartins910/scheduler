package handler

import (
	m "app/model"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/labstack/echo/v4"
	"log"
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
	// db
	var task m.Task

	c.Bind(&task)

	h.DB.Create(&task)

	// asynq
	payload, err := json.Marshal(task)

	if err != nil {
		return err
	}

	t := asynq.NewTask("dont.know", payload)

	log.Printf(" [*] Enqueuing task: %+v", t)

	info, err := h.Client.Enqueue(t)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf(" [*] Successfully enqueued task: %+v", info)

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
