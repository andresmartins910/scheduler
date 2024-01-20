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
	db.AutoMigrate(&m.Report{})

	return db
}

func populateTestTable(db *gorm.DB) []m.Report {
	reports := []m.Report{
		{Name: "Task 1", Description: "Description 1"},
		{Name: "Task 2", Description: "Description 2"},
		{Name: "Task 3", Description: "Description 3"},
	}

	for _, report := range reports {
		db.Create(&report)
	}

	return reports
}

func TestGetReportsHandler(t *testing.T) {
	db := newTestDatabase(t) // Cria uma instância do banco de dados de teste
	populateTestTable(db)    // Popula a tabela de teste com dados de teste

	tx := db.Begin()    // Inicia uma transação
	defer tx.Rollback() // Rollback da transação se ocorrer algum erro durante o teste

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := h.Handler{DB: db} // Cria uma instância do handler do DB

	if err := h.GetReportsHandler(c, &handler); err != nil {
		t.Fatal(err)
	}

	got := []m.Report{}

	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatal(err)
	}

	want := append([]m.Report{}, got...)

	assert.Equal(t, http.StatusOK, rec.Code, "they should be equal") // Verifica se o código de status da resposta é o esperado
	assert.Len(t, got, len(want), "they should be equal")            // Verifica se o tamanho da resposta é o esperado
	assert.Equal(t, want, got, "they should be equal")               // Verifica se as respostas são iguais
}
