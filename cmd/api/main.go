package main

import (
	"log"
	"scheduler/config"
	h "scheduler/pkg/handler"
	m "scheduler/pkg/model"

	"github.com/hibiken/asynq"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})

	if client == nil {
		log.Fatal("failed to create asynq client")
	}

	dsn := "root:220422@ndrE@tcp(127.0.0.1:3306)/scheduler?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&m.Report{})

	handler := &h.Handler{DB: db, Client: client}

	e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == config.GetToken(), nil
	}))

	// Routes
	e.GET("/", func(c echo.Context) error {
		return h.GetReportsHandler(c, handler)
	})
	e.GET("/:id", func(c echo.Context) error {
		return h.GetReportByIdHandler(c, handler)
	})
	e.POST("/", func(c echo.Context) error {
		return h.CreateReportHandler(c, handler)
	})
	e.PUT("/:id", func(c echo.Context) error {
		return h.UpdateReportHandler(c, handler)
	})
	e.DELETE("/:id", func(c echo.Context) error {
		return h.DeleteReportHandler(c, handler)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
