package handler

import (
	m "app/model"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
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

	// asynq
	payload, err := json.Marshal(report)

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
