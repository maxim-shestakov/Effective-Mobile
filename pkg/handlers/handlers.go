package handlers

import (
	"log"
	"net/http"
	"strconv"

	st "Effective-Mobile/internal/structures"
	p "Effective-Mobile/pkg/postgresql"

	"github.com/gin-gonic/gin"
)

func filterCars(regNum, mark, model string, cars []st.Car) []st.Car {
	var filteredCars []st.Car

	for _, car := range cars {
		if (regNum == "" || car.Regnum == regNum) &&
			(mark == "" || car.Mark == mark) &&
			(model == "" || car.Model == model) {
			filteredCars = append(filteredCars, car)
		}
	}
	return filteredCars
}

// GetCars godoc
// @Summary Get cars
// @Tags API Functions
// @Description Get all cars from database or filtered by regnum, mark and model (optional). All filters should be written as query parameters.
// @ID get-all-cars
// @Accept json
// @Produce json
// @Param regnum query string false "Registration number"
// @Param mark query string false "Car mark"
// @Param model query string false "Car model"
// @Success 200 {object} st.StatusOKMessage "ok"
// @Failure 500 {object} st.StatusInternalServerErrorMessage "internal server error"
// @Failure 400 {object} st.StatusBadRequestMessage "bad request"
// @Failure 404 {object} st.StatusNotFoundMessage "not found"
// @Router /info [get]
func GetCars(c *gin.Context) {
	cars, err := p.GetCars()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can't get cars from db"})
		return
	}
	filterRegnum := c.Query("regnum")
	filterMark := c.Query("mark")
	filterModel := c.Query("model")
	pg, ok := c.GetQuery("page")
	if !ok {
		pg = "1"
	}
	page, _ := strconv.Atoi(pg)
	lm, ok := c.GetQuery("limit")
	if !ok {
		lm = "20"
	}
	limit, _ := strconv.Atoi(lm)

	start := (page - 1) * limit
	end := page * limit

	cars = filterCars(filterRegnum, filterMark, filterModel, cars)
	if end > len(cars) {
		end = len(cars)
	}
	
	PaginatedCars := cars[start:end]

	log.Println("cars received") //debug log
	c.JSON(http.StatusOK, PaginatedCars)
}

// CreateCar godoc
// @Summary Create car
// @Tags API Functions
// @Description Add a new car to the database from JSON input.
// @ID create-car
// @Accept json
// @Produce json
// @Param input body st.Car true "Car info (regnum, mark, model, owner_id are required)."
// @Success 200 {object} st.StatusOKMessage "ok"
// @Failure 500 {object} st.StatusInternalServerErrorMessage "internal server error"
// @Failure 400 {object} st.StatusBadRequestMessage "bad request"
// @Failure 404 {object} st.StatusNotFoundMessage "not found"
// @Router /cars [post]
func CreateCar(c *gin.Context) {
	var car st.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't bind json body to struct"})
		return
	}
	err := p.AddCar(car)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can't add a car to the db"})
		return
	}
	log.Println("car created") //debug log
	c.JSON(http.StatusOK, gin.H{"message": "car created"})
}

// UpdateCar godoc
// @Summary Update car
// @Tags API Functions
// @Description Update car info in the database by id.
// @ID update-car
// @Accept json
// @Produce json
// @Param input body st.Car true "Change car info, ID and other one or more fields are required. If the field is empty, it will not be changed ."
// @Success 200 {object} st.StatusOKMessage "ok"
// @Failure 500 {object} st.StatusInternalServerErrorMessage "internal server error"
// @Failure 400 {object} st.StatusBadRequestMessage "bad request"
// @Failure 404 {object} st.StatusNotFoundMessage "not found"
// @Router /cars/{id} [put]
func UpdateCar(c *gin.Context) {
	var car st.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't bind json body to struct"})
		return
	}
	err := p.UpdateCar(car)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can't update a car in the db"})
		return
	}
	log.Println("car updated") //debug log
	c.JSON(http.StatusOK, gin.H{"message": "car updated"})
}

// DeleteCar godoc
// @Summary Delete car
// @Tags API Functions
// @Description Delete a car from the database by ID.
// @ID delete-car
// @Accept json
// @Produce json
// @Param id path string true "Car ID"
// @Success 200 {object} st.StatusOKMessage "ok"
// @Failure 500 {object} st.StatusInternalServerErrorMessage "internal server error"
// @Failure 400 {object} st.StatusBadRequestMessage "bad request"
// @Failure 404 {object} st.StatusNotFoundMessage "not found"
// @Router /cars/{id} [delete]
func DeleteCar(c *gin.Context) {
	id := c.Param("id")
	err := p.DeleteCar(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can't delete a car from the db"})
		return
	}
	log.Println("car deleted") //debug log
	c.JSON(http.StatusOK, gin.H{"message": "car deleted"})
}

// AddOwner godoc
// @Summary Add owner
// @Tags API Functions
// @Description Add a new owner to the database from JSON input body.
// @ID add-owner
// @Accept json
// @Produce json
// @Param input body st.Owner true "Owner info (only name, surname are required)."
// @Success 200 {object} st.StatusOKMessage "ok"
// @Failure 500 {object} st.StatusInternalServerErrorMessage "internal server error"
// @Failure 400 {object} st.StatusBadRequestMessage "bad request"
// @Failure 404 {object} st.StatusNotFoundMessage "not found"
// @Router /owners [post]
func AddOwner(c *gin.Context) {
	var owner st.Owner
	if err := c.ShouldBindJSON(&owner); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't bind json body to struct"})
		return
	}
	err := p.AddOwner(owner)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can't add an owner to the db"})
		return
	}
	log.Println("owner created") //debug log
	c.JSON(http.StatusOK, gin.H{"message": "owner created"})
}
