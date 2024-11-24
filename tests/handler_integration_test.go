package tests

import (
	"Effective-Mobile/internal/db"
	"context"
	"net/http"
)

func (s *TestSuite) TestGetCarsWithoutFilter() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var cars []db.Car
	resp, err := s.httpCli.Get(s.url + "/info").
		JsonResponseBody(&cars).
		Do(ctx)

	s.test.Nil(err)
	s.test.NotEmpty(cars)
	s.test.Equal(http.StatusOK, resp.StatusCode())
}

func (s *TestSuite) TestGetCarsWithFilter() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var cars []db.Car
	resp, err := s.httpCli.Get(s.url + "/info?mark=Toyota").
		JsonResponseBody(&cars).
		Do(ctx)

	s.test.Nil(err)
	s.test.Len(cars, 1)
	s.test.Equal("Camry", cars[0].Model)
	s.test.Equal(http.StatusOK, resp.StatusCode())
}

func (s *TestSuite) TestCreateCar() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	requestBody := db.Car{
		Regnum:  "X777XX159",
		Mark:    "BMW",
		Model:   "M8",
		Year:    2023,
		OwnerID: 1,
	}

	resp, err := s.httpCli.Post(s.url + "/cars").
		JsonRequestBody(requestBody).
		Do(ctx)

	s.test.Nil(err)
	s.test.Equal(http.StatusOK, resp.StatusCode())

	var cars []db.Car
	resp, err = s.httpCli.Get(s.url + "/info?mark=BMW&model=M8").
		JsonResponseBody(&cars).
		Do(ctx)

	s.test.Nil(err)
	s.test.Len(cars, 1)

	// trying to add the same car
	resp, err = s.httpCli.Post(s.url + "/cars").
		JsonRequestBody(requestBody).
		Do(ctx)

	s.test.Nil(err)
	s.test.Equal(http.StatusInternalServerError, resp.StatusCode())
}

func (s *TestSuite) TestUpdateCar() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	requestBody := db.Car{
		Regnum: "X123XX150",
		Mark:   "Lada",
		Model:  "Chocolada",
		Year:   2002,
	}

	resp, err := s.httpCli.Put(s.url + "/cars/1").
		JsonRequestBody(requestBody).
		Do(ctx)

	s.test.Nil(err)
	s.test.Equal(http.StatusOK, resp.StatusCode())

	var cars []db.Car
	resp, err = s.httpCli.Get(s.url + "/info?mark=Lada").
		JsonResponseBody(&cars).
		Do(ctx)

	s.test.Nil(err)
	s.test.Len(cars, 1)
	s.test.Equal("Chocolada", cars[0].Model)
}

func (s *TestSuite) TestDeleteCar() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resp, err := s.httpCli.Delete(s.url + "/cars/3").
		Do(ctx)

	s.test.Nil(err)
	s.test.Equal(http.StatusOK, resp.StatusCode())

	var cars []db.Car
	resp, err = s.httpCli.Get(s.url + "/info?mark=BMW&model=X5").
		JsonResponseBody(&cars).
		Do(ctx)

	s.test.Nil(err)
	s.test.Empty(cars)

	// invalid id test
	resp, err = s.httpCli.Delete(s.url + "/cars/first").
		Do(ctx)

	s.test.Nil(err)
	s.test.Equal(http.StatusInternalServerError, resp.StatusCode())
}

func (s *TestSuite) TestAddOwner() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	requestBody := db.Owner{
		Name:       "Maxim",
		Surname:    "Shestakov",
		Patronymic: "Olegovich",
	}

	resp, err := s.httpCli.Post(s.url + "/owners").
		JsonRequestBody(requestBody).
		Do(ctx)

	s.test.Nil(err)
	s.test.Equal(http.StatusOK, resp.StatusCode())
}
