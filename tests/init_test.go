package tests

import (
	"Effective-Mobile/internal/db"
	h "Effective-Mobile/pkg/handlers"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/txix-open/isp-kit/http/httpcli"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestSuite struct {
	suite.Suite
	url     string
	test    *assert.Assertions
	httpCli *httpcli.Client
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) SetupSuite() {
	s.T().Helper()
	s.test = assert.New(s.T())
	s.httpCli = httpcli.New()

	// Настройка тестовой базы данных
	db := db.New("test", "test", "localhost", "5432", "test")
	err := goose.Up(db.GetDB(), "../db/migrations")
	if err != nil {
		s.T().Fatal(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}()

	r := setupHttpRouter(db)
	srv := httptest.NewServer(r)

	s.url = srv.URL

	defer srv.Close()
}

func (s *TestSuite) TestCreateCar() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resp, err := s.httpCli.Get(s.url + "/info").Do(ctx)
	s.test.Nil(err)
	s.test.Equal(http.StatusOK, resp.StatusCode())
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
