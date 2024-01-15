package handler_test

import (
	"encoding/json"
	// "fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	h "app/handler"
	m "app/model"
)

func newTestDatabase(t *testing.T) *gorm.DB {
	dsn := "root:220422@ndrE@tcp(127.0.0.1:3306)/scheduler_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Unix(0, 0)
		},
	})

	if err != nil {
		t.Fatal("Failed to connect to database")
	}

	db.Exec("DROP TABLE IF EXISTS tasks")
	db.AutoMigrate(&m.Task{})

	return db
}

func populateTestTable(db *gorm.DB) []m.Task {
	tasks := []m.Task{
		{Title: "Task 1", Status: "TODO"},
		{Title: "Task 2", Status: "In Progress"},
		{Title: "Task 3", Status: "Done"},
	}

	for _, task := range tasks {
		db.Create(&task)
	}

	return tasks
}

func TestGetTasksHandler(t *testing.T) {
	db := newTestDatabase(t) // Cria uma instância do banco de dados de teste
	populateTestTable(db)    // Popula a tabela de teste com dados de teste

	tx := db.Begin()    // Inicia uma transação
	defer tx.Rollback() // Rollback da transação se ocorrer algum erro durante o teste

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := h.Handler{DB: db} // Cria uma instância do handler do DB

	if err := h.GetTasksHandler(c, &handler); err != nil {
		t.Fatal(err)
	}

	got := []m.Task{}

	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatal(err)
	}

	want := append([]m.Task{}, got...)

	assert.Equal(t, http.StatusOK, rec.Code, "they should be equal") // Verifica se o código de status da resposta é o esperado
	assert.Len(t, got, len(want), "they should be equal")            // Verifica se o tamanho da resposta é o esperado
	assert.Equal(t, want, got, "they should be equal")               // Verifica se as respostas são iguais
}
