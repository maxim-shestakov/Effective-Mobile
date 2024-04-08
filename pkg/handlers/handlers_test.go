package handlers

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"strconv"
	"testing"

	"Effective-Mobile/internal/db"
	"Effective-Mobile/internal/mocks"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestCreateCar(t *testing.T) {
	type mockBehavior func(s *mocks.MockStorage, car db.Car)

	tests := []struct {
		name                string
		inputBody           string
		inputCar            db.Car
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"regnum":"A111AA155","mark":"BMW","model":"X5","year":2015,"owner_id":1}`,
			inputCar: db.Car{
				Regnum:  "A111AA155",
				Mark:    "BMW",
				Model:   "X5",
				Year:    2015,
				OwnerID: 1,
			},
			mockBehavior: func(s *mocks.MockStorage, car db.Car) {
				s.EXPECT().AddCar(car).Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"message":"car created"}`,
		},
		{
			name:                "Wrong input",
			inputBody:           `{"regnum":"","mark":"BMW","model":"X5","year":2015,"owner_id":""}`,
			mockBehavior:        func(s *mocks.MockStorage, car db.Car) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"error":"can't bind json body to struct"}`,
		},
		{
			name:      "Server error",
			inputBody: `{"regnum":"A111AA155","mark":"BMW","model":"X5","year":2015,"owner_id":1}`,
			inputCar: db.Car{
				Regnum:  "A111AA155",
				Mark:    "BMW",
				Model:   "X5",
				Year:    2015,
				OwnerID: 1,
			},
			mockBehavior: func(s *mocks.MockStorage, car db.Car) {

				s.EXPECT().AddCar(car).Return(errors.New("can't add a car to the db"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"error":"can't add a car to the db"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			// Init Dependencies
			carStorage := mocks.NewMockStorage(c)
			test.mockBehavior(carStorage, test.inputCar)

			// Init gin
			router := gin.New()
			router.POST("/cars", CreateCar(carStorage))

			// Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/cars", bytes.NewBufferString(test.inputBody))

			// Perform request
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedRequestBody, w.Body.String())

		})

	}

}

func TestDeleteCar(t *testing.T) {
	type mockBehavior func(s *mocks.MockStorage, id string)

	tests := []struct {
		name                string
		inputId             string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:    "OK",
			inputId: "1",
			mockBehavior: func(s *mocks.MockStorage, id string) {
				s.EXPECT().DeleteCar(id).Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"message":"car deleted"}`,
		},
		{
			name:    "Server error",
			inputId: "1",
			mockBehavior: func(s *mocks.MockStorage, id string) {
				s.EXPECT().DeleteCar(id).Return(errors.New("can't delete a car from the db"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"error":"can't delete a car from the db"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			// Init Dependencies
			carStorage := mocks.NewMockStorage(c)
			test.mockBehavior(carStorage, test.inputId)

			// Init gin
			router := gin.New()
			router.DELETE("/cars/:id", DeleteCar(carStorage))

			// Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", "/cars/"+test.inputId, nil)

			// Perform request
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedRequestBody, w.Body.String())

		})

	}
}

