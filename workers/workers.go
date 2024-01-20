package main

import (
	"app/task"
	"log"

	"github.com/hibiken/asynq"
)

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{Concurrency: 10},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(task.TypeWelcomeEmail, task.HandleWelcomeEmailTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
