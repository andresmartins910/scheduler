package handler

import (
	m "app/model"
	"app/task"
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetReportsHandler(c echo.Context, h *Handler) error {
	var reports []m.Report

	h.DB.Find(&reports)

	return c.JSONPretty(http.StatusOK, reports, "  ")
}

func GetReportByIdHandler(c echo.Context, h *Handler) error {
	var report m.Report

	h.DB.First(&report, c.Param("id"))

	return c.JSONPretty(http.StatusOK, report, "  ")
}

func CreateReportHandler(c echo.Context, h *Handler) error {
	// db
	var report m.Report

	c.Bind(&report)

	h.DB.Create(&report)

	ctx := context.Background()

	// asynq
	t, err := task.NewReportTask(ctx, report)
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}

	info, err := h.Client.Enqueue(t)
    if err != nil {
        log.Fatalf("could not enqueue task: %v", err)
    }
    log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	return c.JSONPretty(http.StatusOK, report, "  ")
}

func UpdateReportHandler(c echo.Context, h *Handler) error {
	var report m.Report

	h.DB.First(&report, c.Param("id"))

	c.Bind(&report)

	h.DB.Save(&report)

	return c.JSONPretty(http.StatusOK, report, "  ")
}

func DeleteReportHandler(c echo.Context, h *Handler) error {
	var report m.Report

	h.DB.First(&report, c.Param("id"))

	h.DB.Delete(&report)

	return c.JSONPretty(http.StatusOK, report, "  ")
}
