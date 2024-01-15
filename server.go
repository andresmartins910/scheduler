package main

import (
	// "github.com/hibiken/asynq"
	h "app/handler"
	m "app/model"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// func init() {
// 	// Inicia o servidor de mensagens
// 	srv := asynq.NewServer(
// 		asynq.RedisClientOpt{Addr: "127.0.0.1:6379"},
// 		asynq.Config{
// 			// Specify how many concurrent workers to use
// 			Concurrency: 10,
// 			// Optionally specify multiple queues with different priority.
// 			Queues: map[string]int{
// 				"critical": 6,
// 				"default":  3,
// 				"low":      1,
// 			},
// 			// See the godoc for other configuration options
// 		},
// 	)

// 	// Inicia o servidor de mensagens
// 	if err := srv.Run(); err != nil {
// }

func main() {
	e := echo.New()

	dsn := "root:220422@ndrE@tcp(127.0.0.1:3306)/scheduler?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&m.Task{})

	handler := &h.Handler{DB: db}

	// Routes
	e.GET("/", func(c echo.Context) error {
		return h.GetTasksHandler(c, handler)
	})
	e.GET("/:id", func(c echo.Context) error {
		return h.GetTaskByIdHandler(c, handler)
	})
	e.POST("/", func(c echo.Context) error {
		return h.CreateTaskHandler(c, handler)
	})
	e.PUT("/:id", func(c echo.Context) error {
		return h.UpdateTaskHandler(c, handler)
	})
	e.DELETE("/:id", func(c echo.Context) error {
		return h.DeleteTaskHandler(c, handler)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