func TestUpdateCar(t *testing.T) {
	type mockBehavior func(s *mocks.MockStorage, car db.Car)

	tests := []struct {
		name                string
		inputBody           string
		inputCar            db.Car
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"id":1,"regnum":"A111AA155","mark":"BMW","model":"X5","year":2015,"owner_id":1}`,
			inputCar: db.Car{
				ID:      1,
				Regnum:  "A111AA155",
				Mark:    "BMW",
				Model:   "X5",
				Year:    2015,
				OwnerID: 1,
			},
			mockBehavior: func(s *mocks.MockStorage, car db.Car) {
				s.EXPECT().UpdateCar(car).Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"message":"car updated"}`,
		},
		{
			name:                "Wrong input",
			inputBody:           `{"mark":"BMW","model":"X5","year":2015,"owner_id":""}`,
			mockBehavior:        func(s *mocks.MockStorage, car db.Car) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"error":"can't bind json body to struct"}`,
		},
		{
			name:      "Server error",
			inputBody: `{"id":1,"regnum":"A111AA155","mark":"BMW","model":"X5","year":2015,"owner_id":1}`,
			inputCar: db.Car{
				ID:      1,
				Regnum:  "A111AA155",
				Mark:    "BMW",
				Model:   "X5",
				Year:    2015,
				OwnerID: 1,
			},
			mockBehavior: func(s *mocks.MockStorage, car db.Car) {
				s.EXPECT().UpdateCar(car).Return(errors.New("can't update a car in the db"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"error":"can't update a car in the db"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			// Init Dependencies
			carStorage := mocks.NewMockStorage(c)
			test.mockBehavior(carStorage, test.inputCar)

			// Init gin
			router := gin.New()
			router.PUT("/cars/:id", UpdateCar(carStorage))

			// Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "/cars/"+strconv.Itoa(test.inputCar.ID), bytes.NewBufferString(test.inputBody))

			// Perform request
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedRequestBody, w.Body.String())

		})
	}
}

func TestGetCars(t *testing.T) {
	type mockBehavior func(s *mocks.MockStorage, cars []db.Car)

	tests := []struct {
		name                string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
		cars                []db.Car
	}{
		{
			name: "OK",
			mockBehavior: func(s *mocks.MockStorage, cars []db.Car) {
				s.EXPECT().GetCars().Return(cars, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `[{"id":1,"regnum":"A111AA155","mark":"BMW","model":"X5","year":2015,"owner_id":1}]`,
			cars: []db.Car{
				{
					ID:      1,
					Regnum:  "A111AA155",
					Mark:    "BMW",
					Model:   "X5",
					Year:    2015,
					OwnerID: 1,
				},
			},
		},
		{
			name: "Server error",
			mockBehavior: func(s *mocks.MockStorage, cars []db.Car) {
				s.EXPECT().GetCars().Return(nil, errors.New("can't get cars from the db"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"error":"can't get cars from the db"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			// Init Dependencies
			carStorage := mocks.NewMockStorage(c)
			test.mockBehavior(carStorage, test.cars)

			// Init gin
			router := gin.New()
			router.GET("/cars", GetCars(carStorage))

			// Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/cars", nil)

			// Perform request
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedRequestBody, w.Body.String())

		})
	}
}

func TestAddOwner(t *testing.T) {
	type mockBehavior func(s *mocks.MockStorage, owner db.Owner)

	tests := []struct {
		name                string
		inputBody           string
		inputOwner          db.Owner
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"Max","surname":"Shestakov","patronymic":"Olegovich"}`,
			inputOwner: db.Owner{
				Name:       "Max",
				Surname:    "Shestakov",
				Patronymic: "Olegovich",
			},
			mockBehavior: func(s *mocks.MockStorage, owner db.Owner) {
				s.EXPECT().AddOwner(owner).Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"message":"owner created"}`,
		},
		{
			name:                "Wrong input",
			inputBody:           `{"name": 1,"surname":"Shestakov","patronymic":"Olegovich"}`,
			mockBehavior:        func(s *mocks.MockStorage, owner db.Owner) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"error":"can't bind json body to struct"}`,
		},
		{
			name:      "Server error",
			inputBody: `{"name":"Max","surname":"Shestakov","patronymic":"Olegovich"}`,
			inputOwner: db.Owner{
				Name:       "Max",
				Surname:    "Shestakov",
				Patronymic: "Olegovich",
			},
			mockBehavior: func(s *mocks.MockStorage, owner db.Owner) {
				s.EXPECT().AddOwner(owner).Return(errors.New("can't add an owner to the db"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"error":"can't add an owner to the db"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			// Init Dependencies
			ownerStorage := mocks.NewMockStorage(c)
			test.mockBehavior(ownerStorage, test.inputOwner)

			// Init gin
			router := gin.New()
			router.POST("/owners", AddOwner(ownerStorage))

			// Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/owners", bytes.NewBufferString(test.inputBody))

			// Perform request
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedRequestBody, w.Body.String())

		})
	}
}

func TestFilterCars(t *testing.T) {
	regnum := "A111AA155"
	mark := "BMW"
	model := "X5"
	cars := []db.Car{
		{
			ID:      1,
			Regnum:  "A111AA155",
			Mark:    "BMW",
			Model:   "X5",
			Year:    2015,
			OwnerID: 1,
		},
		{
			ID:      2,
			Regnum:  "A111AA156",
			Mark:    "BMW",
			Model:   "X6",
			Year:    2016,
			OwnerID: 1,
		},
		{
			ID:      3,
			Regnum:  "A111AA157",
			Mark:    "BMW",
			Model:   "X7",
			Year:    2017,
			OwnerID: 1,
		},
	}

	expectedFilteredCars := []db.Car{
		{
			ID:      1,
			Regnum:  "A111AA155",
			Mark:    "BMW",
			Model:   "X5",
			Year:    2015,
			OwnerID: 1,
		}}

	realResult := filterCars(regnum, mark, model, cars)

	assert.Equal(t, expectedFilteredCars, realResult)
}
