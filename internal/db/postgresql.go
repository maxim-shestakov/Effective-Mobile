package db

import (
	"fmt"
	"log"


	_ "github.com/lib/pq"
)



// AddOwner inserts a new owner into the database.
//
// Parameter: owner of type Owner.
// Return type: error.
func (p *Postgresql) AddOwner(owner Owner) error {
	_, err := p.db.Exec("INSERT INTO owners (name, surname, patronymic) VALUES ($1, $2, $3)", owner.Name, owner.Surname, owner.Patronymic)
	if err != nil {
		log.Println(err)
	}
	return err
}


// AddCar inserts a car into the database.
//
// Parameter: car of type Car.
// Return type: error.
func (p *Postgresql) AddCar(car Car) error {
	_, err := p.db.Exec("INSERT INTO cars (regnum, mark, model, year, owner_id) VALUES ($1, $2, $3, $4, $5)", car.Regnum, car.Mark, car.Model, car.Year, car.OwnerID)
	if err != nil {
		log.Println(err)
	}
	return err
}

// GetCars retrieves all cars from the database.
//
// No parameters.
// Returns a slice of Car structs and an error.
func (p *Postgresql) GetCars() ([]Car, error) {
	rows, err := p.db.Query("SELECT * FROM cars")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	cars := []Car{}
	for rows.Next() {
		var car Car
		if err := rows.Scan(&car.ID, &car.Regnum, &car.Mark, &car.Model, &car.Year, &car.OwnerID); err != nil {
			log.Println(err)
			return nil, err
		}
		cars = append(cars, car)
	}
	return cars, nil
}


// UpdateCar updates the information of a car in the database.
//
// Parameter:
//   car Car: the car struct containing the information to be updated.
// Return:
//   error: an error if the update operation fails.
func (p *Postgresql) UpdateCar(car Car) error {
	args := []interface{}{}
	counter := 1
	query := "UPDATE cars SET"
	if car.Regnum!= "" { //When regnum is not empty we should update it
		query += fmt.Sprintf(" regnum=$%d,", counter)
		args = append(args, car.Regnum)
		counter++
	}
	if car.Mark != "" {
		query += fmt.Sprintf(" mark=$%d,", counter)
		args = append(args, car.Mark)
		counter++
	}
	if car.Model != "" {
		query += fmt.Sprintf(" model=$%d,", counter)
		args = append(args, car.Model)
		counter++
	}
	if car.Year != 0 {
		query += fmt.Sprintf(" year=$%d,", counter)
		args = append(args, car.Year)
		counter++
	}
	
	if counter == 1 {
		log.Println("no fields to update")
		return nil
	}

	query = query[:len(query)-1]
	query += fmt.Sprintf(" WHERE id=$%d", counter)
	args = append(args, car.ID)
	_, err := p.db.Exec(query, args...)
	if err != nil {
		log.Println("problem with updating information about film", err)
		return err
	}
	return nil
}


// DeleteCar deletes a car from the database based on the provided id.
//
// Parameter: id string
// Return type: error
func (p *Postgresql) DeleteCar(id string) error {
	_, err := p.db.Exec("DELETE FROM cars WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}