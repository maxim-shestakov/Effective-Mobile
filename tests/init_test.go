package tests

import (
	"Effective-Mobile/internal/db"
	h "Effective-Mobile/pkg/handlers"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/txix-open/isp-kit/http/httpcli"
	"log"
	"math/rand/v2"
	"net/http/httptest"
	"strconv"
	"testing"
)

type TestSuite struct {
	suite.Suite
	url     string
	test    *assert.Assertions
	httpCli *httpcli.Client
	schema  string
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) SetupSuite() {
	s.T().Helper()
	s.test = assert.New(s.T())
	s.httpCli = httpcli.New()

	// Настройка тестовой базы данных
	schema, err := createTestSchema()
	if err != nil {
		s.T().Fatal(err)
	}
	s.schema = schema
	db := db.New("test", "test", "localhost", "5432", "test", schema)
	err = goose.Up(db.GetDB(), "../db/migrations")
	if err != nil {
		s.T().Fatal(err)
	}

	r := setupHttpRouter(db)
	srv := httptest.NewServer(r)

	s.url = srv.URL
}

func (s *TestSuite) TearDownSuite() {
	s.T().Helper()
	// Удаляем тестовую схему
	schema := s.schema
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", "test", "test", "localhost", "5432", "test")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	query := fmt.Sprintf("DROP SCHEMA IF EXISTS %s CASCADE", schema)
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func setupHttpRouter(db *db.Repository) *gin.Engine {
	// Запуск сервера на отдельном порту
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/info", h.GetCars(db))
	r.PUT("/cars/:id", h.UpdateCar(db))
	r.POST("/cars", h.CreateCar(db))
	r.POST("/owners", h.AddOwner(db))
	r.DELETE("/cars/:id", h.DeleteCar(db))

	return r
}

func createTestSchema() (string, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", "test", "test", "localhost", "5432", "test")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Создаем схему для теста
	schema := "test_" + strconv.Itoa(rand.IntN(10000))
	query := fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schema)
	_, err = db.Exec(query)
	if err != nil {
		return "", err
	}

	return schema, nil
}
